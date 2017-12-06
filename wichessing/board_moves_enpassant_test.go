// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

var EnPassantMovesCases = []AvailableMovesCase{
	{
		Name:   "No Passant Available, Black Double Take",
		Active: Black,
		Position: PointSet{
			{
				AbsPoint: AbsPoint{3, 0},
				Piece: &Piece{
					Kind:        King,
					Orientation: White,
					Moved:       true,
				},
			},
			BlackKingStart,
			{
				AbsPoint: AbsPoint{5, 3},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{4, 3},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: Black,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{6, 3},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: Black,
					Moved:       true,
				},
			},
		},
		PreviousFrom: AbsPoint{4, 0},
		PreviousTo:   AbsPoint{3, 0},
		Moves: map[AbsPoint]AbsPointSet{
			{3, 0}: {
				{2, 0},
				{2, 1},
				{3, 1},
				{4, 1},
				{4, 0},
			},
			{4, 7}: {
				{3, 7},
				{3, 6},
				{4, 6},
				{5, 6},
				{5, 7},
			},
			{5, 3}: {
				{5, 4},
			},
			{4, 3}: {
				{4, 2},
			},
			{6, 3}: {
				{6, 2},
			},
		},
	},
	{
		Name:   "Black Double En Passant Take",
		Active: Black,
		Position: PointSet{
			WhiteKingStart,
			BlackKingStart,
			{
				AbsPoint: AbsPoint{5, 3},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: White,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{4, 3},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: Black,
					Moved:       true,
				},
			},
			{
				AbsPoint: AbsPoint{6, 3},
				Piece: &Piece{
					Kind:        Pawn,
					Orientation: Black,
					Moved:       true,
				},
			},
		},
		PreviousFrom: AbsPoint{5, 1},
		PreviousTo:   AbsPoint{5, 3},
		Moves: map[AbsPoint]AbsPointSet{
			{4, 0}: {
				{3, 0},
				{3, 1},
				{4, 1},
				{5, 1},
				{5, 0},
			},
			{4, 7}: {
				{3, 7},
				{3, 6},
				{4, 6},
				{5, 6},
				{5, 7},
			},
			{5, 3}: {
				{5, 4},
			},
			{4, 3}: {
				{5, 2},
				{4, 2},
			},
			{6, 3}: {
				{5, 2},
				{6, 2},
			},
		},
	},
}
