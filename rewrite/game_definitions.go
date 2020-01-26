package main

import (
	"strings"
	"time"

	"github.com/pciet/wichess/rules"
)

const (
	games_table = "games"

	games_piece                 = "piece"
	games_competitive           = "competitive"
	games_recorded              = "recorded"
	games_conceded              = "conceded"
	games_white                 = "white"
	games_white_acknowledge     = "white_ack"
	games_white_latest_move     = "white_latestmove"
	games_white_elapsed         = "white_elapsed"
	games_white_elapsed_updated = "white_elapsedupdated"
	games_black                 = "black"
	games_black_acknowledge     = "black_ack"
	games_black_latest_move     = "black_latestmove"
	games_black_elapsed         = "black_elapsed"
	games_black_elapsed_updated = "black_elapsedupdated"
	games_active                = "active"
	games_previous_active       = "previous_active"
	games_from                  = "move_from"
	games_to                    = "move_to"
	games_draw_turns            = "draw_turns"
	games_turn                  = "turn"
	games_identifier            = "game_id"

	no_move = 64 // value set for an initialized game with no move in from/to
)

type GameIdentifier int

type GameHeader struct {
	ID             GameIdentifier  `json:"-"`
	Competitive    bool            `json:"-"`
	PrizePiece     rules.PieceKind `json:"-"`
	Recorded       bool            `json:"-"`
	Conceded       bool            `json:"-"`
	White          PlayerGameHeader
	Black          PlayerGameHeader
	Active         rules.Orientation
	PreviousActive rules.Orientation
	From           int
	To             int
	DrawTurns      int `json:"-"`
	Turn           int
}

type PlayerGameHeader struct {
	Name          string
	Acknowledge   bool `json:"-"`
	LatestMove    time.Time
	Elapsed       time.Duration
	ElapsedUpdate time.Time
}

type Game struct {
	GameHeader
	rules.Game
	PieceIdentifiers []AddressedPieceID
}

const game_header_query_select = []string{
	games_piece,
	games_competitive,
	games_recorded,
	games_conceded,
	games_white,
	games_white_acknowledge,
	games_white_latest_move,
	games_white_elapsed,
	games_white_elapsed_updated,
	games_black,
	games_black_acknowledge,
	games_black_latest_move,
	games_black_elapsed,
	games_black_elapsed_updated,
	games_active,
	games_previous_active,
	games_from,
	games_to,
	games_draw_turns,
	games_turn,
}

var game_header_query = BuildSQLQuery(game_header_query_select, games_table, games_identifier)

const game_opponent_and_active_query_select = []string{
	games_active,
	games_white,
	games_black,
	games_conceded,
}

var game_opponent_and_active_query = BuildSQLQuery(game_opponent_and_active_query_select, games_table, games_identifier)

const game_insert_values = games_header_query_select

var games_new_insert = func() string {
	var s strings.Builder

	s.WriteString("INSERT INTO ")
	s.WriteString(games_table)
	s.WriteString(" (")
	for i, v := range game_insert_values {
		s.WriteString(v)
		s.WriteString(", ")
	}
	for i := 0; i < 64; i++ {
		fmt.Fprintf(&s, "s%d", i)
		if i != 63 {
			s.WriteString(", ")
		}
	}
	s.WriteString(") VALUES (")
	for i := 0; i < 64+len(game_insert_values); i++ {
		fmt.Fprintf(&s, "$%d", i)
		if i != 64+len(game_insert_values) {
			s.WriteString(", ")
		}
	}
	s.WriteString(") RETURNING ")
	s.WriteString(games_identifier)
	s.WriteRune(';')
	return s.String()
}()

var basic_army [16]rules.PieceKind

func init() {
	for i := 0; i < 8; i++ {
		basic_army[i] = rules.Pawn
	}

	basic_army[8] = rules.Rook
	basic_army[15] = rules.Rook

	basic_army[9] = rules.Knight
	basic_army[14] = rules.Knight

	basic_army[10] = rules.Bishop
	basic_army[13] = rules.Bishop

	basic_army[11] = rules.Queen
	basic_army[12] = rules.King
}
