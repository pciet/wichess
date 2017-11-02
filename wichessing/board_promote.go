// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

func (b Board) AfterPromote(from AbsPoint, kind Kind) Board {
	board := b.Copy()
	for point, _ := range board.PromotePawn(from, kind) {
		board[point.Index()] = *point
	}
	return board
}

func (b Board) HasPawnToPromote() bool {
	for i := 0; i < 8; i++ {
		p := b[i]
		if p.Piece != nil {
			if (p.Orientation == Black) && (p.Base == Pawn) {
				return true
			}
		}
	}
	for i := 56; i < 64; i++ {
		p := b[i]
		if p.Piece != nil {
			if (p.Orientation == White) && (p.Base == Pawn) {
				return true
			}
		}
	}
	return false
}

// The caller is responsible for verifying the correct orientation for the requester of the request.
func (b Board) PromotePawn(at AbsPoint, to Kind) PointSet {
	if (at.Rank != 0) && (at.Rank != 7) {
		return nil
	}
	point := b[at.Index()]
	if point.Piece == nil {
		return nil
	}
	if point.Base != Pawn {
		return nil
	}
	if (at.Rank == 0) && (point.Orientation != Black) {
		return nil
	}
	if (at.Rank == 7) && (point.Orientation != White) {
		return nil
	}
	if (to != Knight) && (to != Bishop) && (to != Rook) && (to != Queen) {
		return nil
	}
	point.Kind = to
	point.Base = BaseForKind(to)
	set := make(PointSet)
	set[&point] = struct{}{}
	return set
}
