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

	player_friend_game_prefix = "f"
	player_friend_game_count  = 6

	friend_table = "friends"

	friend_requester = "requester"
	friend_setup     = "setup"
	friend_friend    = "friend"
	friend_slot      = "slot"
)

var (
	FriendRows = func() []string {
		s := make([]string, player_friend_game_count)
		for i := 0; i < player_friend_game_count; i++ {
			s[i] = fmt.Sprintf("%s%d", player_friend_game_prefix, i)
		}
		return s
	}()

	player_new_insert = BuildSQLInsert(player_table, func() []string {
		s := []string{
			player_name,
			player_crypt,
			player_wins,
			player_losses,
			player_draws,
			player_rating,
			player_c5,
			player_c15,
		}
		for _, v := range FriendRows {
			s = append(s, v)
		}
		return s
	}())

	player_record_query = BuildSQLQuery([]string{
		player_wins,
		player_losses,
		player_draws,
	}, player_table, player_name)

	player_friend_games_query = BuildSQLQuery(FriendRows, player_table, player_name)

	friend_matching_query = BuildSQLQuery([]string{
		friend_friend,
		friend_slot,
	}, friend_table, friend_requester)

	player_timed_game_query = BuildSQLQuery([]string{
		player_c5,
		player_c15,
	}, player_table, player_name)

	player_crypt_query = BuildSQLQuery([]string{player_crypt}, player_table, player_name)
)
