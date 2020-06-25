package rules

import "github.com/pciet/wichess/piece"

func (a Board) PromotionNeeded() (Orientation, bool) {
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

func (a Board) DoPromotion(with piece.Kind) AddressedSquare {
	for i := 0; i < 8; i++ {
		s := a[i]
		if (s.Kind.Basic() == piece.Pawn) && (s.Orientation == Black) {
			s.Kind = with
			return AddressedSquare{
				Address: AddressIndex(i).Address(),
				Square:  s,
			}
		}
	}
	for i := 56; i < 64; i++ {
		s := a[i]
		if (s.Kind.Basic() == piece.Pawn) && (s.Orientation == White) {
			s.Kind = with
			return AddressedSquare{
				Address: AddressIndex(i).Address(),
				Square:  s,
			}
		}
	}
	Panic("didn't find promotion")
	return AddressedSquare{}
}
