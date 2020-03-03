package rules

type State int

const (
	Normal State = iota
	Promotion
	Check
	Checkmate
	Draw
	Conceded // this and the following aren't used by the rules package
	TimeOver
)
