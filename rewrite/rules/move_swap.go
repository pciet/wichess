package rules

func (a Board) SwapMove(changes []AddressedSquare, m Move) []AddressedSquare {
	s := a[m.From.Index()]
	s.Moved = true
	changes = append(changes, AddressedSquare{m.To, s})

	s = a[m.To.Index()]
	s.Moved = true
	return append(changes, AddressedSquare{m.From, s})
}
