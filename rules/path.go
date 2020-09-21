package rules

import (
	"strings"

	"github.com/pciet/wichess/piece"
)

type (
	path struct {
		Truncated bool // MustEnd pieces need to know if the path was ended early
		Addresses []Address
	}

	pathVariations [piece.PathVariationCount][]Path
)

// TODO: is it worth caching the calculation result of applying paths?

// appliedPaths changes the piece.Address that's relative to the piece into absolute board
// addresses. Other pieces aren't considered. If a piece's path leaves the board then the
// path is truncated at the last square on the board edge.
func appliedPaths(f piece.Kind, at Address, o Orientation) PathVariations {
	var out PathVariations
	rv := piece.Paths(f)

	offBoard := func(r piece.Address) (bool, Address) {
		x := at.File
		if o == White {
			x = x + r.File
		} else {
			x = x - r.File
		}
		if (x < 0) || (x > 7) {
			return true, Address{}
		}

		y := at.Rank
		if o == White {
			y = y + r.Rank
		} else {
			y = y - r.Rank
		}
		if (y < 0) || (y > 7) {
			return true, Address{}
		}

		return false, Address{x, y}
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
