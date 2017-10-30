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
)

// Returns address that have changed, Kind 0 piece for a now empty point, but does not update the board points. Returns true if promoting.
func (g game) move(from, to int, mover string, timeoutMove bool) (map[string]piece, bool) {
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
	b := wichessingBoard(g.Points)
	if b.HasPawnToPromote() {
		return nil, false
	}
	if b.Draw(orientation) {
		return nil, false
	}
	diff := make(map[string]piece)
	difference, taken := b.Move(absPoint(from), absPoint(to), orientation)
	for point, _ := range difference {
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
			for index, ppt := range b {
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
		return diff, false
	}
	takenPieces := make(map[int]struct{})
	for point, _ := range taken {
		takenPieces[g.Points[point.Index()].Identifier] = struct{}{}
	}
	after := b.AfterMove(absPoint(from), absPoint(to), orientation)
	var promoting bool
	if after.HasPawnToPromote() {
		g.DB.updateGame(g.ID, diff, mover)
		promoting = true
	} else {
		g.DB.updateGame(g.ID, diff, nextMover)
	}
	if timeoutMove == false {
		go func() {
			gameMonitorsLock.RLock()
			c, has := gameMonitors[g.ID]
			if has {
				c.move <- time.Now()
			}
			gameMonitorsLock.RUnlock()
		}()
	}
	if (mover != easy_computer_player) && (mover != hard_computer_player) {
		gameListeningLock.RLock()
		listeners, has := gameListening[g.ID]
		if ((orientation == wichessing.White) || timeoutMove) && has {
			if listeners.black != nil {
				listeners.black <- diff
			}
		}
		if ((orientation == wichessing.Black) || timeoutMove) && has {
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
		} else if after.Draw(nextOrientation) {
			g.DB.updatePlayerRecords(g.White, g.Black, true)
		}
	}
	return diff, promoting
}

func (g game) promote(from int, player string, kind wichessing.Kind, timeoutMove bool) map[string]piece {
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
	if (point.Orientation != orientation) || (point.Base != wichessing.Pawn) {
		return nil
	}
	b := wichessingBoard(g.Points)
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
	g.DB.updateGame(g.ID, diff, nextMover)
	if timeoutMove == false {
		go func() {
			gameMonitorsLock.RLock()
			c, has := gameMonitors[g.ID]
			if has {
				c.move <- time.Now()
			}
			gameMonitorsLock.RUnlock()
		}()
	}
	if (player != easy_computer_player) && (player != hard_computer_player) {
		gameListeningLock.RLock()
		listeners, has := gameListening[g.ID]
		if ((orientation == wichessing.White) || timeoutMove) && has {
			if listeners.black != nil {
				listeners.black <- diff
			}
		}
		if ((orientation == wichessing.Black) || timeoutMove) && has {
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
	if g.timeLoss(active, total) {
		if g.Competitive {
			moves[time_key] = map[string]struct{}{
				fmt.Sprintf("%d", g.Piece): {},
			}
		} else {
			moves[time_key] = nil
		}
		return moves
	}
	board := wichessingBoard(g.Points)
	if board.HasPawnToPromote() {
		moves[promote_key] = nil
		return moves
	}
	if board.Draw(active) {
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
		moves[point.String()] = set.Strings()
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
	return moves
}
