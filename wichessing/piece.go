// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import (
	"crypto/rand"
	"math"
	"math/big"
	prand "math/rand"
)

type Piece struct {
	Kind
	Takes      int
	Identifier int
}

func CopyFromPiece(from Piece, to *Piece) {
	to.Kind = from.Kind
	to.Takes = from.Takes
}

type Kind int

const kinds = 27

const (
	King Kind = iota + 1
	Queen
	Rook
	Bishop
	Knight
	Pawn
	WhiteKing
	WhiteQueen
	WhiteRook
	WhiteBishop
	WhiteKnight
	WhitePawn
	BlackKing
	BlackQueen
	BlackRook
	BlackBishop
	BlackKnight
	BlackPawn
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

func NameForKind(k Kind) string {
	switch k {
	case King:
		return "King"
	case Queen:
		return "Queen"
	case Rook:
		return "Rook"
	case Bishop:
		return "Bishop"
	case Knight:
		return "Knight"
	case Pawn:
		return "Pawn"
	case Swap:
		return "Swap"
	case Lock:
		return "Lock"
	case Recon:
		return "Recon"
	case Detonate:
		return "Detonate"
	case Ghost:
		return "Ghost"
	case Steal:
		return "Steal"
	case Guard:
		return "Guard"
	case Rally:
		return "Rally"
	case Fortify:
		return "Fortify"
	}
	return ""
}

var random *prand.Rand

func init() {
	seed, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		panic(err.Error())
	}
	random = prand.New(prand.NewSource(seed.Int64()))
}

func randomHeroKind() Kind {
	return Kind(random.Int63n(9) + 1 + 18)
}

func RandomPiece() Piece {
	return Piece{
		Kind: randomHeroKind(),
	}
}
