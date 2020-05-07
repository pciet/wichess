package main

import (
	"database/sql"

	"github.com/lib/pq"
	"github.com/pciet/wichess/rules"
)

// NewGame creates a new game in the database, including loading the requested
// pieces from the database. If a piece request isn't valid then 0 is returned.
func NewGame(tx *sql.Tx, losesPieces bool,
	white string, whiteArmy ArmyRequest,
	black string, blackArmy ArmyRequest) GameIdentifier {
	var wp, bp [16]EncodedPiece

	enc := func(to *[16]EncodedPiece, with ArmyRequest,
		o rules.Orientation, name string) bool {
		var c Collection
		if name != ComputerPlayerName {
			c = PlayerCollection(tx, PlayerID(tx, name))
		}
		for i := 0; i < 16; i++ {
			p := ConfigurePiece(c[i], BasicArmy[i], o)
			if p.Kind == rules.NoKind {
				DebugPrintln("bad request to LoadPiece for player",
					name, "piece ID", with[i])
				return false
			}
			(*to)[i] = p.Encode()
		}
		return true
	}

	ok := enc(&wp, whiteArmy, rules.White, white)
	if ok == false {
		return 0
	}
	ok = enc(&bp, blackArmy, rules.Black, black)
	if ok == false {
		return 0
	}

	// QueryRow instead of Exec: https://github.com/lib/pq/issues/24
	var id GameIdentifier
	err := tx.QueryRow(GamesNewInsert,
		false,
		white, false,
		black, false,
		rules.White, rules.Black,
		NoMove, NoMove,
		0, 1,
		pq.Array([]EncodedPiece{
			wp[8], wp[9], wp[10], wp[11], wp[12], wp[13], wp[14], wp[15],
			wp[0], wp[1], wp[2], wp[3], wp[4], wp[5], wp[6], wp[7],
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			bp[0], bp[1], bp[2], bp[3], bp[4], bp[5], bp[6], bp[7],
			bp[8], bp[9], bp[10], bp[11], bp[12], bp[13], bp[14], bp[15],
		}),
	).Scan(&id)
	if err != nil {
		Panic("failed to insert new game:", err)
	}

	return id
}
