// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

func (b Board) AfterMove(from AbsPoint, to AbsPoint, turn Orientation) Board {
	board := b.Copy()
	for point, _ := range board.Move(from, to, turn) {
		board[point.Index()] = *point
	}
	return board
}

// An empty PointSet return indicates no changes to the board - an invalid move.
// The board itself is not returned so no modifications are made to the receiver Board.
func (b Board) Move(from AbsPoint, to AbsPoint, turn Orientation) PointSet {
	fromPoint := b[from.Index()]
	if fromPoint.Piece == nil {
		return PointSet{}
	}
	if fromPoint.Orientation != turn {
		return PointSet{}
	}
	if b.HasPawnToPromote() {
		return PointSet{}
	}
	toPoint := b[to.Index()]
	if toPoint.Piece != nil {
		if (toPoint.Orientation == turn) && (fromPoint.Swaps == false) {
			return PointSet{}
		}
	}
	set := make(PointSet)
	for pt, _ := range b.SurroundingPoints(fromPoint) {
		if pt.Piece != nil {
			if (pt.Orientation != fromPoint.Orientation) && pt.Locks {
				return PointSet{}
			}
		}
	}
	// this check is here so detonations happen before guard chaining occurs
	if toPoint.Piece != nil {
		if toPoint.Detonates && (toPoint.Orientation != turn) {
			set[&Point{
				AbsPoint: from,
			}] = struct{}{}
			dset := b.DetonationsFrom(to, nil)
			if len(dset) > 0 {
				for pt, _ := range dset {
					if (pt.File == from.File) && (pt.Rank == from.Rank) {
						continue
					}
					set[&Point{
						AbsPoint: *pt,
					}] = struct{}{}
				}
				b.UpdatePiecePrevious(turn)
				return set
			}
		}
	}
	for pt, _ := range b.SurroundingPoints(toPoint) {
		if pt.Piece != nil {
			if (pt.Orientation != fromPoint.Orientation) && pt.Guards {
				set.SetPointPiece(fromPoint.AbsPoint, nil)
				set.SetPointPiece(pt.AbsPoint, nil)
				if fromPoint.Detonates {
					board := b.Copy()
					board[fromPoint.AbsPoint.Index()].Piece = nil
					board[pt.AbsPoint.Index()].Piece = nil
					board[to.Index()].Piece = fromPoint.Piece
					dset := board.DetonationsFrom(to, nil)
					if len(dset) > 0 {
						for dpt, _ := range dset {
							if (dpt.File == to.File) && (dpt.Rank == to.Rank) {
								continue
							}
							set.SetPointPiece(*dpt, nil)
						}
					}
				} else {
					// a detonator taken by a guard will remove all other possible chaining guards
					// but in these other cases we need to chain all of the other possible guards
					board := b.Copy()
					board[fromPoint.AbsPoint.Index()].Piece = nil
					board[pt.AbsPoint.Index()].Piece = nil
					board[to.Index()].Piece = pt.Piece
					// in the diff we'll at least have the first piece here, maybe updated if another opponent guard is around
					set.SetPointPiece(to, pt.Piece)
				ALLGUARDS:
					for {
						for gpt, _ := range board.SurroundingPoints(toPoint) {
							if gpt.Piece != nil {
								if (gpt.Orientation != board[to.Index()].Piece.Orientation) && gpt.Guards {
									set.SetPointPiece(gpt.AbsPoint, nil)
									set.SetPointPiece(to, gpt.Piece)
									board[gpt.AbsPoint.Index()].Piece = nil
									board[to.Index()].Piece = gpt.Piece
									continue ALLGUARDS
								}
							}
						}
						break
					}
				}
				b.UpdatePiecePrevious(turn)
				pt.Piece.Previous = pt.AbsPoint.Index()
				return set
			}
		}
	}
	if b.MovesFromPoint(fromPoint).Has(to) == false {
		return PointSet{}
	}
	// en passant / in passing
	if (fromPoint.Base == Pawn) && (toPoint.Piece == nil) {
		var piece *Piece
		var expectedToRank uint8
		if turn == Black {
			piece = b[AbsPoint{File: to.File, Rank: 3}.Index()].Piece
			expectedToRank = 2
		} else {
			piece = b[AbsPoint{File: to.File, Rank: 4}.Index()].Piece
			expectedToRank = 5
		}
		if (piece != nil) && (to.Rank == expectedToRank) {
			if (piece.Orientation != turn) && (piece.Base == Pawn) {
				if turn == Black {
					if piece.Previous == (AbsPoint{File: to.File, Rank: 1}).Index() {
						set[&Point{
							AbsPoint: AbsPoint{File: to.File, Rank: 3},
						}] = struct{}{}
						set[&Point{
							AbsPoint: from,
						}] = struct{}{}
						set[&Point{
							AbsPoint: to,
							Piece:    fromPoint.Piece,
						}] = struct{}{}
						b.UpdatePiecePrevious(turn)
						fromPoint.Piece.Previous = from.Index()
						return set
					}
				} else {
					if piece.Previous == (AbsPoint{File: to.File, Rank: 6}).Index() {
						set[&Point{
							AbsPoint: AbsPoint{File: to.File, Rank: 4},
						}] = struct{}{}
						set[&Point{
							AbsPoint: from,
						}] = struct{}{}
						set[&Point{
							AbsPoint: to,
							Piece:    fromPoint.Piece,
						}] = struct{}{}
						b.UpdatePiecePrevious(turn)
						fromPoint.Piece.Previous = from.Index()
						return set
					}
				}
			}
		}
	}
	// castling can only be done when in-between points are unoccupied and check isn't entered, as verified by MovesFromPoint above
	if (fromPoint.Base == King) && (fromPoint.Moved == false) {
		// these to moves are only available when castling is available
		if (to.File == 2) && (to.Rank == 0) {
			set[&Point{
				AbsPoint: from,
			}] = struct{}{}
			set[&Point{
				AbsPoint: AbsPoint{
					File: 0,
					Rank: 0,
				},
			}] = struct{}{}
			set[&Point{
				AbsPoint: to,
				Piece:    fromPoint.Piece,
			}] = struct{}{}
			fromPoint.Moved = true
			set[&Point{
				AbsPoint: AbsPoint{
					File: 3,
					Rank: 0,
				},
				Piece: b[0].Piece,
			}] = struct{}{}
			b[0].Moved = true
			b.UpdatePiecePrevious(turn)
			b[0].Piece.Previous = AbsPoint{File: 0, Rank: 0}.Index()
			fromPoint.Piece.Previous = from.Index()
			return set
		} else if (to.File == 6) && (to.Rank == 0) {
			set[&Point{
				AbsPoint: from,
			}] = struct{}{}
			set[&Point{
				AbsPoint: AbsPoint{
					File: 7,
					Rank: 0,
				},
			}] = struct{}{}
			set[&Point{
				AbsPoint: to,
				Piece:    fromPoint.Piece,
			}] = struct{}{}
			fromPoint.Moved = true
			set[&Point{
				AbsPoint: AbsPoint{
					File: 5,
					Rank: 0,
				},
				Piece: b[7].Piece,
			}] = struct{}{}
			b[7].Moved = true
			b.UpdatePiecePrevious(turn)
			b[7].Piece.Previous = AbsPoint{File: 7, Rank: 0}.Index()
			fromPoint.Piece.Previous = from.Index()
			return set
		} else if (to.File == 2) && (to.Rank == 7) {
			set[&Point{
				AbsPoint: from,
			}] = struct{}{}
			set[&Point{
				AbsPoint: AbsPoint{
					File: 0,
					Rank: 7,
				},
			}] = struct{}{}
			set[&Point{
				AbsPoint: to,
				Piece:    fromPoint.Piece,
			}] = struct{}{}
			fromPoint.Moved = true
			set[&Point{
				AbsPoint: AbsPoint{
					File: 3,
					Rank: 7,
				},
				Piece: b[56].Piece,
			}] = struct{}{}
			b[56].Moved = true
			b.UpdatePiecePrevious(turn)
			b[56].Piece.Previous = AbsPoint{File: 0, Rank: 7}.Index()
			fromPoint.Piece.Previous = from.Index()
			return set
		} else if (to.File == 6) && (to.Rank == 7) {
			set[&Point{
				AbsPoint: from,
			}] = struct{}{}
			set[&Point{
				AbsPoint: AbsPoint{
					File: 7,
					Rank: 7,
				},
			}] = struct{}{}
			set[&Point{
				AbsPoint: to,
				Piece:    fromPoint.Piece,
			}] = struct{}{}
			fromPoint.Moved = true
			set[&Point{
				AbsPoint: AbsPoint{
					File: 5,
					Rank: 7,
				},
				Piece: b[63].Piece,
			}] = struct{}{}
			b[63].Moved = true
			b.UpdatePiecePrevious(turn)
			b[63].Piece.Previous = AbsPoint{File: 7, Rank: 7}.Index()
			fromPoint.Piece.Previous = from.Index()
			return set
		}
	}
	if toPoint.Piece != nil {
		if fromPoint.Swaps && (toPoint.Orientation == turn) {
			set[&Point{
				Piece:    toPoint.Piece,
				AbsPoint: from,
			}] = struct{}{}
			// TODO: no side effects on the receiver
			toPoint.Moved = true
			set[&Point{
				Piece:    fromPoint.Piece,
				AbsPoint: to,
			}] = struct{}{}
			fromPoint.Moved = true
			b.UpdatePiecePrevious(turn)
			fromPoint.Piece.Previous = from.Index()
			toPoint.Piece.Previous = to.Index()
			return set
		}
	}
	set[&Point{
		AbsPoint: from,
	}] = struct{}{}
	fromPoint.Moved = true
	set[&Point{
		Piece:    fromPoint.Piece,
		AbsPoint: to,
	}] = struct{}{}
	b.UpdatePiecePrevious(turn)
	fromPoint.Piece.Previous = from.Index()
	return set
}
