// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import (
	"fmt"
)

type Point struct {
	*Piece // nil for no piece
	AbsPoint
}

type PointSet map[*Point]struct{}

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

func (s AbsPointSet) String() map[string]struct{} {
	m := make(map[string]struct{})
	for p, _ := range s {
		m[p.String()] = struct{}{}
	}
	return m
}

func (p AbsPoint) Index() uint8 {
	return (p.File + (8 * p.Rank))
}

func FileFromIndex(i uint8) uint8 {
	return i % 8
}

func RankFromIndex(i uint8) uint8 {
	return i / 8
}

func (p AbsPoint) String() string {
	return fmt.Sprintf("%v-%v", p.File, p.Rank)
}

func IndexFromAddressString(address string) uint8 {
	var file, rank int
	_, err := fmt.Sscanf(address, "%d-%d", &file, &rank)
	if err != nil {
		panic(err.Error())
	}
	return IndexFromFileAndRank(uint8(file), uint8(rank))
}

func IndexFromFileAndRank(file, rank uint8) uint8 {
	return file + (rank * 8)
}
