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

	games_white      = "white"
	games_black      = "black"
	games_active     = "active"
	games_identifier = "game_id"
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
func playersFor(gameID int) (string, string) {
	row := database.QueryRow(game_players_query, gameID)
	var white, black string
	err := row.Scan(&white, &black)
	if err != nil {
		panicExit(err.Error())
	}
	return white, black
}

type encodedGame struct {
	ID     int
	White  string
	Black  string
	Active string
	Points [64]pieceEncoding
}

type game struct {
	ID     int
	White  string
	Black  string
	Active string
	Points [64]piece
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

// Returns address that have changed, Kind 0 piece for a now empty point, but does not update the board points.
func (g game) move(from, to int, mover string) map[string]piece {
	var nextMover string
	var orientation wichessing.Orientation
	if g.White == g.Active {
		if g.White != mover {
			return nil
		}
		orientation = wichessing.White
		nextMover = g.Black
	} else {
		if g.Black != mover {
			return nil
		}
		orientation = wichessing.Black
		nextMover = g.White
	}
	diff := make(map[string]piece)
	for point, _ := range g.wichessingBoard().Move(absPoint(from), absPoint(to), orientation) {
		if point.Piece == nil {
			diff[point.AbsPoint.String()] = piece{
				Piece: wichessing.Piece{
					Kind: 0,
				},
			}
		} else {
			diff[point.AbsPoint.String()] = piece{
				Piece:      *point.Piece,
				Identifier: g.Points[from].Identifier,
			}
		}
	}
	writeGameChangesToDatabase(g.ID, diff, nextMover)
	if len(diff) != 0 {
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
	return diff
}

// The map keys are wichessing.AbsPoint converted to "x/file-y/rank" formatted string.
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
	for point, set := range board.Moves() {
		moves[point.String()] = set.String()
	}
	return moves
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

const game_query = "SELECT * FROM " + games_table + " WHERE " + games_identifier + "=$1"

func gameWithIdentifier(id int) game {
	row := database.QueryRow(game_query, id)
	g := encodedGame{}
	err := row.Scan(&g.ID, &g.White, &g.Black, &g.Active, &g.Points[0], &g.Points[1], &g.Points[2], &g.Points[3], &g.Points[4], &g.Points[5], &g.Points[6], &g.Points[7], &g.Points[8], &g.Points[9], &g.Points[10], &g.Points[11], &g.Points[12], &g.Points[13], &g.Points[14], &g.Points[15], &g.Points[16], &g.Points[17], &g.Points[18], &g.Points[19], &g.Points[20], &g.Points[21], &g.Points[22], &g.Points[23], &g.Points[24], &g.Points[25], &g.Points[26], &g.Points[27], &g.Points[28], &g.Points[29], &g.Points[30], &g.Points[31], &g.Points[32], &g.Points[33], &g.Points[34], &g.Points[35], &g.Points[36], &g.Points[37], &g.Points[38], &g.Points[39], &g.Points[40], &g.Points[41], &g.Points[42], &g.Points[43], &g.Points[44], &g.Points[45], &g.Points[46], &g.Points[47], &g.Points[48], &g.Points[49], &g.Points[50], &g.Points[51], &g.Points[52], &g.Points[53], &g.Points[54], &g.Points[55], &g.Points[56], &g.Points[57], &g.Points[58], &g.Points[59], &g.Points[60], &g.Points[61], &g.Points[62], &g.Points[63])
	if err != nil {
		panicExit(err.Error())
	}
	return game{
		ID:     g.ID,
		White:  g.White,
		Black:  g.Black,
		Active: g.Active,
		Points: decodedPoints(g.Points),
	}
}

const new_game_insert = "INSERT INTO " + games_table + " (" + games_white + ", " + games_black + ", " + games_active + ", s0, s1, s2, s3, s4, s5, s6, s7, s8, s9, s10, s11, s12, s13, s14, s15, s16, s17, s18, s19, s20, s21, s22, s23, s24, s25, s26, s27, s28, s29, s30, s31, s32, s33, s34, s35, s36, s37, s38, s39, s40, s41, s42, s43, s44, s45, s46, s47, s48, s49, s50, s51, s52, s53, s54, s55, s56, s57, s58, s59, s60, s61, s62, s63) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63, $64, $65, $66, $67) RETURNING " + games_identifier

// Returns the game identifier.
func newGameIntoDatabase(player1 string, player1setup gameSetup, player2 string, player2setup gameSetup) int {
	// https://github.com/lib/pq/issues/24
	var id int
	whitePawn := pieceWithID(0, wichessing.Pawn, wichessing.White, player1).encode()
	blackPawn := pieceWithID(0, wichessing.Pawn, wichessing.Black, player2).encode()
	err := database.QueryRow(new_game_insert, player1, player2, player1,
		pieceWithID(player1setup.leftRookID, wichessing.Rook, wichessing.White, player1).encode(),
		pieceWithID(player1setup.leftKnightID, wichessing.Knight, wichessing.White, player1).encode(),
		pieceWithID(player1setup.leftBishopID, wichessing.Bishop, wichessing.White, player1).encode(),
		pieceWithID(0, wichessing.King, wichessing.White, player1).encode(),
		pieceWithID(0, wichessing.Queen, wichessing.White, player1).encode(),
		pieceWithID(player1setup.rightBishopID, wichessing.Bishop, wichessing.White, player1).encode(),
		pieceWithID(player1setup.rightKnightID, wichessing.Knight, wichessing.White, player1).encode(),
		pieceWithID(player1setup.rightRookID, wichessing.Rook, wichessing.White, player1).encode(),
		whitePawn, whitePawn, whitePawn, whitePawn, whitePawn, whitePawn, whitePawn, whitePawn,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		blackPawn, blackPawn, blackPawn, blackPawn, blackPawn, blackPawn, blackPawn, blackPawn,
		pieceWithID(player2setup.leftRookID, wichessing.Rook, wichessing.Black, player2).encode(),
		pieceWithID(player2setup.leftKnightID, wichessing.Knight, wichessing.Black, player2).encode(),
		pieceWithID(player2setup.leftBishopID, wichessing.Bishop, wichessing.Black, player2).encode(),
		pieceWithID(0, wichessing.King, wichessing.Black, player2).encode(),
		pieceWithID(0, wichessing.Queen, wichessing.Black, player2).encode(),
		pieceWithID(player2setup.rightBishopID, wichessing.Bishop, wichessing.Black, player2).encode(),
		pieceWithID(player2setup.rightKnightID, wichessing.Knight, wichessing.Black, player2).encode(),
		pieceWithID(player2setup.rightRookID, wichessing.Rook, wichessing.Black, player2).encode()).Scan(&id)
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
