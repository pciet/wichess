package rules

type (
	// Moves are ordered in a path.
	Path struct {
		Truncated bool // MustEnd pieces need to know if the path was ended early
		Addresses []Address
	}

	RelPath []RelAddress

	// Pieces have varying paths, like the pawn with a different first and take moves.
	// This integer is an array index of a PathVariations or RelPathVariations var.
	PathVariation int
)

const (
	First PathVariation = iota
	NormalMove
	RallyMove
	Take
	PathVariationCount // not a variation
)

type (
	PathVariations    [PathVariationCount][]Path
	RelPathVariations [PathVariationCount][]RelPath
)
