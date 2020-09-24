package game

import (
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

	changes := an.Board.DoPromotion(with)

	an.PreviousActive = an.Active
	an.Active = an.opponentOf(an.Active)

	return []rules.Square{changes}
}
