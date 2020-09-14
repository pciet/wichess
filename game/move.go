package game

import (
	"fmt"

	"github.com/pciet/wichess/piece"
	"github.com/pciet/wichess/rules"
)

// Move attempts to do the requested move onto the Instance. The changed squares, captured pieces,
// and whether a promotion is now needed are returned. If the move couldn't be done then the
// changed squares slice is nil.
func (an Instance) Move(with rules.Move) ([]rules.AddressedSquare, []rules.Piece, bool) {
	if an.moveLegal(with) == false {
		return nil, nil, false
	}

	changes, captures := &(an.Board).DoMove(with)
	if changes == nil {
		return nil, nil, false
	}

	// clear moves cache
	an.moves = nil

	an.PreviousMove = with

	// determine draw turn count which is reset by pawn move or capture
	an.DrawTurns++
	if len(captures) > 0 {
		an.DrawTurns = 0
	} else {
		for _, change := range changes {
			if change.Kind.Basic() == piece.Pawn {
				an.DrawTurns = 0
				break
			}
		}
	}

	// record captures to game memory
	wCaptIndex := memory.FirstAvailableCaptureIndex(&(an.White.Captures))
	bCaptIndex := memory.FirstAvailableCaptureIndex(&(an.Black.Captures))
	for _, capt := range captures {
		if capt.Kind == piece.NoKind {
			panic("capture list with piece.NoKind")
		}
		if capt.Orientation == rules.White {
			an.White.Captures[wCaptIndex] = capt.Kind
			wCaptIndex++
		} else if capt.Orientation == rules.Black {
			an.Black.Captures[bCaptIndex] = capt.Kind
			bCaptIndex++
		} else {
			panic(fmt.Sprint("unknown orientation %v", capt.Kind))
		}
	}

	// set active and previous active player
	an.PreviousActive = an.Active
	promoter, promotionNeeded := an.PromotionNeeded()
	if promotionNeeded {
		if promoter != an.Active {
			// reverse promotion
			an.Active = promoter
		}
		// else active player stays the same
	} else {
		an.Active = an.OpponentOf(an.Active)
	}

	an.Changed()

	return changes, captures, promotionNeeded
}

func (an Instance) moveLegal(m rules.Move) bool {
	from := an.Board[m.From.Index()]
	if (from.Kind == piece.NoKind) || (from.Orientation != an.Active) {
		return false
	}

	// TODO: cache move calculations

	moves, state := an.Moves()
	if (state != rules.Normal) && (state != rules.Check) {
		return false
	}

	if rules.MoveSetSliceHasMove(moves, m) == false {
		return false
	}

	return true
}
