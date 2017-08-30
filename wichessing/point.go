// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

type Point struct {
	*Piece // nil for no piece
	AbsPoint
}

// Absolute Point represents a specific point on the board.
type AbsPoint struct {
	File uint8
	Rank uint8
}

type RelPoint struct {
	XOffset int8
	YOffset int8
}

type AbsPointSet map[*AbsPoint]struct{}

func (p AbsPoint) Index() uint8 {
	return (p.File + (8 * p.Rank))
}
