package game

import (
	"log"

	"github.com/pciet/wichess/piece"
	"github.com/pciet/wichess/rules"
)

// Move attempts to do the requested move onto the Instance. The changed squares, captured pieces,
// and whether a promotion is now needed are returned. If the move couldn't be done then the
// changed squares slice is nil.
func (an Instance) Move(with rules.Move) ([]rules.Square, []rules.Piece, bool) {
	if an.moveLegal(with) == false {
		return nil, nil, false
	}

	changes, captures := an.Game.Board.DoMove(with)
	if changes == nil {
		return nil, nil, false
	}
	for _, capt := range captures {
		if capt.Piece.Kind != piece.King {
			continue
		}
		log.Panicln("captured king in move", with, "\nchanges", changes, "\ncaptures", captures,
			"\n", an.String())
	}

	(&an.Game.Board).ApplyChanges(changes)

	// clear moves cache
	an.MovesCache = nil
	an.StateCache = rules.NoState

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
	wCaptIndex := an.White.Captures.FirstAvailableIndex()
	bCaptIndex := an.Black.Captures.FirstAvailableIndex()
	for _, capt := range captures {
		if capt.Kind == piece.NoKind {
			log.Panicln("rules.Board.DoMove at", with, "returned empty square capture. Captures\n",
				captures, "\nchanges:\n", changes, "\nboard with changes applied:\n", an.Game.Board)
		}
		if capt.Orientation == rules.White {
			an.White.Captures[wCaptIndex] = capt.Kind
			wCaptIndex++
		} else if capt.Orientation == rules.Black {
			an.Black.Captures[bCaptIndex] = capt.Kind
			bCaptIndex++
		} else {
			log.Panicln("unknown orientation %v", capt.Kind)
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
		an.Active = an.opponentOf(an.Active)
	}

	capturedPieces := make([]rules.Piece, len(captures))
	for i, p := range captures {
		capturedPieces[i] = p.Piece
	}

	return changes, capturedPieces, promotionNeeded
}

func (an Instance) moveLegal(m rules.Move) bool {
	from := an.Board[m.From.Index()]
	if (from.Kind == piece.NoKind) || (from.Orientation != an.Active) {
		return false
	}

	// moves and state are read from the cache in an.Moves if cached
	moves, state := an.Moves()
	if (state != rules.Normal) && (state != rules.Check) {
		return false
	}

	if rules.MoveSetSliceHasMove(moves, m) == false {
		return false
	}

	return true
}
