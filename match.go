// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"log"
)

// An identifier of 0 means request the normal basic piece instead of a hero piece.
func requestMatch(name string, leftRookID, leftKnightID, leftBishopID, rightBishopID, rightKnightID, rightRookID int) {
	log.Printf("%v requesting %v %v %v %v %v %v", name, leftRookID, leftKnightID, leftBishopID, rightBishopID, rightKnightID, rightRookID)
}
