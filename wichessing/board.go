// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import (
	"bytes"
	"fmt"
)

type Board struct {
	Points       [64]Point
	PreviousFrom AbsPoint
	PreviousTo   AbsPoint
}

func (b Board) Copy() Board {
	var board Board
	for i, pt := range b.Points {
		board.Points[i] = Point{
			AbsPoint: AbsPoint{
				File: pt.File,
				Rank: pt.Rank,
			},
			Piece: pt.Piece.Copy(),
		}
	}
	board.PreviousFrom = b.PreviousFrom
	board.PreviousTo = b.PreviousTo
	return board
}

func (b Board) String() string {
	var buffer bytes.Buffer
	for rank := 7; rank >= 0; rank-- {
		for file := 0; file < 8; file++ {
			buffer.WriteString(fmt.Sprintf("%v ", b.Points[AbsPoint{uint8(file), uint8(rank)}.Index()]))
		}
		buffer.WriteString("\n")
	}
	buffer.WriteString(fmt.Sprintf("Previous From: %v, Previous To: %v\n", b.PreviousFrom, b.PreviousTo))
	return buffer.String()
}
