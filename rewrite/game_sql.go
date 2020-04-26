package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	GamesTable = "games"

	GamesPiece            = "piece"
	GamesConceded         = "conceded"
	GamesWhite            = "white"
	GamesWhiteAcknowledge = "white_ack"
	GamesBlack            = "black"
	GamesBlackAcknowledge = "black_ack"
	GamesActive           = "active"
	GamesPreviousActive   = "previous_active"
	GamesFrom             = "move_from"
	GamesTo               = "move_to"
	GamesDrawTurns        = "draw_turns"
	GamesTurn             = "turn"
	GamesIdentifier       = "game_id"
)

var (
	// GameHeaderSelects are also used for the new game insert.
	GamesHeaderSelects = []string{
		GamesConceded,
		GamesWhite,
		GamesWhiteAcknowledge,
		GamesBlack,
		GamesBlackAcknowledge,
		GamesActive,
		GamesPreviousActive,
		GamesFrom,
		GamesTo,
		GamesDrawTurns,
		GamesTurn,
	}

	GamesHeaderQuery = SQLQuery(GamesHeaderSelects, GamesTable, GamesIdentifier)

	GamesBoardSelects = func() []string {
		s := make([]string, 64)
		for i, _ := range s {
			s[i] = "s" + strconv.Itoa(i)
		}
		return s
	}()

	GamesBoardQuery = SQLQuery(GamesBoardSelects, GamesTable, GamesIdentifier)

	GamesOpponentAndActiveQuery = SQLQuery([]string{
		GamesActive,
		GamesWhite,
		GamesBlack,
		GamesConceded,
	}, GamesTable, GamesIdentifier)

	// TODO: is there a useful generalized SQL func to make for GamesNewInsert?

	GamesNewInsert = func() string {
		var s strings.Builder

		s.WriteString("INSERT INTO ")
		s.WriteString(GamesTable)
		s.WriteString(" (")
		for _, v := range GamesHeaderSelects {
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
		for i := 1; i <= 64+len(GamesHeaderSelects); i++ {
			fmt.Fprintf(&s, "$%d", i)
			if i != 64+len(GamesHeaderSelects) {
				s.WriteString(", ")
			}
		}
		s.WriteString(") RETURNING ")
		s.WriteString(GamesIdentifier)
		s.WriteRune(';')
		if DebugSQL {
			fmt.Println(s.String())
		}
		return s.String()
	}()

	GamesHasPlayerQuery = SQLGeneralizedWhereQuery(nil, GamesTable,
		GamesIdentifier+"=$1 AND ("+GamesWhite+"=$2 OR "+GamesBlack+"=$2)")

	GamesTurnQuery           = SQLQuery([]string{GamesTurn}, GamesTable, GamesIdentifier)
	GamesOpponentQuery       = SQLQuery([]string{GamesWhite, GamesBlack}, GamesTable, GamesIdentifier)
	GamesPreviousActiveQuery = SQLQuery([]string{GamesPreviousActive}, GamesTable, GamesIdentifier)

	GamesAcknowledgeUpdate = func() string {
		var s strings.Builder
		s.WriteString("UPDATE ")
		s.WriteString(GamesTable)
		s.WriteString(" SET $1=$2 WHERE ")
		s.WriteString(GamesIdentifier)
		s.WriteString("=$3;")
		if DebugSQL {
			fmt.Println(s.String())
		}
		return s.String()
	}()

	GamesDelete = SQLDelete(GamesTable, GamesIdentifier)
)
