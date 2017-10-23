// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"strconv"

	"github.com/pciet/wichess/wichessing"
)

const (
	identifier_bit  = 0
	identifier_mask = 0xFFFFFFFF

	orientation_bit  = 32
	orientation_mask = 0x1

	moved_bit  = 33
	moved_mask = 0x1

	previous_point_bit  = 34
	previous_point_mask = 0x3F

	kind_bit  = 47
	kind_mask = 0xFFFF
)

type pieceEncoding uint64

func (p piece) encode() pieceEncoding {
	var enc uint64
	enc |= (uint64(p.Identifier) & identifier_mask) << identifier_bit
	enc |= (uint64(p.Orientation) & orientation_mask) << orientation_bit
	enc |= (uint64(btoi(p.Moved)) & moved_mask) << moved_bit
	enc |= (uint64(p.Kind) & kind_mask) << kind_bit
	enc |= (uint64(p.Previous) & previous_point_mask) << previous_point_bit
	return pieceEncoding(enc)
}

func (e pieceEncoding) decode() piece {
	return piece{
		Identifier: int((e >> identifier_bit) & identifier_mask),
		Piece: wichessing.Piece{
			Orientation: wichessing.Orientation((e >> orientation_bit) & orientation_mask),
			Moved:       itob(int((e >> moved_bit) & moved_mask)),
			Kind:        wichessing.Kind((e >> kind_bit) & kind_mask),
			Base:        wichessing.BaseForKind(wichessing.Kind((e >> kind_bit) & kind_mask)),
			Previous:    uint8(((e >> previous_point_bit) & previous_point_mask)),
		},
	}
}

func (e pieceEncoding) String() string {
	return strconv.Itoa(int(e))
}

func btoi(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

func itob(i int) bool {
	if i == 0 {
		return false
	} else {
		return true
	}
}
