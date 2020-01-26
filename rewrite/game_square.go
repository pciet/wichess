package main

import (
	"github.com/pciet/wichess/rules"
)

type AddressedSquare struct {
	rules.Square
	rules.BoardAddress
}
