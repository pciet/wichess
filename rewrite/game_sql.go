package main

import (
	"fmt"
	"strings"
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
)

var (
	// these keys are used to select the game header values but also to generate the new game insert
	game_header_selects = []string{
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

	game_header_query = BuildSQLQuery(game_header_selects, games_table, games_identifier)

	game_opponent_and_active_query = BuildSQLQuery([]string{
		games_active,
		games_white,
		games_black,
		games_conceded,
	}, games_table, games_identifier)

	// TODO: is there a useful generalized BuildSQL func to make with this?
	games_new_insert = func() string {
		var s strings.Builder

		s.WriteString("INSERT INTO ")
		s.WriteString(games_table)
		s.WriteString(" (")
		for _, v := range game_header_selects {
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
		for i := 0; i < 64+len(game_header_selects); i++ {
			fmt.Fprintf(&s, "$%d", i)
			if i != 64+len(game_header_selects) {
				s.WriteString(", ")
			}
		}
		s.WriteString(") RETURNING ")
		s.WriteString(games_identifier)
		s.WriteRune(';')
		return s.String()
	}()
)
