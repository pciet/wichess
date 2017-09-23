// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

// Metadata such as number of takes, owner, identifier, or specific board location are represented in the application, not this wichessing logic that simply provides what moves can be made for a given board state.

type Piece struct {
	Kind
	Orientation
	Moved     bool
	Ghost     bool `json:"-"`  // can move through other pieces
	MustEnd   bool `json:"-"`  // can only move to last point in path
	MustTake  bool `json:"-"`  // if taking is possible then only take moves can be made
	Swaps     bool `json:"-"`  // may move to swap with friendly pieces
	Locks     bool `json:"-"`  // surrounding enemy pieces cannot move
	Recons    bool `json:"-"`  // friendly pieces in one of the three behind points can move to the one ahead point when empty
	Detonates bool `json:"-"`  // takes all surrounding pieces when taken
	Steals    bool `json:"-"`  // instead of taking convert the other piece
	Guards    bool `json:"-"`  // adjacent enemy pieces are taken by this piece
	Rallies   bool `'json:"-"` // adjacent friendly pieces gain their rally moves
	Fortified bool `'json:"-"` // can't be taken by pawns
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
	Swap  // can switch with friendly pieces by normal moves
	Lock  // surrounding enemy pieces cannot move
	Recon // pieces can move from one of the three behind points to the empty one ahead of this piece
	// bishop kinds
	Detonate // takes all surrounding pieces when taken, friend and enemy
	Ghost    // can move through other pieces
	Steal    // instead of taking, this piece converts the other piece and moves adjacent
	// rook kinds
	Guard   // can only move one but any adjacent enemy is taken
	Rally   // adjacent friendly pieces gain additional rally move paths
	Fortify // can't be taken by pawns
)

func (the Piece) SetKindFlags() Piece {
	switch the.Kind {
	case Knight:
		the.Ghost = true
		the.MustEnd = true
	case Swap:
		the.Ghost = true
		the.MustEnd = true
		the.Swaps = true
	case Lock:
		the.Ghost = true
		the.MustEnd = true
		the.Locks = true
	case Recon:
		the.Ghost = true
		the.MustEnd = true
		the.Recons = true
	case Detonate:
		the.Detonates = true
	case Ghost:
		the.Ghost = true
	case Steal:
		the.Steals = true
	case Guard:
		the.Guards = true
	case Rally:
		the.Rallies = true
	case Fortify:
		the.Fortified = true
	case Pawn:
		the.MustTake = true
	}
	return the
}

func (the *Piece) Copy() *Piece {
	if the == nil {
		return nil
	}
	p := Piece{
		Kind:        the.Kind,
		Orientation: the.Orientation,
		Moved:       the.Moved,
	}
	p = p.SetKindFlags()
	return &p
}
