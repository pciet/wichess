package main

import "database/sql"

type PlayerFriendStatus struct {
	ActiveFriends   [PlayerFriendGameCount]string
	MatchingFriends [PlayerFriendGameCount]string
	ActiveTurn      [PlayerFriendGameCount]bool
}

func LoadPlayerFriendStatus(tx *sql.Tx, name string) PlayerFriendStatus {
	var s PlayerFriendStatus
	s.ActiveTurn, s.ActiveFriends = PlayerFriendActiveAndOpponentName(tx, name)
	s.MatchingFriends = FriendMatchings(tx, name)
	return s
}

// PlayerFriendActiveAndOpponentName shows if a friend is waiting for the player's move in their
// six friend games, and it also returns the friend names.
func PlayerFriendActiveAndOpponentName(tx *sql.Tx, name string) ([PlayerFriendGameCount]bool, [PlayerFriendGameCount]string) {
	games := PlayerFriendGames(tx, name)
	var active [PlayerFriendGameCount]bool
	var opponents [PlayerFriendGameCount]string
	for i, g := range games {
		if g == 0 {
			continue
		}
		active[i], opponents[i] = GameActiveAndOpponentName(tx, g, name)
	}
	return active, opponents
}

func PlayerFriendGames(tx *sql.Tx, name string) [PlayerFriendGameCount]GameIdentifier {
	var games [PlayerFriendGameCount]GameIdentifier
	s := make([]interface{}, PlayerFriendGameCount)
	for i := 0; i < PlayerFriendGameCount; i++ {
		s[i] = &games[i]
	}
	err := tx.QueryRow(PlayerFriendGamesQuery, name).Scan(s...)
	if err != nil {
		Panic(err)
	}
	return games
}
