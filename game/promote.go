package game

import (
	"github.com/pciet/wichess/piece"
	"github.com/pciet/wichess/rules"
)

// Promote attempts to do the requested promotion for the active player. Nil is returned if the
// promotion couldn't be done.
func (an Instance) Promote(with piece.Kind) []rules.AddressedSquare {
	if an.promotionNeeded() == false {
		return nil, nil, false
	}

	changes := &(an.Board).DoPromotion(with)

	an.PreviousActive = an.Active
	an.Active = an.OpponentOf(an.Active)

	return changes
}
