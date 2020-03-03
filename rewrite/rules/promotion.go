package rules

func (a Board) PromotionNeeded() bool {
	for i := 0; i < 8; i++ {
		p := a[i]
		if (p.BasicKind() == Pawn) && (p.Orientation == Black) {
			return true
		}
	}
	for i := 56; i < 64; i++ {
		p := a[i]
		if (p.BasicKind() == Pawn) && (p.Orientation == White) {
			return true
		}
	}
	return false
}
