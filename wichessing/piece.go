// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import (
	"fmt"
)

// Metadata such as number of takes, owner, identifier, or specific board location are represented in the application, not this wichessing logic that simply provides what moves can be made for a given board state.

type Piece struct {
	Kind
	Orientation
	Base     Kind
	Moved    bool  `json:"-"`
	Previous uint8 `json:"-"` // previous point

	Ghost     bool `json:"-"` // can move through other pieces
	MustEnd   bool `json:"-"` // can only move to last point in path
	MustTake  bool `json:"-"` // if taking is possible then only take moves can be made
	Swaps     bool `json:"-"` // may move to swap with friendly pieces
	Locks     bool `json:"-"` // surrounding enemy pieces cannot move
	Recons    bool `json:"-"` // friendly pieces in one of the three behind points can move to the one ahead point when empty
	Detonates bool `json:"-"` // takes all surrounding pieces when taken
	Guards    bool `json:"-"` // adjacent enemy pieces are taken by this piece
	Rallies   bool `json:"-"` // adjacent friendly pieces gain their rally moves
	Fortified bool `json:"-"` // can't be taken by pawns
}

func (p Piece) String() string {
	return fmt.Sprintf("(%v %v)", p.Kind, p.Orientation)
}

type PieceSet map[*Piece]struct{}

func (the PieceSet) Add(a Piece) PieceSet {
	the[&a] = struct{}{}
	return the
}

type Orientation int

const (
	White Orientation = 0
	Black Orientation = 1
)

func (o Orientation) String() string {
	if o == White {
		return "white"
	} else {
		return "black"
	}
}

type Kind int

// TODO: detonate doesn't take fortify

// add new kinds to the end because the client has to match this numbering
const (
	King Kind = iota + 1
	Queen
	Rook
	Bishop
	Knight
	Pawn
	SwapPawn // pawns and knights don't have a ghost kind
	LockPawn
	ReconPawn
	DetonatePawn
	GuardPawn
	RallyPawn
	FortifyPawn
	ExtendedPawn
	SwapKnight
	LockKnight
	ReconKnight
	DetonateKnight
	GuardKnight
	RallyKnight
	FortifyKnight
	ExtendedKnight
	SwapBishop
	LockBishop
	ReconBishop
	DetonateBishop
	GhostBishop
	GuardBishop
	RallyBishop
	FortifyBishop
	ExtendedBishop
	SwapRook
	LockRook
	ReconRook
	DetonateRook
	GhostRook
	GuardRook
	RallyRook
	FortifyRook
	ExtendedRook
)

func BaseForKind(the Kind) Kind {
	switch the {
	case King:
		return King
	case Queen:
		return Queen
	case Rook, SwapRook, LockRook, ReconRook, DetonateRook, GhostRook, GuardRook, RallyRook, FortifyRook, ExtendedRook:
		return Rook
	case Bishop, SwapBishop, LockBishop, ReconBishop, DetonateBishop, GhostBishop, GuardBishop, RallyBishop, FortifyBishop, ExtendedBishop:
		return Bishop
	case Knight, SwapKnight, LockKnight, ReconKnight, DetonateKnight, GuardKnight, RallyKnight, FortifyKnight, ExtendedKnight:
		return Knight
	case Pawn, SwapPawn, LockPawn, ReconPawn, DetonatePawn, GuardPawn, RallyPawn, FortifyPawn, ExtendedPawn:
		return Pawn
	default:
		panic(fmt.Sprintf("wichessing: unexpected kind %v", the))
	}
}

func (the Piece) SetKindFlags() Piece {
	the.Base = BaseForKind(the.Kind)
	switch the.Kind {
	case Knight:
		the.Ghost = true
		the.MustEnd = true
	case SwapPawn:
		the.Swaps = true
	case LockPawn:
		the.Locks = true
	case ReconPawn:
		the.Recons = true
	case DetonatePawn:
		the.Detonates = true
	case GuardPawn:
		the.Guards = true
	case RallyPawn:
		the.Rallies = true
	case FortifyPawn:
		the.Fortified = true
	case ExtendedPawn:
	case SwapKnight:
		the.Ghost = true
		the.MustEnd = true
		the.Swaps = true
	case LockKnight:
		the.Ghost = true
		the.MustEnd = true
		the.Locks = true
	case ReconKnight:
		the.Ghost = true
		the.MustEnd = true
		the.Recons = true
	case DetonateKnight:
		the.Ghost = true
		the.MustEnd = true
		the.Detonates = true
	case GuardKnight:
		the.Ghost = true
		the.MustEnd = true
		the.Guards = true
	case RallyKnight:
		the.Ghost = true
		the.MustEnd = true
		the.Rallies = true
	case FortifyKnight:
		the.Ghost = true
		the.MustEnd = true
		the.Fortified = true
	case ExtendedKnight:
		the.Ghost = true
		the.MustEnd = true
	case SwapBishop:
		the.Swaps = true
	case LockBishop:
		the.Locks = true
	case ReconBishop:
		the.Recons = true
	case DetonateBishop:
		the.Detonates = true
	case GhostBishop:
		the.Ghost = true
	case GuardBishop:
		the.Guards = true
	case RallyBishop:
		the.Rallies = true
	case FortifyBishop:
		the.Fortified = true
	case SwapRook:
		the.Swaps = true
	case LockRook:
		the.Locks = true
	case ReconRook:
		the.Recons = true
	case DetonateRook:
		the.Detonates = true
	case GhostRook:
		the.Ghost = true
	case GuardRook:
		the.Guards = true
	case RallyRook:
		the.Rallies = true
	case FortifyRook:
		the.Fortified = true
	}
	return the
}

func (the *Piece) Copy() *Piece {
	if the == nil {
		return nil
	}
	p := Piece{
		Kind:        the.Kind,
		Base:        the.Base,
		Orientation: the.Orientation,
		Moved:       the.Moved,
		Previous:    the.Previous,
	}
	p = p.SetKindFlags()
	return &p
}
