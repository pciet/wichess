package main

import (
	"fmt"
	"strings"
)

const (
	GamesTable = "games"

	GamesIdentifier = "id"

	GamesConceded = "conceded"

	GamesWhite            = "white"
	GamesWhiteAcknowledge = "white_ack"

	GamesBlack            = "black"
	GamesBlackAcknowledge = "black_ack"

	GamesActive         = "active"
	GamesPreviousActive = "previous_active"

	GamesMoveFrom = "move_from"
	GamesMoveTo   = "move_to"

	GamesDrawTurns = "draw_turns"
	GamesTurn      = "turn"

	GamesBoard       = "board"
	GamesBoardLength = 64
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
		GamesMoveFrom,
		GamesMoveTo,
		GamesDrawTurns,
		GamesTurn,
	}

	GamesHeaderQuery          = SQLQuery(GamesHeaderSelects, GamesTable, GamesIdentifier)
	GamesHeaderForUpdateQuery = SQLForUpdateQuery(GamesHeaderSelects,
		GamesTable, GamesIdentifier)

	GamesBoardQuery          = SQLQuery([]string{GamesBoard}, GamesTable, GamesIdentifier)
	GamesBoardForUpdateQuery = SQLForUpdateQuery([]string{GamesBoard},
		GamesTable, GamesIdentifier)

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
		s.WriteString(GamesBoard)
		s.WriteString(") VALUES (")
		for i := 1; i <= len(GamesHeaderSelects)+1; i++ {
			fmt.Fprintf(&s, "$%d", i)
			if i != (len(GamesHeaderSelects) + 1) {
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

	GamesTurnQuery     = SQLQuery([]string{GamesTurn}, GamesTable, GamesIdentifier)
	GamesOpponentQuery = SQLQuery([]string{GamesWhite, GamesBlack},
		GamesTable, GamesIdentifier)
	GamesPreviousActiveQuery = SQLQuery([]string{GamesPreviousActive},
		GamesTable, GamesIdentifier)
	GamesPlayersQuery = SQLQuery([]string{GamesWhite, GamesBlack},
		GamesTable, GamesIdentifier)

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
