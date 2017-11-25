// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"fmt"

	"github.com/pciet/wichess/wichessing"
)

func nameForKind(the wichessing.Kind) string {
	switch the {
	case wichessing.King:
		return "King"
	case wichessing.Queen:
		return "Queen"
	case wichessing.Rook:
		return "Rook"
	case wichessing.Bishop:
		return "Bishop"
	case wichessing.Knight:
		return "Knight"
	case wichessing.Pawn:
		return "Pawn"
	case wichessing.SwapPawn:
		return "Swap Pawn"
	case wichessing.LockPawn:
		return "Lock Pawn"
	case wichessing.ReconPawn:
		return "Recon Pawn"
	case wichessing.DetonatePawn:
		return "Detonate Pawn"
	case wichessing.GuardPawn:
		return "Guard Pawn"
	case wichessing.RallyPawn:
		return "Rally Pawn"
	case wichessing.FortifyPawn:
		return "Fortify Pawn"
	case wichessing.ExtendedPawn:
		return "Extended Pawn"
	case wichessing.SwapKnight:
		return "Swap Knight"
	case wichessing.LockKnight:
		return "Lock Knight"
	case wichessing.ReconKnight:
		return "Recon Knight"
	case wichessing.DetonateKnight:
		return "Detonate Knight"
	case wichessing.GuardKnight:
		return "Guard Knight"
	case wichessing.RallyKnight:
		return "Rally Knight"
	case wichessing.FortifyKnight:
		return "Fortify Knight"
	case wichessing.ExtendedKnight:
		return "Extended Knight"
	case wichessing.SwapBishop:
		return "Swap Bishop"
	case wichessing.LockBishop:
		return "Lock Bishop"
	case wichessing.ReconBishop:
		return "Recon Bishop"
	case wichessing.DetonateBishop:
		return "Detonate Bishop"
	case wichessing.GhostBishop:
		return "Ghost Bishop"
	case wichessing.GuardBishop:
		return "Guard Bishop"
	case wichessing.RallyBishop:
		return "Rally Bishop"
	case wichessing.FortifyBishop:
		return "Fortify Bishop"
	case wichessing.ExtendedBishop:
		return "Extended Bishop"
	case wichessing.SwapRook:
		return "Swap Rook"
	case wichessing.LockRook:
		return "Lock Rook"
	case wichessing.ReconRook:
		return "Recon Rook"
	case wichessing.DetonateRook:
		return "Detonate Rook"
	case wichessing.GhostRook:
		return "Ghost Rook"
	case wichessing.GuardRook:
		return "Guard Rook"
	case wichessing.RallyRook:
		return "Rally Rook"
	case wichessing.FortifyRook:
		return "Fortify Rook"
	case wichessing.ExtendedRook:
		return "Extended Rook"
	}
	if debug {
		fmt.Println("nameForKind: no name found")
	}
	return ""
}
