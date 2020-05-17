package rules

func (a Board) IsEnPassantMove(m Move) bool {
	// this will need to be revisited if a new pawn can move to change file without taking
	return (m.From.File != m.To.File) &&
		(BasicKind(a[m.From.Index()].Kind) == Pawn) && a[m.To.Index()].Empty()
}

func (a Board) EnPassantMove(changes, takes []AddressedSquare,
	m Move) ([]AddressedSquare, []AddressedSquare) {
	s := a[m.From.Index()]
	var taking Address
	if s.Orientation == White {
		taking = Address{m.To.File, m.To.Rank - 1}
	} else {
		taking = Address{m.To.File, m.To.Rank + 1}
	}
	changes = append(changes, AddressedSquare{taking, Square{}})
	changes = append(changes, AddressedSquare{m.From, Square{}})

	return append(changes, AddressedSquare{m.To, s}),
		append(takes, AddressedSquare{taking, a[taking.Index()]})
}
