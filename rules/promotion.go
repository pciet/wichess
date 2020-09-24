package rules

import "github.com/pciet/wichess/piece"

// PromotionNeeded returns whether a promotion must be done, and if so it returns the Orientation
// of the player that needs to promote.
func (a *Board) PromotionNeeded() (Orientation, bool) {
	for i := 0; i < 8; i++ {
		p := a[i]
		if (p.Kind.Basic() == piece.Pawn) && (p.Orientation == Black) {
			return Black, true
		}
	}
	for i := 56; i < 64; i++ {
		p := a[i]
		if (p.Kind.Basic() == piece.Pawn) && (p.Orientation == White) {
			return White, true
		}
	}
	return White, false
}

// DoPromotion finds the piece that needs to be promoted and returns the value of its Square after
// the promotion is done. The Board is not changed.
func (a *Board) DoPromotion(with piece.Kind) Square {
	for i := 0; i < 8; i++ {
		s := a[i]
		if (s.Kind.Basic() == piece.Pawn) && (s.Orientation == Black) {
			s.Kind = with
			return Square{
				Address: AddressIndex(i).Address(),
				Piece:   s,
			}
		}
	}
	for i := 56; i < 64; i++ {
		s := a[i]
		if (s.Kind.Basic() == piece.Pawn) && (s.Orientation == White) {
			s.Kind = with
			return Square{
				Address: AddressIndex(i).Address(),
				Piece:   s,
			}
		}
	}
	panic("didn't find promotion")
	return Square{}
}
