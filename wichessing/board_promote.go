// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

func (b Board) AfterPromote(from AbsPoint, kind Kind) Board {
	board := b.Copy()
	diff := board.PromotePawn(from, kind)
	if (diff == nil) || (len(diff) == 0) {
		panic("failed to promote pawn")
	}
	for point, _ := range diff {
		board.Points[point.Index()] = *point
	}
	return board
}

func (b Board) HasPawnToPromote() (bool, Orientation) {
	for i := 0; i < 8; i++ {
		p := b.Points[i]
		if p.Piece != nil {
			if (p.Orientation == Black) && (p.Base == Pawn) {
				return true, Black
			}
		}
	}
	for i := 56; i < 64; i++ {
		p := b.Points[i]
		if p.Piece != nil {
			if (p.Orientation == White) && (p.Base == Pawn) {
				return true, White
			}
		}
	}
	return false, White
}

// The caller is responsible for verifying the correct orientation for the requester of the request.
func (b Board) PromotePawn(at AbsPoint, to Kind) PointSet {
	if (at.Rank != 0) && (at.Rank != 7) {
		return nil
	}
	point := b.Points[at.Index()]
	if point.Piece == nil {
		return nil
	}
	// don't affect the input
	point.Piece = point.Piece.Copy()
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
