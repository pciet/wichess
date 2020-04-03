package main

import "github.com/pciet/wichess/rules"

// An Update is used by the response to /move and the
// alert message to update the webpage's board
// representation and to indicate promotion communication
// variations as described by docs/promotion.md.
type Update struct {
	Squares []rules.AddressedSquare `json:"d"` // diff
	State   string                  `json:"s,omitempty"`
}

const (
	PromotionNeededUpdate = "p"
	WaitUpdate            = "w"
	ContinueUpdate        = "c"
)
