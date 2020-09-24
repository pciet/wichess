package rules

func (a *Board) pieceEnabled(at Address) bool {
	s := a[at.Index()]
	for _, p := range a.surroundingSquares(at) {
		if (p.Orientation == s.Orientation) && p.flags.enables {
			return true
		}
	}
	return false
}
