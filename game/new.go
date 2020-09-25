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
	wp, wpicks, bp, bpicks, err := reserveArmies(wa, ba, white, black)
	if err != nil {
		return memory.NoGame
	}

	var w, b [16]rules.Piece
	for i, k := range wp {
		var start rules.AddressIndex
		if i < 8 {
			start = rules.AddressIndex(8 + i)
		} else {
			start = rules.AddressIndex(i - 8)
		}
		w[i] = rules.NewPiece(k, rules.White, false, start.Address())
	}
	for i, k := range bp {
		start := rules.AddressIndex(i + 48)
		b[i] = rules.NewPiece(k, rules.Black, false, start.Address())
	}

	g := memory.Game{
		Active:         rules.White,
		PreviousActive: rules.Black,
		White: memory.GamePlayer{
			PlayerIdentifier: white,
			Left:             wpicks.Left,
			Right:            wpicks.Right,
			Reward:           piece.RandomSpecialKind(),
		},
		Black: memory.GamePlayer{
			PlayerIdentifier: black,
			Left:             bpicks.Left,
			Right:            bpicks.Right,
			Reward:           piece.RandomSpecialKind(),
		},
		PreviousMove: rules.NoMove,
	}

	for i := 0; i < 8; i++ {
		g.Board[i] = w[i+8]
	}
	for i := 8; i < 16; i++ {
		g.Board[i] = w[i-8]
	}
	for i := 16; i < 48; i++ {
		g.Board[i] = rules.NoPiece
	}
	// the black army is a mirror of the white army, not a rotation
	for i := 48; i < 55; i++ {
		g.Board[i] = b[i-48]
	}
	for i := 55; i < 64; i++ {
		g.Board[i] = b[i-47]
	}

	return memory.AddGame(&g)
}
