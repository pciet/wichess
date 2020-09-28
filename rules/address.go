package rules

import "strconv"

// Address is the board square address. File is the column, or X value, starting from the left,
// and rank is the row, or Y value, starting from the bottom.
//
// Address ordering is from the perspective of the white player. The left white rook is at 0,0,
// the right white rook is at 7,0, the left black rook from the black player's perspective is
// at 7,7.
type Address struct {
	File int `json:"f"`
	Rank int `json:"r"`
}

// The AddressIndex is another form of board addressing that's integers from 0 to 63, starting
// at the bottom left from the white player's perspective going to the right then moving up a row
// and restarting at the left.
type AddressIndex int

// NoAddress is the value of an Address when it doesn't point at a square on the board.
var NoAddress = Address{0, 8}

// NoAddressIndex is the value of an AddressIndex when it doesn't point at a square on the board.
const NoAddressIndex = 64

func (an Address) Index() AddressIndex   { return AddressIndex(an.File + (8 * an.Rank)) }
func (an AddressIndex) File() int        { return int(an % 8) }
func (an AddressIndex) Rank() int        { return int(an / 8) }
func (an AddressIndex) Address() Address { return Address{an.File(), an.Rank()} }
func (an AddressIndex) Int() int         { return int(an) }

// AddressSliceHasCount returns the count of the address in the slice.
func AddressSliceHasCount(a []Address, of Address) int {
	count := 0
	for _, addr := range a {
		if addr == of {
			count++
		}
	}
	return count
}

func (an Address) String() string {
	return strconv.Itoa(int(an.File)) + "-" + strconv.Itoa(int(an.Rank))
}

var (
	whiteKingStart      = Address{4, 0}
	blackKingStart      = Address{4, 7}
	whiteLeftRookStart  = Address{0, 0}
	whiteRightRookStart = Address{7, 0}
	blackLeftRookStart  = Address{7, 7}
	blackRightRookStart = Address{0, 7}
)

func removeAddressSliceDuplicates(a []Address) []Address {
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

func (an Address) squareEven() bool {
	if an.Rank%2 == 0 {
		if an.File%2 == 0 {
			return false
		} else {
			return true
		}
	} else {
		if an.File%2 == 0 {
			return true
		} else {
			return false
		}
	}
	panic("bad return")
	return true
}

func addressSliceHas(slice []Address, has Address) bool {
	for _, a := range slice {
		if a == has {
			return true
		}
	}
	return false
}
