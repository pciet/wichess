package rules

import (
	"log"
	"strings"
)

// TODO: is it worth caching the calculation result of applying rel paths?

// Paths relative to the piece are applied to the board in AppliedRelPaths which returns board addresses.
// Interaction of pieces is not considered. Paths that leave the board are truncated.
func AppliedRelPaths(f PieceKind, at Address, o Orientation) PathVariations {
	var out PathVariations

	rv, ok := PieceRelPaths[f]
	if ok == false {
		log.Panicln("no basic paths defined for piece", f)
	}

	offBoard := func(r RelAddress) (bool, Address) {
		x := int8(at.File)
		if o == White {
			x = x + r.X
		} else {
			x = x - r.X
		}
		if (x < 0) || (x > 7) {
			return true, Address{}
		}

		y := int8(at.Rank)
		if o == White {
			y = y + r.Y
		} else {
			y = y - r.Y
		}
		if (y < 0) || (y > 7) {
			return true, Address{}
		}

		return false, Address{uint8(x), uint8(y)}
	}

	for v, relpaths := range rv {
		paths := make([]Path, 0, len(relpaths))
		for _, rp := range relpaths {
			if len(rp) == 0 {
				log.Panicln("zero length basic path for piece", f)
			}

			off, _ := offBoard(rp[0])
			if off {
				continue
			}

			p := Path{Addresses: make([]Address, 0, len(rp))}
			for _, ra := range rp {
				off, addr := offBoard(ra)
				if off {
					p.Truncated = true
					break
				}
				p.Addresses = append(p.Addresses, addr)
			}
			paths = append(paths, p)
		}
		out[v] = paths
	}

	return out
}

// The resulting slice only has one of any path.
func CombineRelPathSlices(s ...[]RelPath) []RelPath {
	if len(s) == 0 {
		log.Panic("no slices")
	}

	z := s[0]
	for _, ps := range s[1:] {
	LOOP:
		for _, p := range ps {
			for _, ep := range z {
				if RelPathEqual(p, ep) {
					continue LOOP
				}
			}
			z = append(z, p)
		}
	}

	return z
}

func RelPathEqual(a, b RelPath) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func (a Path) String() string {
	var s strings.Builder
	for _, address := range a.Addresses {
		s.WriteString(address.String())
		s.WriteRune(' ')
	}
	if a.Truncated {
		s.WriteString("trunc")
	}
	return s.String()
}
