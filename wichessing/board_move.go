// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

func (b Board) AfterMove(from AbsPoint, to AbsPoint, turn Orientation, previousFrom AbsPoint, previousTo AbsPoint) Board {
	board := b.Copy()
	diff, _ := board.Move(from, to, turn, previousFrom, previousTo)
	for point, _ := range diff {
		board[point.Index()] = *point
	}
	return board
}

// An empty PointSet return indicates no changes to the board - an invalid move. Any taken pieces are returned as the second value.
// The board itself is not returned so no modifications are made to the receiver Board.
func (b Board) Move(from AbsPoint, to AbsPoint, turn Orientation, previousFrom AbsPoint, previousTo AbsPoint) (PointSet, map[AbsPoint]*Piece) {
	fromPoint := b[from.Index()]
	if fromPoint.Piece == nil {
		return PointSet{}, map[AbsPoint]*Piece{}
	}
	if fromPoint.Orientation != turn {
		return PointSet{}, map[AbsPoint]*Piece{}
	}
	promoting, _ := b.HasPawnToPromote()
	if promoting {
		return PointSet{}, map[AbsPoint]*Piece{}
	}
	toPoint := b[to.Index()]
	if toPoint.Piece != nil {
		if (toPoint.Orientation == turn) && (fromPoint.Swaps == false) {
			return PointSet{}, map[AbsPoint]*Piece{}
		}
	}
	for pt, _ := range b.SurroundingPoints(fromPoint) {
		if pt.Piece != nil {
			if (pt.Orientation != fromPoint.Orientation) && pt.Locks {
				return PointSet{}, map[AbsPoint]*Piece{}
			}
		}
	}
	set := make(PointSet)
	pieceset := make(map[AbsPoint]*Piece)
	// this check is here so detonations happen before guard chaining occurs
	if toPoint.Piece != nil {
		if toPoint.Detonates && (toPoint.Orientation != turn) {
			set[&Point{
				AbsPoint: from,
			}] = struct{}{}
			pieceset[from] = b[from.Index()].Piece
			dset := b.DetonationsFrom(to, nil)
			if len(dset) > 0 {
				for pt, _ := range dset {
					if (pt.File == from.File) && (pt.Rank == from.Rank) {
						continue
					}
					set[&Point{
						AbsPoint: *pt,
					}] = struct{}{}
					pieceset[*pt] = b[(*pt).Index()].Piece
				}
				return set, pieceset
			}
		}
	}
	for pt, _ := range b.SurroundingPoints(toPoint) {
		if pt.Piece != nil {
			if (pt.Orientation != fromPoint.Orientation) && pt.Guards {
				if (pt.Base == Pawn) && fromPoint.Fortified {
					continue
				}
				set.SetPointPiece(fromPoint.AbsPoint, nil)
				set.SetPointPiece(pt.AbsPoint, nil)
				pieceset[from] = b[from.Index()].Piece
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
							pieceset[*dpt] = b[(*dpt).Index()].Piece
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
					pieceset[from] = b[from.Index()].Piece
					if b[to.Index()].Piece != nil {
						pieceset[to] = b[to.Index()].Piece
					}
				ALLGUARDS:
					for {
						for gpt, _ := range board.SurroundingPoints(toPoint) {
							if gpt.Piece != nil {
								if (gpt.Orientation != board[to.Index()].Piece.Orientation) && gpt.Guards {
									set.SetPointPiece(gpt.AbsPoint, nil)
									set.SetPointPiece(to, gpt.Piece)
									pieceset[to] = board[to.Index()].Piece
									board[gpt.AbsPoint.Index()].Piece = nil
									board[to.Index()].Piece = gpt.Piece
									continue ALLGUARDS
								}
							}
						}
						break
					}
				}
				return set, pieceset
			}
		}
	}
	if b.MovesFromPoint(fromPoint, previousFrom, previousTo).Has(to) == false {
		return PointSet{}, pieceset
	}
	// en passant / in passing
	// previous points are marked as AbsPoint{0, 8} if the game has no moves made yet
	if (fromPoint.Base == Pawn) && (toPoint.Piece == nil) && (previousFrom.Rank != 8) && (previousTo.Rank != 8) {
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
					if (previousFrom == AbsPoint{File: to.File, Rank: 1}) && (previousTo == AbsPoint{File: to.File, Rank: 3}) {
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
						pieceset[to] = b[to.Index()].Piece
						return set, pieceset
					}
				} else {
					if (previousFrom == AbsPoint{File: to.File, Rank: 6}) && (previousTo == AbsPoint{File: to.File, Rank: 4}) {
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
						pieceset[to] = b[to.Index()].Piece
						return set, pieceset
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
			return set, pieceset
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
			return set, pieceset
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
			return set, pieceset
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
			return set, pieceset
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
			return set, pieceset
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
	if toPoint.Piece != nil {
		pieceset[to] = toPoint.Piece
	}
	return set, pieceset
}
