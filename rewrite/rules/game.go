package rules

// A square is one point on the grid board.
// If the square's PieceKind is 0 then there's no piece.
type Square Piece

// A Wisconsin Chess board is a regular 8x8 chess board.
type Board [8 * 8]Square

type BoardAddress struct {
	File uint8
	Rank uint8
}

// A BoardAddress can be represented as an index 0-63.
// Bottom left corner square is 0, count to the right, then continue at far left square of next rank.
type BoardAddressIndex uint8

func (a BoardAddress) Index() BoardAddressIndex   { return BoardAddressIndex(a.File + (8 * a.Rank)) }
func (a BoardAddressIndex) File() uint8           { return uint8(a % 8) }
func (a BoardAddressIndex) Rank() uint8           { return uint8(a / 8) }
func (a BoardAddressIndex) Address() BoardAddress { return BoardAddress{a.File(), a.Rank()} }

type Move struct {
	From BoardAddress
	To   BoardAddress
}

type Game struct {
	Board
	Previous Move // The en passant take option depends on the previous move.
}
