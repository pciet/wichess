package game

import (
	"log"

	"github.com/pciet/wichess/piece"
	"github.com/pciet/wichess/rules"
)

// Promote attempts to do the requested promotion for the active player. Nil is returned if the
// promotion couldn't be done.
func (an Instance) Promote(with piece.Kind) []rules.Square {
	_, pn := an.PromotionNeeded()
	if pn == false {
		return nil
	}

	change := an.Board.DoPromotion(with)
	if change.Piece.Kind == piece.NoKind {
		return nil
	}

	changes := []rules.Square{change}
	(&an.Game.Board).ApplyChanges(changes)

	an.MovesCache = nil
	an.StateCache = rules.NoState

	an.PreviousActive = an.Active
	an.Active = an.opponentOf(an.Active)

	return changes
}

func (an Instance) PromotingOrientation() rules.Orientation {
	o, needed := an.Game.Board.PromotionNeeded()
	if needed == false {
		log.Panicln("promotion not needed\n", an)
	}
	return o
}
