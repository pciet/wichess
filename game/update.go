package game

import "github.com/pciet/wichess/rules"

// TODO: define CapturedPiece at the right place

type (
	// The UpdateState of an Update is the game state or signals about promotion as described in
	// docs/promotion.md.
	UpdateState string

	// Updates are the information sent in an Alert. This can be changes to the game such as board
	// changes and captures, notification of a check, draw, or checkmate calculation, or an alert
	// that the opponent conceded.
	Update struct {
		Squares     []Square `json:"d,omitempty"`
		Captures    []Piece  `json:"c,omitempty"`
		UpdateState `json:"s,omitempty"`
		FromMove    rules.Move `json:"m"` // can't be empty, rules.NoMove for empty
	}
)

// These const strings are the possible values for an UpdateState.
const (
	PromotionNeededUpdate     = "p"
	WaitUpdate                = "w"
	ContinueUpdate            = "c"
	CheckCalculatedUpdate     = "ch"
	DrawCalculatedUpdate      = "dr"
	CheckmateCalculatedUpdate = "chm"
	ConcededUpdate            = "co"
)
