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

type PointSet []Point

func (p PointSet) Add(the Point) PointSet {
	return append(p, the)
}

func (p PointSet) SetPointPiece(at AbsPoint, piece *Piece) PointSet {
	for i, point := range p {
		if point.AbsPoint == at {
			p[i].Piece = piece
			return p
		}
	}
	return append(p, Point{AbsPoint: at, Piece: piece})
}

func (s PointSet) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("(")
	length := len(s)
	for i, point := range s {
		buffer.WriteString(fmt.Sprintf("%v", point))
		if i != (length - 1) {
			buffer.WriteString(" ")
		}
	}
	buffer.WriteString(")")
	return buffer.String()
}

func (p PointSet) Board(previousFrom, previousTo AbsPoint) Board {
	b := Board{
		PreviousFrom: previousFrom,
		PreviousTo:   previousTo,
	}
	for i := 0; i < 64; i++ {
		b.Points[i].AbsPoint = AbsPointFromIndex(uint8(i))
	}
	for _, point := range p {
		if point.Piece != nil {
			piece := point.Piece.SetKindFlags()
			b.Points[AbsPointToIndex(point.AbsPoint)].Piece = &piece
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

type AbsPointSet []AbsPoint

func (s AbsPointSet) Add(the AbsPoint) AbsPointSet {
	return append(s, the)
}

func (s AbsPointSet) Combine(sets ...AbsPointSet) AbsPointSet {
	length := len(s)
	for _, set := range sets {
		length += len(set)
	}
	out := make(AbsPointSet, length)
	i := 0
	for _, point := range s {
		out[i] = point
		i++
	}
	for _, set := range sets {
		for _, point := range set {
			out[i] = point
			i++
		}
	}
	return out
}

func (s AbsPointSet) Reduce() AbsPointSet {
	out := make(AbsPointSet, 0, len(s))
	for _, pt := range s {
		if out.Has(pt) {
			continue
		}
		out = append(out, pt)
	}
	return s
}

func (s AbsPointSet) Has(the AbsPoint) bool {
	for _, pt := range s {
		if pt == the {
			return true
		}
	}
	return false
}

func (s AbsPointSet) Equal(an AbsPointSet) bool {
	if len(s) != len(an) {
		return false
	}
	for _, point := range s {
		if an.Has(point) == false {
			return false
		}
	}
	for _, point := range an {
		if s.Has(point) == false {
			return false
		}
	}
	return true
}

func (s AbsPointSet) ReducedDiff(from AbsPointSet) AbsPointSet {
	diff := make(AbsPointSet, 0, len(s))
	for _, point := range s {
		if from.Has(point) == false {
			diff = append(diff, point)
		}
	}
	for _, point := range from {
		if s.Has(point) == false {
			diff = append(diff, point)
		}
	}
	return diff.Reduce()
}

func (s AbsPointSet) Remove(the AbsPoint) AbsPointSet {
	out := make(AbsPointSet, 0, len(s))
	for _, point := range s {
		if point == the {
			continue
		}
		out = append(out, point)
	}
	return out
}

func (s AbsPointSet) Strings() []string {
	m := make([]string, len(s))
	for i, p := range s {
		m[i] = p.String()
	}
	return m
}

func (s AbsPointSet) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("(")
	length := len(s)
	for i, point := range s {
		buffer.WriteString(fmt.Sprintf("%v", point))
		i++
		if i != (length - 1) {
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

func AbsPointFromAddressString(address string) AbsPoint {
	var file, rank int
	_, err := fmt.Sscanf(address, "%d-%d", &file, &rank)
	if err != nil {
		panic(err.Error())
	}
	return AbsPoint{uint8(file), uint8(rank)}
}

func IndexFromFileAndRank(file, rank uint8) uint8 {
	return file + (rank * 8)
}

func AbsPointToIndex(the AbsPoint) uint8 {
	return IndexFromFileAndRank(the.File, the.Rank)
}
