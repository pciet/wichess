// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import (
	"bytes"
	"fmt"
)

type Point struct {
	*Piece // nil for no piece
	AbsPoint
}

func (p Point) String() string {
	if p.Piece == nil {
		return p.AbsPoint.String()
	} else {
		return fmt.Sprintf("%v:%v", p.AbsPoint, p.Piece)
	}
}

type PointSet map[*Point]struct{}

func (s PointSet) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("(")
	length := len(s)
	i := 0
	for point, _ := range s {
		buffer.WriteString(fmt.Sprintf("%v", point))
		i++
		if i != length {
			buffer.WriteString(" ")
		}
	}
	buffer.WriteString(")")
	return buffer.String()
}

func (p PointSet) Board() Board {
	b := Board{}
	for i := 0; i < 64; i++ {
		b[i].AbsPoint = AbsPointFromIndex(uint8(i))
	}
	for point, _ := range p {
		if point.Piece != nil {
			piece := (*(point.Piece)).SetKindFlags()
			b[AbsPointToIndex(point.AbsPoint)].Piece = &piece
		}
	}
	return b
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

func (s AbsPointSet) Add(the AbsPointSet) AbsPointSet {
	if len(the) == 0 {
		return s
	}
	newset := make(AbsPointSet)
	for pt, _ := range s {
		newset[&AbsPoint{
			File: pt.File,
			Rank: pt.Rank,
		}] = struct{}{}
	}
OUTER:
	for pt, _ := range the {
		for ep, _ := range newset {
			if (pt.File == ep.File) && (pt.Rank == ep.Rank) {
				continue OUTER
			}
		}
		newset[&AbsPoint{
			File: pt.File,
			Rank: pt.Rank,
		}] = struct{}{}
	}
	return newset
}

func (s AbsPointSet) Reduce() AbsPointSet {
	for pt, _ := range s {
		for opt, _ := range s {
			if pt == opt {
				continue
			}
			if (pt.File == opt.File) && (pt.Rank == opt.Rank) {
				delete(s, opt)
			}
		}
	}
	return s
}

func (s AbsPointSet) Has(the AbsPoint) bool {
	for pt, _ := range s {
		if (pt.File == the.File) && (pt.Rank == the.Rank) {
			return true
		}
	}
	return false
}

func (s AbsPointSet) Equal(an AbsPointSet) bool {
	for point, _ := range s {
		if an.Has(*point) == false {
			return false
		}
	}
	for point, _ := range an {
		if s.Has(*point) == false {
			return false
		}
	}
	return true
}

func (s AbsPointSet) Diff(from AbsPointSet) AbsPointSet {
	diff := make(AbsPointSet)
	for point, _ := range s {
		if from.Has(*point) == false {
			diff[point] = struct{}{}
		}
	}
	for point, _ := range from {
		if s.Has(*point) == false {
			diff[point] = struct{}{}
		}
	}
	return diff.Reduce()
}

func (s AbsPointSet) Strings() map[string]struct{} {
	m := make(map[string]struct{})
	for p, _ := range s {
		m[p.String()] = struct{}{}
	}
	return m
}

func (s AbsPointSet) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("(")
	length := len(s)
	i := 0
	for point, _ := range s {
		buffer.WriteString(fmt.Sprintf("%v", point))
		i++
		if i != length {
			buffer.WriteString(" ")
		}
	}
	buffer.WriteString(")")
	return buffer.String()
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

func AbsPointFromIndex(i uint8) AbsPoint {
	return AbsPoint{
		File: FileFromIndex(i),
		Rank: RankFromIndex(i),
	}
}

func (p AbsPoint) String() string {
	return fmt.Sprintf("%v-%v", p.File, p.Rank)
}

func (p RelPoint) String() string {
	return fmt.Sprintf("%v:%v", p.XOffset, p.YOffset)
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

func AbsPointToIndex(the AbsPoint) uint8 {
	return IndexFromFileAndRank(the.File, the.Rank)
}
