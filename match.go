// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import ()

const (
	easy_computer_player = "Easy Computer Player"
	hard_computer_player = "Hard Computer Player"
)

// A zero for an ID value means use the regular piece. Index starts at left pawn and ends at right rook.
type gameSetup [16]int

func requestEasyComputerMatch(player string, setup gameSetup) {
	newGameIntoDatabase(player, setup, easy_computer_player, gameSetup{})
}
