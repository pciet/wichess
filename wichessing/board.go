// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import (
	"bytes"
	"fmt"
)

type Board [64]Point

func (b Board) Copy() Board {
	var board Board
	for i, pt := range b {
		board[i] = Point{
			AbsPoint: AbsPoint{
				File: pt.File,
				Rank: pt.Rank,
			},
			Piece: pt.Piece.Copy(),
		}
	}
	return board
}

func (b Board) String() string {
	var buffer bytes.Buffer
	for rank := 7; rank >= 0; rank-- {
		for file := 0; file < 8; file++ {
			buffer.WriteString(fmt.Sprintf("%v ", b[AbsPoint{uint8(file), uint8(rank)}.Index()]))
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}
