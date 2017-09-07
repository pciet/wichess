// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

// Metadata such as number of takes, owner, identifier, or specific board location are represented in the application, not this wichessing logic that simply provides what moves can be made for a given board state.

type Piece struct {
	Kind
	Orientation
	Moved   bool
	Ghost   bool `json:"-"` // can move through other pieces
	MustEnd bool `json:"-"` // can only move to last point in path
}

type Orientation int

const (
	White Orientation = 0
	Black Orientation = 1
)

type Kind int

const (
	King Kind = iota + 1
	Queen
	Rook
	Bishop
	Knight
	Pawn
	// knight kinds
	Swap
	Lock
	Recon
	// bishop kinds
	Detonate
	Ghost
	Steal
	// rook kinds
	Guard
	Rally
	Fortify
)

func (the Piece) SetKindFlags() Piece {
	if the.Kind == Knight {
		the.Ghost = true
		the.MustEnd = true
	}
	return the
}
