// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/pciet/wichess/wichessing"
)

const (
	games_table = "games"

	games_competitive           = "competitive"
	games_recorded              = "recorded"
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
	games_identifier            = "game_id"
)

type GameInfo struct {
	ID                  int
	Competitive         bool
	Recorded            bool
	White               string
	WhiteAcknowledge    bool
	WhiteLatestMove     time.Time
	WhiteElapsed        time.Duration
	WhiteElapsedUpdated time.Time
	Black               string
	BlackAcknowledge    bool
	BlackLatestMove     time.Time
	BlackElapsed        time.Duration
	BlackElapsedUpdated time.Time
	Active              string
}

type game struct {
	GameInfo
	Points [64]piece
	DB     DB `json:"-"`
}

func (db DB) gameRecorded(gameID int) bool {
	var recorded bool
	err := db.QueryRow("SELECT "+games_recorded+" FROM "+games_table+" WHERE "+games_identifier+"=$1;", gameID).Scan(&recorded)
	if err != nil {
		panicExit(err.Error())
	}
	return recorded
}

func (db DB) setGameRecorded(id int) {
	result, err := db.Exec("UPDATE "+games_table+" SET "+games_recorded+" = $1 WHERE "+games_identifier+" = $2;", true, id)
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
}

func (g game) activeOrientation() wichessing.Orientation {
	if g.Active == g.White {
		return wichessing.White
	} else if g.Active == g.Black {
		return wichessing.Black
	} else {
		panicExit(fmt.Sprintf("%v is not white (%v) or black (%v)", g.Active, g.White, g.Black))
		return wichessing.White
	}
}

func (db DB) gameInfo(id int) GameInfo {
	g := GameInfo{ID: id}
	err := db.QueryRow("SELECT "+games_competitive+", "+games_recorded+", "+games_white+", "+games_white_acknowledge+", "+games_white_latest_move+", "+games_white_elapsed+", "+games_white_elapsed_updated+", "+games_black+", "+games_black_acknowledge+", "+games_black_latest_move+", "+games_black_elapsed+", "+games_black_elapsed_updated+", "+games_active+" FROM "+games_table+" WHERE "+games_identifier+"=$1;", id).Scan(&g.Competitive, &g.Recorded, &g.White, &g.WhiteAcknowledge, &g.WhiteLatestMove, &g.WhiteElapsed, &g.WhiteElapsedUpdated, &g.Black, &g.BlackAcknowledge, &g.BlackLatestMove, &g.BlackElapsed, &g.BlackElapsedUpdated, &g.Active)
	if err != nil {
		panicExit(err.Error())
	}
	return g
}

func (db DB) gameWithIdentifier(id int) game {
	row := db.QueryRow("SELECT * FROM "+games_table+" WHERE "+games_identifier+"=$1;", id)
	g := GameInfo{}
	var Points [64]pieceEncoding
	err := row.Scan(&g.ID, &g.Competitive, &g.Recorded, &g.White, &g.WhiteAcknowledge, &g.WhiteLatestMove, &g.WhiteElapsed, &g.WhiteElapsedUpdated, &g.Black, &g.BlackAcknowledge, &g.BlackLatestMove, &g.BlackElapsed, &g.BlackElapsedUpdated, &g.Active, &Points[0], &Points[1], &Points[2], &Points[3], &Points[4], &Points[5], &Points[6], &Points[7], &Points[8], &Points[9], &Points[10], &Points[11], &Points[12], &Points[13], &Points[14], &Points[15], &Points[16], &Points[17], &Points[18], &Points[19], &Points[20], &Points[21], &Points[22], &Points[23], &Points[24], &Points[25], &Points[26], &Points[27], &Points[28], &Points[29], &Points[30], &Points[31], &Points[32], &Points[33], &Points[34], &Points[35], &Points[36], &Points[37], &Points[38], &Points[39], &Points[40], &Points[41], &Points[42], &Points[43], &Points[44], &Points[45], &Points[46], &Points[47], &Points[48], &Points[49], &Points[50], &Points[51], &Points[52], &Points[53], &Points[54], &Points[55], &Points[56], &Points[57], &Points[58], &Points[59], &Points[60], &Points[61], &Points[62], &Points[63])
	if err != nil {
		panicExit(err.Error())
	}
	return game{
		GameInfo: g,
		Points:   decodedPoints(Points),
		DB:       db,
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

func (g *game) acknowledgeGameComplete(player string) bool {
	active := (*g).activeOrientation()
	var totalTime time.Duration
	var c5, c15 int
	if (player != hard_computer_player) && (player != easy_computer_player) {
		c5 = g.DB.playersCompetitive5HourGameID(player)
		if c5 != 0 {
			totalTime = competitive5_total_time
		} else {
			c15 = g.DB.playersCompetitive15HourGameID(player)
			if c15 != 0 {
				totalTime = competitive15_total_time
			}
		}
	}
	timeLoss := false
	if totalTime > time.Duration(0) {
		if active == wichessing.White {
			if g.WhiteElapsed > totalTime {
				timeLoss = true
			}
		} else {
			if g.BlackElapsed > totalTime {
				timeLoss = true
			}
		}
	}
	b := wichessingBoard(g.Points)
	if (b.Checkmate(active) == false) && (b.Draw(active) == false) && (timeLoss == false) {
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
	if (player != easy_computer_player) && (player != hard_computer_player) {
		if c5 != 0 {
			g.DB.removePlayersCompetitive5Game(player)
		} else if c15 != 0 {
			g.DB.removePlayersCompetitive15Game(player)
		} else {
			isCompetitive48, slot := g.DB.gameIsCompetitive48ForPlayer(g.ID, player)
			if isCompetitive48 {
				g.DB.removePlayersCompetitive48Game(player, slot)
			}
		}
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
func (db DB) newGame(player1 string, player1setup gameSetup, player2 string, player2setup gameSetup, competitive bool) int {
	// https://github.com/lib/pq/issues/24
	var id int
	err := db.QueryRow("INSERT INTO "+games_table+" ("+games_competitive+", "+games_recorded+", "+games_white+", "+games_white_acknowledge+", "+games_white_latest_move+", "+games_white_elapsed+", "+games_white_elapsed_updated+", "+games_black+", "+games_black_acknowledge+", "+games_black_latest_move+", "+games_black_elapsed+", "+games_black_elapsed_updated+", "+games_active+", s0, s1, s2, s3, s4, s5, s6, s7, s8, s9, s10, s11, s12, s13, s14, s15, s16, s17, s18, s19, s20, s21, s22, s23, s24, s25, s26, s27, s28, s29, s30, s31, s32, s33, s34, s35, s36, s37, s38, s39, s40, s41, s42, s43, s44, s45, s46, s47, s48, s49, s50, s51, s52, s53, s54, s55, s56, s57, s58, s59, s60, s61, s62, s63) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63, $64, $65, $66, $67, $68, $69, $70, $71, $72, $73, $74, $75, $76, $77) RETURNING "+games_identifier+";",
		competitive, false, player1, false, time.Now(), time.Duration(0), time.Now(), player2, false, time.Now(), time.Duration(0), time.Now(), player1,
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

func decodedPoints(pts [64]pieceEncoding) [64]piece {
	var ret [64]piece
	for i := 0; i < 64; i++ {
		if pts[i] == 0 {
			continue
		}
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
