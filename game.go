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
	WhiteAcknowledge bool `json:"-"`
	Black            string
	BlackAcknowledge bool `json:"-"`
	Active           string
	Points           [64]piece
	DB               DB `json:"-"`
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

func (db DB) gameWithIdentifier(id int) game {
	row := db.QueryRow("SELECT * FROM "+games_table+" WHERE "+games_identifier+"=$1;", id)
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
		DB:               db,
	}
}

func (db DB) gameOpponent(name string, gameID int) string {
	white, black := db.gamePlayers(gameID)
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
func (db DB) gamePlayers(gameID int) (string, string) {
	row := db.QueryRow("SELECT "+games_white+", "+games_black+" FROM "+games_table+" WHERE "+games_identifier+"=$1;", gameID)
	var white, black string
	err := row.Scan(&white, &black)
	if err != nil {
		panicExit(err.Error())
	}
	return white, black
}

func (db DB) updateGame(id int, diff map[string]piece, active string) {
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
	result, err := db.Exec(query.String(), args...)
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

func (g *game) acknowledge(player string) bool {
	var active wichessing.Orientation
	if g.Active == g.Black {
		active = wichessing.Black
	} else {
		active = wichessing.White
	}
	b := wichessingBoard(g.Points)
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
	isCompetitive48, slot := g.DB.gameIsCompetitive48ForPlayer(g.ID, player)
	if isCompetitive48 {
		g.DB.removePlayersCompetitive48Game(player, slot)
	}
	if g.BlackAcknowledge && g.WhiteAcknowledge {
		g.DB.deleteGame(g.ID)
		return true
	}
	result, err := g.DB.Exec("UPDATE "+games_table+" SET "+ackKey+" = $1 WHERE "+games_identifier+" = $2;", true, g.ID)
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

// Returns the game identifier.
func (db DB) newGame(player1 string, player1setup gameSetup, player2 string, player2setup gameSetup) int {
	// https://github.com/lib/pq/issues/24
	var id int
	err := db.QueryRow("INSERT INTO "+games_table+" ("+games_white+", "+games_white_acknowledge+", "+games_black+", "+games_black_acknowledge+", "+games_active+", s0, s1, s2, s3, s4, s5, s6, s7, s8, s9, s10, s11, s12, s13, s14, s15, s16, s17, s18, s19, s20, s21, s22, s23, s24, s25, s26, s27, s28, s29, s30, s31, s32, s33, s34, s35, s36, s37, s38, s39, s40, s41, s42, s43, s44, s45, s46, s47, s48, s49, s50, s51, s52, s53, s54, s55, s56, s57, s58, s59, s60, s61, s62, s63) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63, $64, $65, $66, $67, $68, $69) RETURNING "+games_identifier+";",
		player1, false, player2, false, player1,
		db.pieceWithID(player1setup[8], wichessing.Rook, wichessing.White, player1).encode(),
		db.pieceWithID(player1setup[9], wichessing.Knight, wichessing.White, player1).encode(),
		db.pieceWithID(player1setup[10], wichessing.Bishop, wichessing.White, player1).encode(),
		db.pieceWithID(player1setup[11], wichessing.Queen, wichessing.White, player1).encode(),
		db.pieceWithID(player1setup[12], wichessing.King, wichessing.White, player1).encode(),
		db.pieceWithID(player1setup[13], wichessing.Bishop, wichessing.White, player1).encode(),
		db.pieceWithID(player1setup[14], wichessing.Knight, wichessing.White, player1).encode(),
		db.pieceWithID(player1setup[15], wichessing.Rook, wichessing.White, player1).encode(),
		db.pieceWithID(player1setup[0], wichessing.Pawn, wichessing.White, player1).encode(),
		db.pieceWithID(player1setup[1], wichessing.Pawn, wichessing.White, player1).encode(),
		db.pieceWithID(player1setup[2], wichessing.Pawn, wichessing.White, player1).encode(),
		db.pieceWithID(player1setup[3], wichessing.Pawn, wichessing.White, player1).encode(),
		db.pieceWithID(player1setup[4], wichessing.Pawn, wichessing.White, player1).encode(),
		db.pieceWithID(player1setup[5], wichessing.Pawn, wichessing.White, player1).encode(),
		db.pieceWithID(player1setup[6], wichessing.Pawn, wichessing.White, player1).encode(),
		db.pieceWithID(player1setup[7], wichessing.Pawn, wichessing.White, player1).encode(),
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		db.pieceWithID(player2setup[0], wichessing.Pawn, wichessing.Black, player2).encode(),
		db.pieceWithID(player2setup[1], wichessing.Pawn, wichessing.Black, player2).encode(),
		db.pieceWithID(player2setup[2], wichessing.Pawn, wichessing.Black, player2).encode(),
		db.pieceWithID(player2setup[3], wichessing.Pawn, wichessing.Black, player2).encode(),
		db.pieceWithID(player2setup[4], wichessing.Pawn, wichessing.Black, player2).encode(),
		db.pieceWithID(player2setup[5], wichessing.Pawn, wichessing.Black, player2).encode(),
		db.pieceWithID(player2setup[6], wichessing.Pawn, wichessing.Black, player2).encode(),
		db.pieceWithID(player2setup[7], wichessing.Pawn, wichessing.Black, player2).encode(),
		db.pieceWithID(player2setup[8], wichessing.Rook, wichessing.Black, player2).encode(),
		db.pieceWithID(player2setup[9], wichessing.Knight, wichessing.Black, player2).encode(),
		db.pieceWithID(player2setup[10], wichessing.Bishop, wichessing.Black, player2).encode(),
		db.pieceWithID(player2setup[11], wichessing.Queen, wichessing.Black, player2).encode(),
		db.pieceWithID(player2setup[12], wichessing.King, wichessing.Black, player2).encode(),
		db.pieceWithID(player2setup[13], wichessing.Bishop, wichessing.Black, player2).encode(),
		db.pieceWithID(player2setup[14], wichessing.Knight, wichessing.Black, player2).encode(),
		db.pieceWithID(player2setup[15], wichessing.Rook, wichessing.Black, player2).encode()).Scan(&id)
	if err != nil {
		panicExit(err.Error())
	}
	return id
}

func (db DB) deleteGame(id int) {
	result, err := db.Exec("DELETE FROM "+games_table+" WHERE "+games_identifier+" = $1;", id)
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

func decodedPoints(pts [64]pieceEncoding) [64]piece {
	var ret [64]piece
	for i := 0; i < 64; i++ {
		ret[i] = pts[i].decode()
		ret[i].Piece = ret[i].Piece.SetKindFlags()
	}
	return ret
}

func wichessingBoard(points [64]piece) wichessing.Board {
	var board wichessing.Board
	for i := 0; i < 64; i++ {
		var p *wichessing.Piece
		if points[i].Piece.Kind == 0 {
			p = nil
		} else {
			p = &points[i].Piece
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
