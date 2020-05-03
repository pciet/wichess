package rules

const (
	NoKind PieceKind = iota
	King
	Queen
	Rook
	Bishop
	// knights are quick (ghosts)
	Knight
	Pawn
	// pawn that neutralizes (detonates) but can only move
	// one square forward for its first move
	War
	// pawn that reveals (recon) and enables (rallies)
	Form
	// knight that asserts (guards)
	Constructive
	PieceKindCount // add new piece kinds between this and the previous
)

const BasicPieceKindCount = 6
