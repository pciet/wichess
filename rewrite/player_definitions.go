package main

import (
	"fmt"
)

const (
	player_table = "players"

	player_name   = "name"
	player_crypt  = "crypt"
	player_wins   = "wins"
	player_losses = "losses"
	player_draws  = "draws"
	player_rating = "rating"
	player_c5     = "c5"
	player_c15    = "c15"

	player_friend_prefix = "f"

	friend_table = "friends"

	friend_requester = "requester"
	friend_setup     = "setup"
	friend_friend    = "friend"
	friend_slot      = "slot"

	newPlayerPieceCount = 3
)

type PlayerRecord struct {
	Wins   int
	Losses int
	Draws  int
}

var player_record_query = BuildSQLQuery([]string{
	player_wins,
	player_losses,
	player_draws,
}, player_table, player_name)

type PlayerFriendStatus struct {
	ActiveFriends   [6]string
	MatchingFriends [6]string
	ActiveTurn      [6]bool
}

var player_friend_games_query = BuildSQLQuery(func() [6]string {
	var selects [6]string
	for i := 0; i < 6; i++ {
		selects[i] = fmt.Sprintf("%s%d", player_friend_prefix, i)
	}
	return selects
}(), player_table, player_name)

var friend_matching_query = BuildSQLQuery([]string{
	friend_friend,
	friend_slot,
}, friend_table, friend_requester)

var player_timed_game_query = BuildSQLQuery([]string{
	player_c5,
	player_c15,
}, player_table, player_name)

var player_crypt_query = BuildSQLQuery(player_crypt_key, player_table, player_name_key)
