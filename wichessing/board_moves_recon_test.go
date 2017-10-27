// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

var ReconMovesCases = []AvailableMovesCase{
	{
		Name:   "Recon Out Of King Range",
		Active: White,
		Position: PointSet{
			&WhiteKingStart: {},
			{
				AbsPoint: AbsPoint{2, 1},
				Piece: &Piece{
					Kind:        ReconPawn,
					Orientation: White,
				},
			}: {},
			{
				AbsPoint: AbsPoint{6, 1},
				Piece: &Piece{
					Kind:        ReconPawn,
					Orientation: White,
				},
			}: {},
			&BlackKingStart: {},
			{
				AbsPoint: AbsPoint{2, 6},
				Piece: &Piece{
					Kind:        ReconPawn,
					Orientation: Black,
				},
			}: {},
			{
				AbsPoint: AbsPoint{6, 6},
				Piece: &Piece{
					Kind:        ReconPawn,
					Orientation: Black,
				},
			}: {},
		},
		Moves: map[AbsPoint]AbsPointSet{
			{4, 0}: {
				{3, 0}: {},
				{3, 1}: {},
				{4, 1}: {},
				{5, 1}: {},
				{5, 0}: {},
			},
			{2, 1}: {
				{2, 2}: {},
				{2, 3}: {},
			},
			{6, 1}: {
				{6, 2}: {},
				{6, 3}: {},
			},
			{4, 7}: {
				{3, 7}: {},
				{3, 6}: {},
				{4, 6}: {},
				{5, 6}: {},
				{5, 7}: {},
			},
			{2, 6}: {
				{2, 5}: {},
				{2, 4}: {},
			},
			{6, 6}: {
				{6, 5}: {},
				{6, 4}: {},
			},
		},
	},
	{
		Name:   "White King Recon Rook Right, Black King Recon Bishop Right",
		Active: White,
		Position: PointSet{
			&WhiteKingStart: {},
			{
				AbsPoint: AbsPoint{5, 1},
				Piece: &Piece{
					Kind:        ReconRook,
					Orientation: White,
				},
			}: {},
			&BlackKingStart: {},
			{
				AbsPoint: AbsPoint{5, 6},
				Piece: &Piece{
					Kind:        ReconBishop,
					Orientation: Black,
				},
			}: {},
		},
		Moves: map[AbsPoint]AbsPointSet{
			{4, 0}: {
				{3, 0}: {},
				{3, 1}: {},
				{4, 1}: {},
				{5, 2}: {},
				{5, 0}: {},
			},
			{5, 1}: {
				{5, 0}: {},
				{5, 2}: {},
				{5, 3}: {},
				{5, 4}: {},
				{5, 5}: {},
				{5, 6}: {},
				{4, 1}: {},
				{3, 1}: {},
				{2, 1}: {},
				{1, 1}: {},
				{0, 1}: {},
				{6, 1}: {},
				{7, 1}: {},
			},
			{4, 7}: {
				{3, 7}: {},
				{3, 6}: {},
				{4, 6}: {},
				// Black King cannot move to 5-5 because the White Recon Rook threatens that square
				{5, 7}: {},
			},
			{5, 6}: {
				{6, 7}: {},
				{4, 5}: {},
				{3, 4}: {},
				{2, 3}: {},
				{1, 2}: {},
				{0, 1}: {},
				{6, 5}: {},
				{7, 4}: {},
			},
		},
	},
	{
		Name:   "King Recon Pawn Left",
		Active: White,
		Position: PointSet{
			&WhiteKingStart: {},
			{
				AbsPoint: AbsPoint{3, 1},
				Piece: &Piece{
					Kind:        ReconPawn,
					Orientation: White,
				},
			}: {},
			&BlackKingStart: {},
			{
				AbsPoint: AbsPoint{3, 6},
				Piece: &Piece{
					Kind:        ReconPawn,
					Orientation: Black,
				},
			}: {},
		},
		Moves: map[AbsPoint]AbsPointSet{
			{4, 0}: {
				{3, 0}: {},
				{4, 1}: {},
				{3, 2}: {},
				{5, 1}: {},
				{5, 0}: {},
			},
			{3, 1}: {
				{3, 2}: {},
				{3, 3}: {},
			},
			{4, 7}: {
				{3, 7}: {},
				{3, 5}: {},
				{4, 6}: {},
				{5, 6}: {},
				{5, 7}: {},
			},
			{3, 6}: {
				{3, 5}: {},
				{3, 4}: {},
			},
		},
	},
	{
		Name:   "King Recon Pawn Forward",
		Active: White,
		Position: PointSet{
			&WhiteKingStart: {},
			{
				AbsPoint: AbsPoint{4, 1},
				Piece: &Piece{
					Kind:        ReconPawn,
					Orientation: White,
				},
			}: {},
			&BlackKingStart: {},
			{
				AbsPoint: AbsPoint{4, 6},
				Piece: &Piece{
					Kind:        ReconPawn,
					Orientation: Black,
				},
			}: {},
		},
		Moves: map[AbsPoint]AbsPointSet{
			{4, 0}: {
				{3, 0}: {},
				{3, 1}: {},
				{4, 2}: {},
				{5, 1}: {},
				{5, 0}: {},
			},
			{4, 1}: {
				{4, 2}: {},
				{4, 3}: {},
			},
			{4, 7}: {
				{3, 7}: {},
				{3, 6}: {},
				{4, 5}: {},
				{5, 6}: {},
				{5, 7}: {},
			},
			{4, 6}: {
				{4, 5}: {},
				{4, 4}: {},
			},
		},
	},
}
