package rules

func (a Board) PromotionNeeded() (Orientation, bool) {
	for i := 0; i < 8; i++ {
		p := a[i]
		if (BasicKind(p.Kind) == Pawn) && (p.Orientation == Black) {
			return Black, true
		}
	}
	for i := 56; i < 64; i++ {
		p := a[i]
		if (BasicKind(p.Kind) == Pawn) && (p.Orientation == White) {
			return White, true
		}
	}
	return White, false
}

func (a Board) DoPromotion(with PieceKind) AddressedSquare {
	for i := 0; i < 8; i++ {
		s := a[i]
		if (BasicKind(s.Kind) == Pawn) && (s.Orientation == Black) {
			s.Kind = with
			return AddressedSquare{
				Address: AddressIndex(i).Address(),
				Square:  s,
			}
		}
	}
	for i := 56; i < 64; i++ {
		s := a[i]
		if (BasicKind(s.Kind) == Pawn) && (s.Orientation == White) {
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
