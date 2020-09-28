package rules

import "strconv"

// State represents the current state of the game, such as check, promotion needed, or draw.
type State int

const (
	Normal State = iota
	Promotion
	Check
	Checkmate
	Draw
	Conceded // this and the following aren't used by the rules package
	TimeOver
	ReversePromotion
	NoState
)

func (a State) String() string {
	switch a {
	case Normal:
		return "normal"
	case Promotion:
		return "promotion"
	case Check:
		return "check"
	case Checkmate:
		return "checkmate"
	case Draw:
		return "draw"
	case Conceded:
		return "conceded"
	case TimeOver:
		return "time over"
	case ReversePromotion:
		return "reverse promotion"
	}
	return strconv.Itoa(int(a))
}
