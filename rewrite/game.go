package main

import (
	"database/sql"
	"strconv"
	"strings"
	"time"

	"github.com/pciet/wichess/rules"
)

func LoadGame(tx *sql.Tx, id GameIdentifier) Game {
	return Game{
		Header: LoadGameHeader(tx, id),
		Board:  LoadGameBoard(tx, id),
	}
}

// LoadGameBoard loads from the database and prepares a Board.
// If the game doesn't exist then the PieceIdentifiers member is nil.
func LoadGameBoard(tx *sql.Tx, id GameIdentifier) Board {
	var ep [64]EncodedPiece
	epp := make([]interface{}, 64)
	for i, _ := range ep {
		epp[i] = &(ep[i])
	}

	err := tx.QueryRow(GamesBoardQuery, id).Scan(epp...)
	if err == sql.ErrNoRows {
		return Board{PieceIdentifiers: nil}
	} else if err != nil {
		Panic(err)
	}

	b := Board{PieceIdentifiers: make([]AddressedPieceIdentifier, 0, 8)}

	for i, v := range ep {
		p := v.Decode()
		b.Board[i] = rules.Square(p.Piece.ApplyCharacteristics())
		if p.ID != 0 {
			b.PieceIdentifiers = append(b.PieceIdentifiers, AddressedPieceIdentifier{
				ID:      p.ID,
				Address: rules.AddressIndex(i).Address(),
			})
		}
	}

	return b
}

// LoadGameHeader gets the header from the database. If the header isn't found
// then the ID member is 0.
func LoadGameHeader(tx *sql.Tx, id GameIdentifier) GameHeader {
	h := GameHeader{ID: id}
	err := tx.QueryRow(GamesHeaderQuery, id).Scan(
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
		Panic("failed to query database row:", err)
	}
	return h
}

// NewGame creates a new game in the database, including loading the requested
// pieces from the database. If a piece request isn't valid then 0 is returned.
func NewGame(tx *sql.Tx, losesPieces bool,
	white string, whiteArmy ArmyRequest,
	black string, blackArmy ArmyRequest) GameIdentifier {
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
	err := tx.QueryRow(GamesNewInsert,
		rules.RandomSpecialPieceKind(),
		losesPieces,
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
		1,
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
		Panic("failed to insert new game:", err)
	}

	return id
}

// UpdateGame puts board changes into the database for a game, updates the latest move,
// sets the draw turn count, swaps the active player, and increments the turn number.
func UpdateGame(tx *sql.Tx, id GameIdentifier, white string, black string, active string,
	drawTurns int, turn int, m rules.Move, with []AddressedSquare) {
	var s strings.Buffer
	s.WriteString("UPDATE ")
	s.WriteString(GamesTable)
	s.WriteString(" SET ")

	i := 1

	placeholder := func(last bool) {
		s.WriteString("=$")
		s.WriteString(strconv.Itoa(i))
		if last == false {
			s.WriteString(", ")
		}
		i++
	}

	args := make([]interface{}, 0, 4)

	for _, s := range with {
		args = append(args, s.Encode())
		s.WriteRune('s')
		s.WriteString(strconv.Itoa(s.Address.Index()))
		placeholder(false)
	}

	switch active {
	case white:
		args = append(args, black)
	case black:
		args = append(args, white)
	default:
		Panic("game", h.ID, "active player", active, "not the white or black player")
	}
	s.WriteString(GamesActive)
	placeholder(false)

	// the move is recorded for future en passant calculation
	args = append(args, m.From.Index())
	s.WriteString(GamesFrom)
	placeholder(false)

	args = append(args, m.To.Index())
	s.WriteString(GamesTo)
	placeholder(false)

	// draw turns are reset or incremented depending on the move made
	args = append(args, drawTurns)
	s.WriteString(GamesDrawTurns)
	placeholder(false)

	args = append(args, turn+1)
	s.WriteString(GamesTurn)
	placeholder(true)

	args = append(args, strconv.Itoa(id))
	s.WriteString(" WHERE ")
	s.WriteString(GamesIdentifier)
	s.WriteRune('=')
	s.WriteString(strconv.Itoa(i))
	s.WriteRune(';')

	r, err := tx.Exec(s.String(), args...)
	if err != nil {
		Panic(err)
	}

	c, err := r.RowsAffected()
	if err != nil {
		Panic(err)
	}
	if c != 1 {
		Panic(c, "rows affected by", s.String())
	}
}

// GamesActiveAndOpponentName queries the database to show if this player is the active
// player and the opponent name. An empty string ("") is returned if the game doesn't exist.
// If the game is conceded then the player is always indicated as active.
func GameActiveAndOpponentName(tx *sql.Tx, id GameIdentifier, player string) (bool, string) {
	var conceded bool
	var active, white, black string
	err := tx.QueryRow(GamesOpponentAndActiveQuery, id).Scan(
		&active,
		&white,
		&black,
		&conceded,
	)
	if err == sql.ErrNoRows {
		DebugPrintln("no rows found for id", id, "and player", player)
		return false, ""
	} else if err != nil {
		Panic(err)
	}

	var opponent string
	if player == white {
		opponent = black
	} else if player == black {
		opponent = white
	} else {
		Panic("player", player, "doesn't match white", white, "or black", black)
	}

	if (active == player) || conceded {
		return true, opponent
	}

	return false, opponent
}

func GameHasPlayer(tx *sql.Tx, id GameIdentifier, name string) bool {
	var s sql.NullString
	err := tx.QueryRow(GameHasPlayerQuery, id, name).Scan(&s)
	if err == sql.ErrNoRows {
		return false
	} else if err != nil {
		Panic(err)
	}
	return true
}

func LoadGameTurn(tx *sql.Tx, id GameIdentifier) int {
	var t int
	err := tx.QueryRow(GamesTurnQuery, id).Scan(&t)
	if err != nil {
		Panic(err)
	}
	return t
}

func GameTurnEqual(tx *sql.Tx, id GameIdentifier, turn int) bool {
	t := LoadGameTurn(tx, id)
	if t == turn {
		return true
	}
	return false
}
