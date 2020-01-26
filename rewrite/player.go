package main

import (
	"database/sql"
)

func PlayerRecord(tx *sql.Tx, name string) PlayerRecord {
	var r PlayerRecord
	err := tx.QueryRow(player_record_query, name).Scan(
		&r.Wins,
		&r.Losses,
		&r.Draws,
	)
	if err != nil {
		panic(err)
	}
	return r
}

func PlayerFriendStatus(tx *sql.Tx, name string) PlayerFriendStatus {
	var s PlayerFriendStatus
	s.ActiveTurn, s.ActiveFriends = PlayerFriendActiveAndOpponentName(tx, name)
	s.MatchingFriends = PlayerFriendMatching(tx, name)
	return s
}

func PlayerFriendMatching(tx *sql.Tx, name string) [6]string {
	rows, err := tx.Query(friend_matching_query, name)
	if err != nil {
		panic(err)
	}
	var friends [6]string
	for rows.Next() {
		var friend string
		var slot uint8
		err = rows.Scan(&friend, &slot)
		if err != nil {
			panic(err)
		}
		friends[slot] = friend
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return friends
}

// Returns if this player is active (the game is waiting for their move) for their six friend games, and the names of the player's opponents in those games.
func PlayerFriendActiveAndOpponentName(tx *sql.Tx, name string) ([6]bool, [6]string) {
	games := PlayerFriendGames(tx, name)
	var active [6]bool
	var opponents [6]string
	for i, g := range games {
		if g == 0 {
			continue
		}
		active[i], opponents[i] = GameActiveAndOpponentName(tx, g, name)
	}
	return active, opponents
}

func PlayerFriendGames(tx *sql.Tx, name string) [6]GameIdentifier {
	var games [6]GameIdentifier
	err := tx.QueryRow(player_friend_games_query, name).Scan(
		&(games[0]),
		&(games[1]),
		&(games[2]),
		&(games[3]),
		&(games[4]),
		&(games[5]),
	)
	if err != nil {
		panic(err)
	}
	return games
}

func PlayerHasTimedGame(tx *sql.Tx, name string) bool {
	t5, t15 := PlayerTimedGameIdentifiers(tx, name)
	if (t5 != 0) || (t15 != 0) {
		return true
	}
	return false
}

// The first return is the 5 minute game ID and the second is the 15 minute.
func PlayerTimedGameIdentifiers(tx *sql.Tx, name string) (GameIdentifier, GameIdentifier) {
	var t5, t15 GameIdentifier
	err := tx.QueryRow(player_timed_game_query, name).Scan(&t5, &t15)
	if err != nil {
		panic(err)
	}
	return t5, t15
}
