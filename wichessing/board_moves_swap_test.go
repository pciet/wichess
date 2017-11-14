// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

var SwapMovesCases = []AvailableMovesCase{
	{
		Name:   "Starting Right White Swap Knight",
		Active: White,
		Position: PointSet{
			&WhiteKingStart:        {},
			&WhiteRightBishopStart: {},
			&WhiteRightRookStart:   {},
			{
				AbsPoint: AbsPoint{6, 0},
				Piece: &Piece{
					Kind:        SwapKnight,
					Orientation: White,
				},
			}: {},
			&WhitePawn4Start: {},
			&WhitePawn5Start: {},
			&WhitePawn6Start: {},
			&WhitePawn7Start: {},
			&BlackKingStart:  {},
		},
		PreviousFrom: AbsPoint{3, 3},
		PreviousTo:   AbsPoint{3, 4},
		Moves: map[AbsPoint]AbsPointSet{
			{4, 0}: {
				{3, 0}: {},
				{3, 1}: {},
			},
			{6, 0}: {
				{4, 1}: {},
				{7, 2}: {},
				{5, 2}: {},
			},
			{4, 1}: {
				{4, 2}: {},
				{4, 3}: {},
			},
			{5, 1}: {
				{5, 2}: {},
				{5, 3}: {},
			},
			{6, 1}: {
				{6, 2}: {},
				{6, 3}: {},
			},
			{7, 1}: {
				{7, 2}: {},
				{7, 3}: {},
			},
			{4, 7}: {
				{3, 7}: {},
				{3, 6}: {},
				{4, 6}: {},
				{5, 6}: {},
				{5, 7}: {},
			},
		},
	},
}
