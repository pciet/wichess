package rules

import "strconv"

// Each square has an address.
// File is the column, or X value, starting from the left.
// Rank is the row, or Y value, starting from the bottom.
// Address ordering is from the perspective of the white player.
// The left white rook is at 0,0, the right white rook is at 7,0.
// The left black rook from that player's perspective is at 7,7.
type Address struct {
	File uint8 `json:"f"`
	Rank uint8 `json:"r"`
}

// Relative addressing is from the perspective of a square or piece instead of from the board.
// This kind of addressing is necessary to generically define what moves pieces can make.
type RelAddress struct {
	X int8
	Y int8
}

// The File/Rank address matches an address index 0-63.
type AddressIndex uint8

func (an Address) Index() AddressIndex   { return AddressIndex(an.File + (8 * an.Rank)) }
func (an AddressIndex) File() uint8      { return uint8(an % 8) }
func (an AddressIndex) Rank() uint8      { return uint8(an / 8) }
func (an AddressIndex) Address() Address { return Address{an.File(), an.Rank()} }
func (an AddressIndex) Int() int         { return int(an) }

var NoAddress = Address{0, 8}

func RemoveAddressSliceDuplicates(a []Address) []Address {
	out := make([]Address, 0, len(a))
LOOP:
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				continue LOOP
			}
		}
		out = append(out, a[i])
	}
	return out
}

// TODO: test case for SquareEven since it rarely matters

func (an Address) SquareEven() bool {
	if (an.File%2 + an.Rank%2) == 0 {
		return false
	}
	return true
}

func (an Address) String() string {
	return strconv.Itoa(int(an.File)) + "-" + strconv.Itoa(int(an.Rank))
}
