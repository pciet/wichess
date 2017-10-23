// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

func (b Board) Draw(turn Orientation) bool {
	// TODO: no capture or pawn move in the last fifty moves by either player
	// not in check but no legal move
	moves, check, checkmate := b.Moves(turn)
	if (check == false) && (checkmate == false) && ((moves == nil) || (len(moves) == 0)) {
		return true
	}
	// TODO: same board position has occurred three times
	// insufficient material: king v king, king v king+bishop, king v king+knight, king+bishop v king+bishop of the same bishop color
	friendlyPieces := b.PiecesFor(turn)
	var opponentPieces PieceSet
	if turn == White {
		opponentPieces = b.PiecesFor(Black)
	} else {
		opponentPieces = b.PiecesFor(White)
	}
	// king v king
	if (len(friendlyPieces) == 1) && (len(opponentPieces) == 1) {
		return true
	}
	// insufficient materal only applies to two or less pieces per side
	if (len(friendlyPieces) > 2) || (len(opponentPieces) > 2) {
		return false
	}
	var friendly1, friendly2 *Piece
	var opponent1, opponent2 *Piece
	i := 0
	for piece, _ := range friendlyPieces {
		if i == 0 {
			friendly1 = piece
		} else {
			friendly2 = piece
		}
		i++
	}
	i = 0
	for piece, _ := range opponentPieces {
		if i == 0 {
			opponent1 = piece
		} else {
			opponent2 = piece
		}
		i++
	}
	if friendly2 == nil {
		// king v king+bishop
		if ((opponent1.Base == King) || (opponent1.Base == Bishop)) &&
			((opponent2.Base == King) || (opponent2.Base == Bishop)) {
			return true
		}
		// king v king+knight
		if ((opponent1.Base == King) || (opponent1.Base == Knight)) &&
			((opponent2.Base == King) || (opponent2.Base == Knight)) {
			return true
		}
	} else if ((friendly1.Base == King) || (friendly1.Base == Bishop)) &&
		((friendly2.Base == King) || (friendly2.Base == Bishop)) {
		// king+bishop v king+bishop, where the bishops are on the same color
		if ((opponent1.Base == King) || (opponent1.Base == Bishop)) &&
			((opponent2.Base == King) || (opponent2.Base == Bishop)) {
			var bishop1, bishop2 *Piece
			if friendly1.Base == Bishop {
				bishop1 = friendly1
			} else {
				bishop1 = friendly2
			}
			if opponent1.Base == Bishop {
				bishop2 = opponent1
			} else {
				bishop2 = opponent2
			}
			// TODO: this is the wrong way to check for bishops on the same color
			if (b.IndexPositionOfPiece(bishop1) % 2) == (b.IndexPositionOfPiece(bishop2) % 2) {
				return true
			}
		}
	}
	return false
}
