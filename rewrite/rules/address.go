package rules

// Each square has an address.
// File is the column, or X value, starting from the left.
// Rank is the row, or Y value, starting from the bottom.
// Address ordering is from the perspective of the white player.
// The left white rook is at 0,0, the right white rook is at 7,0.
// The left black rook from that player's perspective is at 7,7.
type Address struct {
	File uint8 `json:"file"`
	Rank uint8 `json:"rank"`
}

// The File/Rank address matches an address index 0-63.
type AddressIndex uint8

func (an Address) Index() AddressIndex   { return AddressIndex(an.File + (8 * an.Rank)) }
func (an AddressIndex) File() uint8      { return uint8(a % 8) }
func (an AddressIndex) Rank() uint8      { return uint8(a / 8) }
func (an AddressIndex) Address() Address { return Address{a.File(), a.Rank()} }

// TODO: test case for SquareEven since it rarely matters

func (an Address) SquareEven() bool {
	if (an.File%2 + an.Rank%2) == 0 {
		return false
	}
	return true
}
