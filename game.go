// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"bytes"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/pciet/wichess/wichessing"
)

const (
	games_table = "games"

	games_white             = "white"
	games_white_acknowledge = "white_ack"
	games_black             = "black"
	games_black_acknowledge = "black_ack"
	games_active            = "active"
	games_identifier        = "game_id"
)

const game_players_query = "SELECT " + games_white + ", " + games_black + " FROM " + games_table + " WHERE " + games_identifier + "=$1"

func opponentFor(name string, gameID int) string {
	row := database.QueryRow(game_players_query, gameID)
	var white, black string
	err := row.Scan(&white, &black)
	if err != nil {
		panicExit(err.Error())
	}
	if name == white {
		return black
	} else if name == black {
		return white
	} else {
		panicExit(fmt.Sprintf("gameID %v has no player %v", gameID, name))
	}
	return ""
}

// White, Black.
func gamePlayers(gameID int) (string, string) {
	row := database.QueryRow(game_players_query, gameID)
	var white, black string
	err := row.Scan(&white, &black)
	if err != nil {
		panicExit(err.Error())
	}
	return white, black
}

type encodedGame struct {
	ID               int
	White            string
	WhiteAcknowledge bool
	Black            string
	BlackAcknowledge bool
	Active           string
	Points           [64]pieceEncoding
}

type game struct {
	ID               int
	White            string
	WhiteAcknowledge bool
	Black            string
	BlackAcknowledge bool
	Active           string
	Points           [64]piece
}

type gameListeners struct {
	white chan map[string]piece
	black chan map[string]piece
}

// TODO: mutex
var gameListening map[int]*gameListeners

func init() {
	gameListening = make(map[int]*gameListeners)
}

func decodedPoints(pts [64]pieceEncoding) [64]piece {
	var ret [64]piece
	for i := 0; i < 64; i++ {
		ret[i] = pts[i].decode()
		ret[i].Piece = ret[i].Piece.SetKindFlags()
	}
	return ret
}

func (g game) wichessingBoard() wichessing.Board {
	var board wichessing.Board
	for i := 0; i < 64; i++ {
		var p *wichessing.Piece
		if g.Points[i].Piece.Kind == 0 {
			p = nil
		} else {
			p = &g.Points[i].Piece
		}
		board[i] = wichessing.Point{
			Piece: p,
			AbsPoint: wichessing.AbsPoint{
				File: wichessing.FileFromIndex(uint8(i)),
				Rank: wichessing.RankFromIndex(uint8(i)),
			},
		}
	}
	return board
}

func absPoint(index int) wichessing.AbsPoint {
	return wichessing.AbsPoint{
		File: wichessing.FileFromIndex(uint8(index)),
		Rank: wichessing.RankFromIndex(uint8(index)),
	}
}

func hardComputerMoveForGame(id int) map[string]piece {
	panicExit("game: hard computer move not implemented")
	return nil
}

func easyComputerMoveForGame(id int) map[string]piece {
	g := gameWithIdentifier(id)
	var orientation wichessing.Orientation
	if easy_computer_player == g.White {
		orientation = wichessing.White
	} else {
		orientation = wichessing.Black
	}
	move := g.wichessingBoard().ComputerMove(orientation)
	if move == nil {
		return nil
	}
	diff, promoting := g.move(int(move.From.Index()), int(move.To.Index()), easy_computer_player)
	if promoting {
		after := g.wichessingBoard().AfterMove(move.From, move.To, orientation)
		var from int
		if orientation == wichessing.White {
			for i := 56; i < 64; i++ {
				p := after[i]
				if p.Piece == nil {
					continue
				}
				if (p.Kind == wichessing.Pawn) && (p.Orientation == wichessing.White) {
					from = i
					break
				}
			}
		} else {
			for i := 0; i < 8; i++ {
				p := after[i]
				if p.Piece == nil {
					continue
				}
				if (p.Kind == wichessing.Pawn) && (p.Orientation == wichessing.Black) {
					from = i
					break
				}
			}
		}
		pdiff := gameWithIdentifier(id).promote(from, easy_computer_player, wichessing.Queen)
		for point, piece := range pdiff {
			diff[point] = piece
		}
	}
	return diff
}

func (g game) promote(from int, player string, kind wichessing.Kind) map[string]piece {
	var nextMover string
	var orientation wichessing.Orientation
	if g.White == player {
		nextMover = g.Black
		orientation = wichessing.White
		if (from < 56) || (from > 63) {
			return nil
		}
	} else {
		nextMover = g.White
		orientation = wichessing.Black
		if (from < 0) || (from > 7) {
			return nil
		}
	}
	point := g.Points[from]
	if point.Kind == 0 {
		return nil
	}
	if (point.Orientation != orientation) || (point.Kind != wichessing.Pawn) {
		return nil
	}
	b := g.wichessingBoard()
	if b.Draw(orientation) {
		return nil
	}
	diff := make(map[string]piece)
	for point, _ := range b.PromotePawn(wichessing.AbsPointFromIndex(uint8(from)), wichessing.Kind(kind)) {
		if point.Piece == nil {
			diff[point.AbsPoint.String()] = piece{
				Piece: wichessing.Piece{
					Kind: 0,
				},
			}
		} else {
			diff[point.AbsPoint.String()] = piece{
				Piece: *point.Piece,
			}
		}
	}
	if (diff == nil) || (len(diff) == 0) {
		return diff
	}
	writeGameChangesToDatabase(g.ID, diff, nextMover)
	if (player != easy_computer_player) && (player != hard_computer_player) {
		if orientation == wichessing.White {
			if gameListening[g.ID].black != nil {
				gameListening[g.ID].black <- diff
			}
		} else {
			if gameListening[g.ID].white != nil {
				gameListening[g.ID].white <- diff
			}
		}
	}
	var checkOrientation wichessing.Orientation
	if orientation == wichessing.White {
		checkOrientation = wichessing.Black
	} else {
		checkOrientation = wichessing.White
	}
	after := b.AfterPromote(absPoint(from), wichessing.Kind(kind))
	if after.Checkmate(checkOrientation) {
		if checkOrientation == wichessing.White {
			writePlayerRecordUpdateToDatabase(g.Black, g.White, false)
		} else {
			writePlayerRecordUpdateToDatabase(g.White, g.Black, false)
		}
	} else if after.Draw(checkOrientation) {
		writePlayerRecordUpdateToDatabase(g.White, g.Black, true)
	}
	return diff
}

// Returns address that have changed, Kind 0 piece for a now empty point, but does not update the board points. Returns true if promoting.
func (g game) move(from, to int, mover string) (map[string]piece, bool) {
	var nextMover string
	var orientation, nextOrientation wichessing.Orientation
	if g.White == g.Active {
		if g.White != mover {
			return nil, false
		}
		orientation = wichessing.White
		nextOrientation = wichessing.Black
		nextMover = g.Black
	} else {
		if g.Black != mover {
			return nil, false
		}
		orientation = wichessing.Black
		nextOrientation = wichessing.White
		nextMover = g.White
	}
	b := g.wichessingBoard()
	if b.HasPawnToPromote() {
		return nil, false
	}
	if b.Draw(orientation) {
		return nil, false
	}
	diff := make(map[string]piece)
	for point, _ := range b.Move(absPoint(from), absPoint(to), orientation) {
		if point.Piece == nil {
			diff[point.AbsPoint.String()] = piece{
				Piece: wichessing.Piece{
					Kind: 0,
				},
			}
		} else {
			diff[point.AbsPoint.String()] = piece{
				Piece:      *point.Piece,
				Identifier: g.Points[from].Identifier, // TODO: this could be incorrect
			}
		}
	}
	if len(diff) == 0 {
		return diff, false
	}
	after := b.AfterMove(absPoint(from), absPoint(to), orientation)
	var promoting bool
	if after.HasPawnToPromote() {
		writeGameChangesToDatabase(g.ID, diff, mover)
		promoting = true
	} else {
		writeGameChangesToDatabase(g.ID, diff, nextMover)
	}
	if (mover != easy_computer_player) && (mover != hard_computer_player) {
		if orientation == wichessing.White {
			if gameListening[g.ID].black != nil {
				gameListening[g.ID].black <- diff
			}
		} else {
			if gameListening[g.ID].white != nil {
				gameListening[g.ID].white <- diff
			}
		}
	}
	if after.Checkmate(nextOrientation) {
		if nextOrientation == wichessing.White {
			writePlayerRecordUpdateToDatabase(g.Black, g.White, false)
		} else {
			writePlayerRecordUpdateToDatabase(g.White, g.Black, false)
		}
	} else if after.Draw(nextOrientation) {
		writePlayerRecordUpdateToDatabase(g.White, g.Black, true)
	}
	return diff, promoting
}

const (
	check_key     = "check"
	checkmate_key = "checkmate"
	promote_key   = "promote"
	draw_key      = "draw"
)

// The map keys are wichessing.AbsPoint converted to "x/file-y/rank" formatted string.
// If the game is in a check or checkmate state, or a piece is to be promoted, then a corresponding key with a nil value will be set.
func (g game) moves() map[string]map[string]struct{} {
	var board wichessing.Board
	for i := 0; i < 64; i++ {
		var p *wichessing.Piece
		if g.Points[i].Piece.Kind == 0 {
			p = nil
		} else {
			p = &g.Points[i].Piece
		}
		board[i] = wichessing.Point{
			Piece: p,
			AbsPoint: wichessing.AbsPoint{
				File: wichessing.FileFromIndex(uint8(i)),
				Rank: wichessing.RankFromIndex(uint8(i)),
			},
		}
	}
	moves := make(map[string]map[string]struct{})
	if board.HasPawnToPromote() {
		moves[promote_key] = nil
		return moves
	}
	var active wichessing.Orientation
	if g.Active == g.White {
		active = wichessing.White
	} else {
		active = wichessing.Black
	}
	if board.Draw(active) {
		moves[draw_key] = nil
		return moves
	}
	m, check, checkmate := board.Moves(active)
	for point, set := range m {
		moves[point.String()] = set.String()
	}
	if checkmate {
		moves[checkmate_key] = nil
	} else if check {
		moves[check_key] = nil
	}
	return moves
}

func (g *game) acknowledge(player string) bool {
	var active wichessing.Orientation
	if g.Active == g.Black {
		active = wichessing.Black
	} else {
		active = wichessing.White
	}
	b := g.wichessingBoard()
	if (b.Checkmate(active) == false) && (b.Draw(active) == false) {
		return false
	}
	var ackKey string
	if player == g.Black {
		ackKey = games_black_acknowledge
		g.BlackAcknowledge = true
	} else if player == g.White {
		ackKey = games_white_acknowledge
		g.WhiteAcknowledge = true
	} else {
		panicExit("player " + player + " is not " + g.Black + " (black) or " + g.White + " (white)")
	}
	if g.BlackAcknowledge && g.WhiteAcknowledge {
		g.deleteFromDatabase()
		return true
	}
	result, err := database.Exec("UPDATE "+games_table+" SET "+ackKey+" = $1 WHERE "+games_identifier+" = $2;", true, g.ID)
	if err != nil {
		panicExit(err.Error())
	}
	count, err := result.RowsAffected()
	if err != nil {
		panicExit(err.Error())
	}
	if count != 1 {
		panicExit(fmt.Sprintf("%v rows affected by ack update exec", count))
	}
	return true
}

func (g game) deleteFromDatabase() {
	result, err := database.Exec("DELETE FROM "+games_table+" WHERE "+games_identifier+" = $1;", g.ID)
	if err != nil {
		panicExit(err.Error())
	}
	count, err := result.RowsAffected()
	if err != nil {
		panicExit(err.Error())
	}
	if count != 1 {
		panicExit(fmt.Sprintf("%v rows affected by delete exec", count))
	}
}

func listeningToGame(name string, white string, black string, id int, socket *websocket.Conn) {
	// TODO: remove game from this map when finished
	_, has := gameListening[id]
	if has == false {
		gameListening[id] = &gameListeners{}
	}
	var l chan map[string]piece
	if name == white {
		gameListening[id].white = make(chan map[string]piece)
		l = gameListening[id].white
	} else if name == black {
		gameListening[id].black = make(chan map[string]piece)
		l = gameListening[id].black
	} else {
		panicExit("unexpected name " + name)
	}
	go func(listenTo chan map[string]piece, conn *websocket.Conn) {
		for {
			err := conn.WriteJSON(<-listenTo)
			if err != nil {
				if name == white {
					gameListening[id].white = nil
				} else {
					gameListening[id].black = nil
				}
				return
			}
		}
	}(l, socket)
}

const game_easy_computer_query = "SELECT " + games_identifier + " FROM " + games_table + " WHERE (" + games_white + " = $1 AND " + games_black + " = $2) OR (" + games_white + " = $3 AND " + games_black + " = $4);"

func easyComputerGame(player string) int {
	row := database.QueryRow(game_easy_computer_query, player, easy_computer_player, easy_computer_player, player)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0
	}
	return id
}

const game_query = "SELECT * FROM " + games_table + " WHERE " + games_identifier + "=$1;"

func gameWithIdentifier(id int) game {
	row := database.QueryRow(game_query, id)
	g := encodedGame{}
	err := row.Scan(&g.ID, &g.White, &g.WhiteAcknowledge, &g.Black, &g.BlackAcknowledge, &g.Active, &g.Points[0], &g.Points[1], &g.Points[2], &g.Points[3], &g.Points[4], &g.Points[5], &g.Points[6], &g.Points[7], &g.Points[8], &g.Points[9], &g.Points[10], &g.Points[11], &g.Points[12], &g.Points[13], &g.Points[14], &g.Points[15], &g.Points[16], &g.Points[17], &g.Points[18], &g.Points[19], &g.Points[20], &g.Points[21], &g.Points[22], &g.Points[23], &g.Points[24], &g.Points[25], &g.Points[26], &g.Points[27], &g.Points[28], &g.Points[29], &g.Points[30], &g.Points[31], &g.Points[32], &g.Points[33], &g.Points[34], &g.Points[35], &g.Points[36], &g.Points[37], &g.Points[38], &g.Points[39], &g.Points[40], &g.Points[41], &g.Points[42], &g.Points[43], &g.Points[44], &g.Points[45], &g.Points[46], &g.Points[47], &g.Points[48], &g.Points[49], &g.Points[50], &g.Points[51], &g.Points[52], &g.Points[53], &g.Points[54], &g.Points[55], &g.Points[56], &g.Points[57], &g.Points[58], &g.Points[59], &g.Points[60], &g.Points[61], &g.Points[62], &g.Points[63])
	if err != nil {
		return game{}
	}
	return game{
		ID:               g.ID,
		White:            g.White,
		WhiteAcknowledge: g.WhiteAcknowledge,
		Black:            g.Black,
		BlackAcknowledge: g.BlackAcknowledge,
		Active:           g.Active,
		Points:           decodedPoints(g.Points),
	}
}

const new_game_insert = "INSERT INTO " + games_table + " (" + games_white + ", " + games_white_acknowledge + ", " + games_black + ", " + games_black_acknowledge + ", " + games_active + ", s0, s1, s2, s3, s4, s5, s6, s7, s8, s9, s10, s11, s12, s13, s14, s15, s16, s17, s18, s19, s20, s21, s22, s23, s24, s25, s26, s27, s28, s29, s30, s31, s32, s33, s34, s35, s36, s37, s38, s39, s40, s41, s42, s43, s44, s45, s46, s47, s48, s49, s50, s51, s52, s53, s54, s55, s56, s57, s58, s59, s60, s61, s62, s63) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63, $64, $65, $66, $67, $68, $69) RETURNING " + games_identifier

// Returns the game identifier.
func newGameIntoDatabase(player1 string, player1setup gameSetup, player2 string, player2setup gameSetup) int {
	// https://github.com/lib/pq/issues/24
	var id int
	err := database.QueryRow(new_game_insert, player1, false, player2, false, player1,
		pieceWithID(player1setup[8], wichessing.Rook, wichessing.White, player1).encode(),
		pieceWithID(player1setup[9], wichessing.Knight, wichessing.White, player1).encode(),
		pieceWithID(player1setup[10], wichessing.Bishop, wichessing.White, player1).encode(),
		pieceWithID(player1setup[11], wichessing.Queen, wichessing.White, player1).encode(),
		pieceWithID(player1setup[12], wichessing.King, wichessing.White, player1).encode(),
		pieceWithID(player1setup[13], wichessing.Bishop, wichessing.White, player1).encode(),
		pieceWithID(player1setup[14], wichessing.Knight, wichessing.White, player1).encode(),
		pieceWithID(player1setup[15], wichessing.Rook, wichessing.White, player1).encode(),
		pieceWithID(player1setup[0], wichessing.Pawn, wichessing.White, player1).encode(),
		pieceWithID(player1setup[1], wichessing.Pawn, wichessing.White, player1).encode(),
		pieceWithID(player1setup[2], wichessing.Pawn, wichessing.White, player1).encode(),
		pieceWithID(player1setup[3], wichessing.Pawn, wichessing.White, player1).encode(),
		pieceWithID(player1setup[4], wichessing.Pawn, wichessing.White, player1).encode(),
		pieceWithID(player1setup[5], wichessing.Pawn, wichessing.White, player1).encode(),
		pieceWithID(player1setup[6], wichessing.Pawn, wichessing.White, player1).encode(),
		pieceWithID(player1setup[7], wichessing.Pawn, wichessing.White, player1).encode(),
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		pieceWithID(player2setup[0], wichessing.Pawn, wichessing.Black, player2).encode(),
		pieceWithID(player2setup[1], wichessing.Pawn, wichessing.Black, player2).encode(),
		pieceWithID(player2setup[2], wichessing.Pawn, wichessing.Black, player2).encode(),
		pieceWithID(player2setup[3], wichessing.Pawn, wichessing.Black, player2).encode(),
		pieceWithID(player2setup[4], wichessing.Pawn, wichessing.Black, player2).encode(),
		pieceWithID(player2setup[5], wichessing.Pawn, wichessing.Black, player2).encode(),
		pieceWithID(player2setup[6], wichessing.Pawn, wichessing.Black, player2).encode(),
		pieceWithID(player2setup[7], wichessing.Pawn, wichessing.Black, player2).encode(),
		pieceWithID(player2setup[8], wichessing.Rook, wichessing.Black, player2).encode(),
		pieceWithID(player2setup[9], wichessing.Knight, wichessing.Black, player2).encode(),
		pieceWithID(player2setup[10], wichessing.Bishop, wichessing.Black, player2).encode(),
		pieceWithID(player2setup[11], wichessing.Queen, wichessing.Black, player2).encode(),
		pieceWithID(player2setup[12], wichessing.King, wichessing.Black, player2).encode(),
		pieceWithID(player2setup[13], wichessing.Bishop, wichessing.Black, player2).encode(),
		pieceWithID(player2setup[14], wichessing.Knight, wichessing.Black, player2).encode(),
		pieceWithID(player2setup[15], wichessing.Rook, wichessing.Black, player2).encode()).Scan(&id)
	if err != nil {
		panicExit(err.Error())
	}
	return id
}

func writeGameChangesToDatabase(id int, diff map[string]piece, active string) {
	if (diff == nil) || (len(diff) == 0) {
		panicExit("no game changes recorded to database")
	}
	var query bytes.Buffer
	_, err := query.WriteString("UPDATE " + games_table + " SET ")
	if err != nil {
		panicExit(err.Error())
	}
	i := 1
	args := make([]interface{}, 0, 4)
	for addr, p := range diff {
		args = append(args, p.encode().String())
		_, err = query.WriteString(fmt.Sprintf("s%v = $%v, ", wichessing.IndexFromAddressString(addr), i))
		if err != nil {
			panicExit(err.Error())
		}
		i++
	}
	args = append(args, active)
	_, err = query.WriteString(fmt.Sprintf(games_active+" = $%v", i))
	if err != nil {
		panicExit(err.Error())
	}
	i++
	args = append(args, fmt.Sprintf("%v", id))
	_, err = query.WriteString(" WHERE " + games_identifier + " = " + fmt.Sprintf("$%v", i) + ";")
	if err != nil {
		panicExit(err.Error())
	}
	result, err := database.Exec(query.String(), args...)
	if err != nil {
		panicExit(err.Error())
	}
	count, err := result.RowsAffected()
	if err != nil {
		panicExit(err.Error())
	}
	if count != 1 {
		panicExit(fmt.Sprintf("%v rows affected by: %v", count, query.String()))
	}
}
