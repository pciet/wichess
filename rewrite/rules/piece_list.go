package rules

const (
	NoKind PieceKind = iota
	King
	Queen
	Rook
	Bishop
	Knight
	Pawn
	SwapPawn // pawns and knights don't have a ghost kind
	LockPawn
	ReconPawn
	DetonatePawn
	GuardPawn
	RallyPawn
	FortifyPawn
	ExtendedPawn
	SwapKnight
	LockKnight
	ReconKnight
	DetonateKnight
	GuardKnight
	RallyKnight
	FortifyKnight
	ExtendedKnight
	SwapBishop
	LockBishop
	ReconBishop
	DetonateBishop
	GhostBishop
	GuardBishop
	RallyBishop
	FortifyBishop
	ExtendedBishop
	SwapRook
	LockRook
	ReconRook
	DetonateRook
	GhostRook
	GuardRook
	RallyRook
	FortifyRook
	ExtendedRook
	PieceKindCount // add new piece kinds between this and the previous
)

const BasicPieceKindCount = 6
