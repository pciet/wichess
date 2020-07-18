package rules

func (a Board) PieceRallied(at Address) bool {
	s := a[at.Index()]
	for _, p := range a.SurroundingSquares(at) {
		if (p.Square.Orientation == s.Orientation) && p.Square.Rallies {
			return true
		}
	}
	return false
}
