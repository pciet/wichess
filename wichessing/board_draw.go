// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

func (b Board) Draw(turn Orientation) bool {
	// not in check but no legal move
	checkmate := b.Checkmate(turn)
	check := b.Check(turn)
	moves := b.AllNaiveMovesFor(turn)
	// remove all check moves
	for pt, mvs := range moves {
		for _, mv := range mvs {
			if b.AfterMove(pt, mv, turn).Check(turn) {
				mvs = mvs.Remove(mv)
			}
			if len(mvs) == 0 {
				delete(moves, pt)
			}
		}
	}
	if (check == false) && (checkmate == false) && ((moves == nil) || (len(moves) == 0)) {
		promoting, promotingOrientation := b.HasPawnToPromote()
		if (promoting == false) || (promotingOrientation != turn) {
			return true
		}
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
	for _, piece := range friendlyPieces {
		if i == 0 {
			friendly1 = piece
		} else {
			friendly2 = piece
		}
		i++
	}
	i = 0
	for _, piece := range opponentPieces {
		if i == 0 {
			opponent1 = piece
		} else {
			opponent2 = piece
		}
		i++
	}
	if b.insufficientMaterial(friendly1, friendly2, opponent1, opponent2) || b.insufficientMaterial(opponent1, opponent2, friendly1, friendly2) {
		return true
	}
	return false
}

// this function is called twice, it does not cover all cases in both directions
func (b Board) insufficientMaterial(friendly1, friendly2, opponent1, opponent2 *Piece) bool {
	// this case is covered by calling this function the second time reversed
	if opponent2 == nil {
		return false
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
			var bishop1Colored, bishop2Colored bool
			b1i := b.IndexPositionOfPiece(bishop1)
			// if on an even row then colored is if odd
			if ((b1i / 8) % 2) == 1 {
				if (b1i % 2) == 0 {
					bishop1Colored = true
				} else {
					bishop1Colored = false
				}
			} else {
				if (b1i % 2) == 0 {
					bishop1Colored = false
				} else {
					bishop1Colored = true
				}
			}
			b2i := b.IndexPositionOfPiece(bishop2)
			if ((b2i / 8) % 2) == 1 {
				if (b2i % 2) == 0 {
					bishop2Colored = true
				} else {
					bishop2Colored = false
				}
			} else {
				if (b2i % 2) == 0 {
					bishop2Colored = false
				} else {
					bishop2Colored = true
				}
			}
			if bishop1Colored == bishop2Colored {
				return true
			}
		}
	}
	return false
}
