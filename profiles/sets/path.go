// Copyright 2017 Matthew Juran
// All Rights Reserved

// Profiling set implementations for a set of grid coordinate slices representing paths on a small grid.
package main

import ()

type Point struct {
	X uint8
	Y uint8
}

func (the Point) Equal(other Point) bool {
	if (the.X == other.X) && (the.Y == other.Y) {
		return true
	}
	return false
}

type Path []Point

func (the Path) Equal(other Path) bool {
	if len(the) != len(other) {
		return false
	}
	for i, point := range the {
		if point.Equal(other[i]) == false {
			return false
		}
	}
	return true
}
