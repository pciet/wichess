// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import (
	"testing"
)

// TODO: test case generator because building these by hand is tricky and tedious

type AvailableMovesCase struct {
	Name      string
	Position  PointSet
	Active    Orientation
	Draw      bool
	Check     bool
	Checkmate bool
	Moves     map[AbsPoint]AbsPointSet
}

var AvailableMovesCases = make([]AvailableMovesCase, 0, 1)

func init() {
	for _, c := range BasicMovesCases {
		AvailableMovesCases = append(AvailableMovesCases, c)
	}
}

// Covers Board.Draw, Board.Moves, and PointSet.Board, depending on the quality of the test cases
func TestMovesCases(t *testing.T) {
CASES:
	for _, c := range AvailableMovesCases {
		b := c.Position.Board()
		draw := b.Draw(c.Active)
		if draw && (c.Draw == false) {
			t.Errorf("\"%v\" failed: unexpected draw", c.Name)
			continue
		} else if (draw == false) && c.Draw {
			t.Errorf("\"%v\" failed: determined not draw", c.Name)
			continue
		}
		moves, check, checkmate := b.Moves(c.Active)
		if check && (c.Check == false) {
			t.Errorf("\"%v\" failed: unexpected check", c.Name)
			continue
		} else if (check == false) && c.Check {
			t.Errorf("\"%v\" failed: determined not check", c.Name)
			continue
		}
		if checkmate && (c.Checkmate == false) {
			t.Errorf("\"%v\" failed: unexpected checkmate", c.Name)
			continue
		} else if (checkmate == false) && c.Checkmate {
			t.Errorf("\"%v\" failed: determined not checkmate", c.Name)
			continue
		}
		for point, targets := range moves {
			expected, has := c.Moves[point]
			if has == false {
				t.Errorf("\"%v\" failed: %v is unexpected moveable location", c.Name, point)
				continue CASES
			}
			// we're assuming board.Moves only shows points that can be moved
			if len(targets) == 0 {
				t.Errorf("\"%v\" failed: %v is marked as moveable but has no moves", c.Name, point)
				continue CASES
			}
			if targets.Equal(expected) == false {
				t.Errorf("\"%v\" failed: %v moves mismatch, %v found, %v expected, %v difference", c.Name, point, targets, expected, targets.Diff(expected))
				continue CASES
			}
		}
	}
}

//TODO: type PositionAfterMoveCase struct {
//	Name            string
//	InitialPosition PointSet
//	From            AbsPoint
//	To              AbsPoint
//	FinalPosition   PointSet
//}
