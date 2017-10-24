// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import (
	"testing"
)

var (
	AvailableMovesCases = make([]AvailableMovesCase, 0, 1)
	AfterMoveCases      = make([]PositionAfterMoveCase, 0, 1)
)

func init() {
	for _, c := range BasicMovesCases {
		AvailableMovesCases = append(AvailableMovesCases, c)
	}
	for _, c := range BasicAfterMoveCases {
		AfterMoveCases = append(AfterMoveCases, c)
	}
}

type AvailableMovesCase struct {
	Name      string
	Position  PointSet
	Active    Orientation
	Draw      bool
	Check     bool
	Checkmate bool
	Moves     map[AbsPoint]AbsPointSet
}

// The active player is inferred by which piece is being moved. The diff is checked by location and piece kind/orientation only.
type PositionAfterMoveCase struct {
	Name    string
	Initial PointSet
	From    AbsPoint
	To      AbsPoint
	Diff    PointSet
}

// Covers Board.Draw and Board.Moves
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

// Covers Board.Move
func TestPositionAfterMoveCases(t *testing.T) {
	for _, c := range AfterMoveCases {
		b := c.Initial.Board()
		from := b[AbsPointToIndex(c.From)]
		if from.Piece == nil {
			t.Errorf("\"%v\" failed: from point %v has no piece", c.Name, c.From)
			continue
		}
		diff := b.Move(c.From, c.To, from.Orientation)
		if (len(c.Diff) == 0) && (len(diff) == 0) {
			continue
		}
		if len(c.Diff) != len(diff) {
			t.Errorf("\"%v\" failed: diff has %v changes but %v changes are expected (%v expected %v found)", c.Name, len(diff), len(c.Diff), c.Diff, diff)
			continue
		}
		// every expected point must have a matching point on the move diff
	DIFFING:
		for expected, _ := range c.Diff {
			for actual, _ := range diff {
				if (expected.File == actual.File) && (expected.Rank == actual.Rank) {
					if (expected.Piece == nil) && (actual.Piece == nil) {
						continue DIFFING
					}
					if expected.Piece == nil {
						t.Errorf("\"%v\" failed: expected no piece at %v but found %v", c.Name, expected.AbsPoint, actual.Piece)
						continue DIFFING
					}
					if actual.Piece == nil {
						t.Errorf("\"%v\" failed: expected %v at %v but found none", c.Name, expected.Piece, expected.AbsPoint)
						continue DIFFING
					}
					if expected.Orientation != actual.Orientation {
						t.Errorf("\"%v\" failed: expected %v piece but found %v piece", c.Name, expected.Orientation, actual.Orientation)
						continue DIFFING
					}
					if expected.Kind != actual.Kind {
						t.Errorf("\"%v\" failed: expected %v kind but found %v kind", c.Name, expected.Kind, actual.Kind)
						continue DIFFING
					}
					continue DIFFING
				}
			}
			t.Errorf("\"%v\" failed: found no difference at %v", c.Name, expected.AbsPoint)
		}
	}
}
