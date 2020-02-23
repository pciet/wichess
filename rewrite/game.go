package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/pciet/wichess/rules"
)

// Returns a GameHeader with ID set to 0 if the game isn't found.
func LoadGameHeader(tx *sql.Tx, id GameIdentifier) GameHeader {
	h := GameHeader{ID: id}
	err := tx.QueryRow(game_header_query, id).Scan(
		&h.PrizePiece,
		&h.Competitive,
		&h.Recorded,
		&h.Conceded,
		&h.White.Name,
		&h.White.Acknowledge,
		&h.White.LatestMove,
		&h.White.Elapsed,
		&h.White.ElapsedUpdated,
		&h.Black.Name,
		&h.Black.Acknowledge,
		&h.Black.LatestMove,
		&h.Black.Elapsed,
		&h.Black.ElapsedUpdated,
		&h.Active,
		&h.PreviousActive,
		&h.From,
		&h.To,
		&h.DrawTurns,
		&h.Turn)
	if err == sql.ErrNoRows {
		DebugPrintln("found no games with id", id)
		h.ID = 0
	} else if err != nil {
		log.Panicln("failed to query database row:", err)
	}
	return h
}

// Returns 0 if the game couldn't be created.
func NewGame(tx *sql.Tx, white string, whiteArmy ArmyRequest, black string, blackArmy ArmyRequest, competitive bool) GameIdentifier {
	var wp, bp [16]EncodedPiece

	enc := func(to *[16]EncodedPiece, with ArmyRequest, o rules.Orientation, name string) bool {
		for i := 0; i < 16; i++ {
			p := LoadPiece(tx, with[i], basic_army[i], o, name)
			if p.Kind == rules.NoKind {
				DebugPrintln("bad request to LoadPiece for player", name, "piece ID", with[i])
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

	now := time.Now()

	// QueryRow instead of Exec: https://github.com/lib/pq/issues/24
	var id GameIdentifier
	err := tx.QueryRow(games_new_insert,
		rules.RandomSpecialPieceKind(),
		competitive,
		false,
		false,
		white,
		false,
		now,
		time.Duration(0),
		now,
		black,
		false,
		now,
		time.Duration(0),
		now,
		white,
		black,
		no_move,
		no_move,
		0,
		0,
		wp[8], wp[9], wp[10], wp[11], wp[12], wp[13], wp[14], wp[15],
		wp[0], wp[1], wp[2], wp[3], wp[4], wp[5], wp[6], wp[7],
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		bp[0], bp[1], bp[2], bp[3], bp[4], bp[5], bp[6], bp[7],
		bp[8], bp[9], bp[10], bp[11], bp[12], bp[13], bp[14], bp[15],
	).Scan(&id)
	if err != nil {
		log.Panicln("failed to insert new game:", err)
	}

	return id
}

// Returns if this player is the active player and the name of the opponent.
// Returns a "" string if the game doesn't exist.
// If the game is conceded then the player is always marked as active.
func GameActiveAndOpponentName(tx *sql.Tx, id GameIdentifier, player string) (bool, string) {
	var conceded bool
	var active, white, black string
	err := tx.QueryRow(game_opponent_and_active_query, id).Scan(
		&active,
		&white,
		&black,
		&conceded,
	)
	if err == sql.ErrNoRows {
		DebugPrintln("no rows found for id", id, "and player", player)
		return false, ""
	} else if err != nil {
		log.Panic(err)
	}

	var opponent string
	if player == white {
		opponent = black
	} else if player == black {
		opponent = white
	} else {
		log.Panicln("player", player, "doesn't match white", white, "or black", black)
	}

	if (active == player) || conceded {
		return true, opponent
	}

	return false, opponent
}
