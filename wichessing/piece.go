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

const kinds = 6

const (
	King Kind = iota + 1
	Queen
	Rook
	Bishop
	Knight
	Pawn
)

func NameForKind(piece Kind) string {
	switch piece {
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

func randomKind() Kind {
	return Kind(random.Int63n(kinds) + 1)
}

func RandomPiece() Piece {
	return Piece{
		Kind: randomKind(),
	}
}
