package game

import "github.com/pciet/wichess/rules"

// TODO: tune draw turns to what works for Wisconsin Chess

// DrawTurnMax is the number of turns without a pawn move or capture after which a draw is declared.
const DrawTurnMax = 50

// TODO: is there less work that can be done to just get State?

// Moves determines the moves available for the active player and what the current game state is.
func (an Instance) Moves() ([]rules.MoveSet, rules.State) {
	if an.Conceded {
		return nil, rules.Conceded
	} else if an.DrawTurns >= DrawTurnMax {
		return nil, rules.Draw
	}

	if an.moves != nil {
		return an.moves, an.state
	}

	an.moves, an.state = an.Board.Moves(an.Active, an.PreviousMove)
	if an.state == rules.Promotion {
		an.moves = []rules.MoveSet{}
		o, _ := an.Board.PromotionNeeded()
		if o != an.Active {
			an.state = rules.ReversePromotion
		}
	}
	return an.moves, an.state
}
