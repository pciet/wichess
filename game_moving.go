// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"fmt"
	"time"

	"github.com/pciet/wichess/wichessing"
)

const (
	check_key     = "check"
	checkmate_key = "checkmate"
	promote_key   = "promote"
	draw_key      = "draw"
	time_key      = "time"
	conceded_key  = "conceded"

	draw_turn_count = 50
)

// TODO: reduce duplication in game.move and game.promote

// Returns address that have changed, Kind 0 piece for a now empty point, but does not update the board points. Returns true and an orientation of the promoter if promoting.
func (g game) move(from, to int, mover string, tx TX) (map[string]piece, bool, wichessing.Orientation) {
	if g.Active != mover {
		if debug {
			fmt.Println("move: active player is not mover")
		}
		return nil, false, wichessing.White
	}
	var nextMover string
	var orientation, nextOrientation wichessing.Orientation
	if g.White == g.Active {
		orientation = wichessing.White
		nextOrientation = wichessing.Black
		nextMover = g.Black
	} else {
		orientation = wichessing.Black
		nextOrientation = wichessing.White
		nextMover = g.White
	}
	b := wichessingBoard(g.Points, g.From, g.To)
	// TODO: check for cases of double-checking
	pring, _ := b.HasPawnToPromote()
	if pring {
		if debug {
			fmt.Println("move: has pawn to promote")
		}
		return nil, false, wichessing.White
	}
	if (g.DrawTurns >= draw_turn_count) || b.Draw(orientation) || g.Conceded {
		if debug {
			fmt.Println("move: draw or conceded determined")
		}
		return nil, false, wichessing.White
	}
	diff := make(map[string]piece)
	difference, taken := b.Move(absPoint(from), absPoint(to), orientation)
	for _, point := range difference {
		if point.Piece == nil {
			diff[point.AbsPoint.String()] = piece{
				Piece: wichessing.Piece{
					Kind: 0,
				},
			}
		} else {
			// the identifier must be set here for the game update to work correctly
			// make taken = map[AbsPoint]Piece instead of PieceSet
			var id int
			for index, ppt := range b.Points {
				if ppt.Piece == point.Piece {
					id = g.Points[index].Identifier
					break
				}
			}
			diff[point.AbsPoint.String()] = piece{
				Piece:      *point.Piece,
				Identifier: id,
			}
		}
	}
	if len(diff) == 0 {
		if debug {
			fmt.Println("move: b.Move returned zero length diff")
		}
		return diff, false, wichessing.White
	}
	takenPieces := make(map[int]struct{})
	for point, _ := range taken {
		takenPieces[g.Points[point.Index()].Identifier] = struct{}{}
	}
	after := b.AfterMove(absPoint(from), absPoint(to), orientation)
	promoting, promotingOrientation := after.HasPawnToPromote()
	if promoting && (promotingOrientation == orientation) {
		tx.updateGame(g.ID, diff, mover, g.Active, from, to, 0, g.Turn)
	} else {
		if (len(taken) == 0) && (b.Points[from].Base != wichessing.Pawn) {
			tx.updateGame(g.ID, diff, nextMover, g.Active, from, to, g.DrawTurns+1, g.Turn)
		} else {
			tx.updateGame(g.ID, diff, nextMover, g.Active, from, to, 0, g.Turn)
		}
	}
	go func() {
		gameMonitorsLock.RLock()
		c, has := gameMonitors[g.ID]
		if has {
			c.move <- time.Now()
		}
		gameMonitorsLock.RUnlock()
	}()
	if (mover != easy_computer_player) && (mover != hard_computer_player) {
		gameListeningLock.RLock()
		listeners, has := gameListening[g.ID]
		if (orientation == wichessing.White) && has {
			if listeners.black != nil {
				listeners.black <- diff
			}
		}
		if (orientation == wichessing.Black) && has {
			if listeners.white != nil {
				listeners.white <- diff
			}
		}
		gameListeningLock.RUnlock()
	}
	if g.Competitive {
		// competitively taken collectible pieces are no longer available to the owning player
		for id, _ := range takenPieces {
			// pieces with no ID (normal chess pieces) have no effect in this function
			g.DB.removePiece(id)
		}
		if after.Checkmate(nextOrientation) {
			if nextOrientation == wichessing.White {
				g.DB.updatePlayerRecords(g.Black, g.White, false)
			} else {
				g.DB.updatePlayerRecords(g.White, g.Black, false)
			}
		} else if ((g.DrawTurns + 1) >= draw_turn_count) || after.Draw(nextOrientation) {
			g.DB.updatePlayerRecords(g.White, g.Black, true)
		}
	} else {
		// in non-competitive games the collectible pieces maintain a reference count, reduced by one when removed
		tx := g.DB.Begin()
		for id, _ := range takenPieces {
			tx.unreservePiece(id)
		}
		tx.Commit()
	}
	return diff, promoting, promotingOrientation
}

func (g game) promote(from int, player string, kind wichessing.Kind, tx TX) map[string]piece {
	if g.Active != player {
		if debug {
			fmt.Println("promote: active not promoting player")
		}
		return nil
	}
	var nextMover string
	var orientation wichessing.Orientation
	if g.White == player {
		nextMover = g.Black
		orientation = wichessing.White
		if (from < 56) || (from > 63) {
			if debug {
				fmt.Println("promote: white out of range")
			}
			return nil
		}
	} else {
		nextMover = g.White
		orientation = wichessing.Black
		if (from < 0) || (from > 7) {
			if debug {
				fmt.Println("promote: black out of range")
			}
			return nil
		}
	}
	point := g.Points[from]
	if point.Kind == 0 {
		if debug {
			fmt.Println("promote: no piece")
		}
		return nil
	}
	if (point.Orientation != orientation) || (point.Base != wichessing.Pawn) {
		if debug {
			fmt.Println("promote: not right orientation or not pawn")
		}
		return nil
	}
	b := wichessingBoard(g.Points, g.From, g.To)
	if b.Draw(orientation) || g.Conceded {
		if debug {
			fmt.Println("promote: draw or conceded determined")
		}
		return nil
	}
	diff := make(map[string]piece)
	promdiff := b.PromotePawn(wichessing.AbsPointFromIndex(uint8(from)), wichessing.Kind(kind))
	if (promdiff == nil) || (len(promdiff) == 0) {
		if debug {
			fmt.Println("promote: b.PromotePawn returned nil or zero length diff")
		}
		return diff
	}
	for _, point := range promdiff {
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
	// guard pawn case, if previous mover was other player then the promoting player gets a move after this one
	if g.PreviousActive == nextMover {
		tx.updateGame(g.ID, diff, g.Active, g.Active, from, from, 0, g.Turn)
	} else {
		tx.updateGame(g.ID, diff, nextMover, g.Active, from, from, 0, g.Turn)
	}
	go func() {
		gameMonitorsLock.RLock()
		c, has := gameMonitors[g.ID]
		if has {
			c.move <- time.Now()
		}
		gameMonitorsLock.RUnlock()
	}()
	if (player != easy_computer_player) && (player != hard_computer_player) {
		gameListeningLock.RLock()
		listeners, has := gameListening[g.ID]
		if (orientation == wichessing.White) && has {
			if listeners.black != nil {
				listeners.black <- diff
			}
		}
		if (orientation == wichessing.Black) && has {
			if listeners.white != nil {
				listeners.white <- diff
			}
		}
		gameListeningLock.RUnlock()
	}
	if g.Competitive {
		var checkOrientation wichessing.Orientation
		if orientation == wichessing.White {
			checkOrientation = wichessing.Black
		} else {
			checkOrientation = wichessing.White
		}
		after := b.AfterPromote(absPoint(from), wichessing.Kind(kind))
		if after.Checkmate(checkOrientation) {
			if checkOrientation == wichessing.White {
				g.DB.updatePlayerRecords(g.Black, g.White, false)
			} else {
				g.DB.updatePlayerRecords(g.White, g.Black, false)
			}
		} else if after.Draw(checkOrientation) {
			g.DB.updatePlayerRecords(g.White, g.Black, true)
		}
	}
	return diff
}

// The map keys are wichessing.AbsPoint converted to "x/file-y/rank" formatted string. If the game is in a check or checkmate state, or a piece is to be promoted, or the active player's elapsed time has exceeded the total clock, then a corresponding key with a nil value will be set.
func (g game) moves(total time.Duration) map[string]map[string]struct{} {
	moves := make(map[string]map[string]struct{})
	active := g.activeOrientation()
	var opponent wichessing.Orientation
	if active == wichessing.White {
		opponent = wichessing.Black
	} else {
		opponent = wichessing.White
	}
	if g.Conceded {
		if g.Competitive {
			moves[conceded_key] = map[string]struct{}{
				fmt.Sprintf("%d", g.Piece): {},
			}
		} else {
			moves[conceded_key] = nil
		}
		return moves
	}
	tx := g.DB.Begin()
	if g.timeLoss(active, total, tx) || g.timeLoss(opponent, total, tx) {
		tx.Commit()
		if g.Competitive {
			moves[time_key] = map[string]struct{}{
				fmt.Sprintf("%d", g.Piece): {},
			}
		} else {
			moves[time_key] = nil
		}
		return moves
	}
	tx.Commit()
	board := wichessingBoard(g.Points, g.From, g.To)
	promoting, _ := board.HasPawnToPromote()
	if promoting {
		moves[promote_key] = nil
		return moves
	}
	if (g.DrawTurns >= draw_turn_count) || board.Draw(active) {
		if g.Competitive {
			moves[draw_key] = map[string]struct{}{
				fmt.Sprintf("%d", g.Piece): {},
			}
		} else {
			moves[draw_key] = nil
		}
		return moves
	}
	m, check, checkmate := board.Moves(active)
	for point, set := range m {
		strings := set.Strings()
		out := make(map[string]struct{})
		for _, str := range strings {
			out[str] = struct{}{}
		}
		moves[point.String()] = out
	}
	if checkmate {
		if g.Competitive {
			moves[checkmate_key] = map[string]struct{}{
				fmt.Sprintf("%d", g.Piece): {},
			}
		} else {
			moves[checkmate_key] = nil
		}
	} else if check {
		moves[check_key] = nil
	}
	if debug {
		if ((m == nil) || (len(m) == 0)) && (checkmate == false) && (check == false) {
			fmt.Println("moves: board.Moves returned nil or zero length set and not check/checkmate")
		}
	}
	return moves
}
