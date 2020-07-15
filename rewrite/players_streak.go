package main

import "database/sql"

// PlayerComputerStreak returns the current and best win streak against the computer player.
func PlayerComputerStreak(tx *sql.Tx, id PlayerIdentifier) (int, int) {
	var streak, best int
	err := tx.QueryRow(PlayersComputerStreakQuery, id).Scan(&streak, &best)
	if err == sql.ErrNoRows {
		return 0, 0
	} else if err != nil {
		Panic(err)
	}
	return streak, best
}

func PlayerComputerStreakIncrement(tx *sql.Tx, id PlayerIdentifier) {
	current, best := PlayerComputerStreak(tx, id)
	current++
	if current > best {
		PlayerUpdateBestComputerStreak(tx, id, current)
	}
	PlayerUpdateComputerStreak(tx, id, current)
}

func PlayerResetComputerStreak(tx *sql.Tx, id PlayerIdentifier) {
	PlayerUpdateComputerStreak(tx, id, 0)
}

func PlayerUpdateBestComputerStreak(tx *sql.Tx, id PlayerIdentifier, value int) {
	SQLExecRow(tx, PlayersBestComputerStreakUpdate, value, id)
}

func PlayerUpdateComputerStreak(tx *sql.Tx, id PlayerIdentifier, value int) {
	SQLExecRow(tx, PlayersComputerStreakUpdate, value, id)
}
