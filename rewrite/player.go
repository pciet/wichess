package main

import (
	"database/sql"
	"log"

	"github.com/pciet/wichess/rules"
)

func NewPlayer(tx *sql.Tx, name, crypt string) {
	result, err := tx.Exec(player_new_insert,
		name,
		crypt,
		0, 0, 0,
		InitialRating,
		0, 0, 0, 0, 0, 0, 0, 0)
	if err != nil {
		log.Panic(err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Panic(err)
	}
	if count != 1 {
		log.Panicln(count, "rows affected by new player insert for", name)
	}

	for i := 0; i < NewPlayerPieceCount; i++ {
		InsertNewPiece(tx, rules.RandomSpecialPieceKind(), name)
	}
}

func LoadPlayerRecord(tx *sql.Tx, name string) PlayerRecord {
	var r PlayerRecord
	err := tx.QueryRow(player_record_query, name).Scan(
		&r.Wins,
		&r.Losses,
		&r.Draws,
	)
	if err != nil {
		log.Panic(err)
	}
	return r
}

func LoadPlayerFriendStatus(tx *sql.Tx, name string) PlayerFriendStatus {
	var s PlayerFriendStatus
	s.ActiveTurn, s.ActiveFriends = PlayerFriendActiveAndOpponentName(tx, name)
	s.MatchingFriends = PlayerFriendMatching(tx, name)
	return s
}

func PlayerFriendMatching(tx *sql.Tx, name string) [player_friend_game_count]string {
	rows, err := tx.Query(friend_matching_query, name)
	if err != nil {
		log.Panic(err)
	}
	var friends [player_friend_game_count]string
	for rows.Next() {
		var friend string
		var slot uint8
		err = rows.Scan(&friend, &slot)
		if err != nil {
			log.Panic(err)
		}
		friends[slot] = friend
	}
	err = rows.Err()
	if err != nil {
		log.Panic(err)
	}
	return friends
}

// Returns if this player is active (the game is waiting for their move) for their six friend games, and the names of the player's opponents in those games.
func PlayerFriendActiveAndOpponentName(tx *sql.Tx, name string) ([player_friend_game_count]bool, [player_friend_game_count]string) {
	games := PlayerFriendGames(tx, name)
	var active [player_friend_game_count]bool
	var opponents [player_friend_game_count]string
	for i, g := range games {
		if g == 0 {
			continue
		}
		active[i], opponents[i] = GameActiveAndOpponentName(tx, g, name)
	}
	return active, opponents
}

func PlayerFriendGames(tx *sql.Tx, name string) [player_friend_game_count]GameIdentifier {
	var games [player_friend_game_count]GameIdentifier
	s := make([]interface{}, player_friend_game_count)
	for i := 0; i < player_friend_game_count; i++ {
		s[i] = &games[i]
	}
	err := tx.QueryRow(player_friend_games_query, name).Scan(s...)
	if err != nil {
		log.Panic(err)
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
		log.Panic(err)
	}
	return t5, t15
}
