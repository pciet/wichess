package game

import (
	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/piece"
	"github.com/pciet/wichess/rules"
)

// New verifies the army requests are legal then adds a new game to memory and returns its
// identifier. If an army request isn't legal then memory.NoGame is returned and no memory is added.
//
//See piece.ReserveArmies for the criteria of a legal army request.
func New(wa, ba piece.ArmyRequest, white, black memory.PlayerIdentifier) memory.GameIdentifier {
	wp, wpicks, bp, bpicks, err := piece.ReserveArmies(wa, ba, white, black)
	if err != nil {
		return memory.NoGame
	}

	g := memory.Game{
		Active:         rules.White,
		PreviousActive: rules.Black,
		White: GamePlayer{
			PlayerIdentifier: white,
			Left:             wpicks.Left,
			Right:            wpicks.Right,
			Reward:           piece.RandomSpecialKind(),
		},
		Black: GamePlayer{
			PlayerIdentifier: black,
			Left:             bpicks.Left,
			Right:            bpicks.Right,
			Reward:           piece.RandomSpecialKind(),
		},
		PreviousMove: rules.NoMove,
		// the black army is a mirror of the white army, not a rotation
		Board: rules.Board{
			wp[8], wp[9], wp[10], wp[11], wp[12], wp[13], wp[14], wp[15],
			wp[0], wp[1], wp[2], wp[3], wp[4], wp[5], wp[6], wp[7],
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			bp[0], bp[1], bp[2], bp[3], bp[4], bp[5], bp[6], bp[7],
			bp[8], bp[9], bp[10], bp[11], bp[12], bp[13], bp[14], bp[15],
		},
	}

	return memory.AddGame(&g)
}
