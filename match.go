// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import ()

// A zero for an ID value means use the regular piece. Index starts at left pawn and ends at right rook.
type gameSetup [16]int

func (db DB) requestEasyComputerMatch(player string, setup gameSetup) {
	db.newGame(player, setup, easy_computer_player, gameSetup{})
}
