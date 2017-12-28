// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"fmt"

	"github.com/pciet/wichess/rating"
)

const (
	player_table = "players"

	player_name_key      = "name"
	player_crypt_key     = "crypt"
	player_wins_key      = "wins"
	player_losses_key    = "losses"
	player_draws_key     = "draws"
	player_rating_key    = "rating"
	player_c5_key        = "c5"
	player_c15_key       = "c15"
	player_friend_prefix = "f"
)

func (db DB) freePlayersFriendSlot(name string, slot uint8) {
	result, err := db.Exec("UPDATE "+player_table+" SET "+player_friend_prefix+fmt.Sprintf("%d", slot)+" = $1 WHERE "+player_name_key+" = $2;", 0, name)
	if err != nil {
		panic(err.Error())
	}
	count, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	if count != 1 {
		panic(fmt.Sprint(count, " rows affected by update to free player's friend slot"))
	}
}

// Returns -1 if the game is not in one of the slots.
func (db DB) playersFriendSlotForGame(name string, id int) int {
	var games [6]int
	err := db.QueryRow("SELECT "+player_friend_prefix+"0, "+player_friend_prefix+"1, "+player_friend_prefix+"2, "+player_friend_prefix+"3, "+player_friend_prefix+"4, "+player_friend_prefix+"5 FROM "+player_table+" WHERE "+player_name_key+"=$1;", name).Scan(&games[0], &games[1], &games[2], &games[3], &games[4], &games[5])
	if err != nil {
		panic(err.Error())
	}
	slot := -1
	for s, gid := range games {
		if gid == id {
			slot = s
			break
		}
	}
	return slot
}

func (db DB) playersGameFromFriendSlot(name string, slot uint8) int {
	var id int
	err := db.QueryRow("SELECT "+player_friend_prefix+fmt.Sprintf("%d", slot)+" FROM "+player_table+" WHERE "+player_name_key+"=$1;", name).Scan(&id)
	if err != nil {
		panic(err.Error())
	}
	return id
}

func (db DB) playersFriendSlotOpponentsAndActive(name string) ([6]string, [6]bool) {
	var games [6]int
	err := db.QueryRow("SELECT "+player_friend_prefix+"0, "+player_friend_prefix+"1, "+player_friend_prefix+"2, "+player_friend_prefix+"3, "+player_friend_prefix+"4, "+player_friend_prefix+"5 FROM "+player_table+" WHERE "+player_name_key+"=$1;", name).Scan(&games[0], &games[1], &games[2], &games[3], &games[4], &games[5])
	if err != nil {
		panic(err.Error())
	}
	var opponents [6]string
	var active [6]bool
	for i, id := range games {
		if id == 0 {
			continue
		}
		opponents[i], active[i] = db.gameOpponentAndActive(name, id)
	}
	return opponents, active
}

func (db DB) setPlayerFriendSlot(name string, slot uint8, gameID int) {
	if slot > 5 {
		panic(fmt.Sprint(slot, "friend slot greater than 5"))
	}
	result, err := db.Exec("UPDATE "+player_table+" SET "+player_friend_prefix+fmt.Sprintf("%d", slot)+" = $1 WHERE "+player_name_key+" = $2;", gameID, name)
	if err != nil {
		panic(err.Error())
	}
	count, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	if count != 1 {
		panic(fmt.Sprint(count, "instead of 1 rows affected"))
	}
}

func (db DB) playersCompetitive15HourGameID(player string) int {
	var id int
	err := db.QueryRow("SELECT "+player_c15_key+" FROM "+player_table+" WHERE "+player_name_key+"=$1;", player).Scan(&id)
	if err != nil {
		panicExit(err.Error())
	}
	return id
}

func (db DB) setPlayerCompetitive15Game(player string, gameID int) {
	_, err := db.Exec("UPDATE "+player_table+" SET "+player_c15_key+" = $1 WHERE "+player_name_key+" = $2;", gameID, player)
	if err != nil {
		panicExit(err.Error())
	}
}

func (db DB) removePlayersCompetitive15Game(player string) {
	_, err := db.Exec("UPDATE "+player_table+" SET "+player_c15_key+" = $1 WHERE "+player_name_key+" = $2;", 0, player)
	if err != nil {
		panicExit(err.Error())
	}
}

func (db DB) playersCompetitive5HourGameID(player string) int {
	var id int
	err := db.QueryRow("SELECT "+player_c5_key+" FROM "+player_table+" WHERE "+player_name_key+"=$1;", player).Scan(&id)
	if err != nil {
		panicExit(err.Error())
	}
	return id
}

func (db DB) setPlayerCompetitive5Game(player string, gameID int) {
	_, err := db.Exec("UPDATE "+player_table+" SET "+player_c5_key+" = $1 WHERE "+player_name_key+" = $2;", gameID, player)
	if err != nil {
		panicExit(err.Error())
	}
}

func (db DB) removePlayersCompetitive5Game(player string) {
	_, err := db.Exec("UPDATE "+player_table+" SET "+player_c5_key+" = $1 WHERE "+player_name_key+" = $2;", 0, player)
	if err != nil {
		panicExit(err.Error())
	}
}

func (db DB) playerRating(name string) int {
	var rating int
	err := database.QueryRow("SELECT "+player_rating_key+" FROM "+player_table+" WHERE "+player_name_key+"=$1;", name).Scan(&rating)
	if err != nil {
		panicExit(err.Error())
	}
	return rating
}

const (
	player_record_win_update  = "UPDATE " + player_table + " SET " + player_wins_key + " = " + player_wins_key + " + 1, " + player_rating_key + " = $1 WHERE " + player_name_key + " = $2;"
	player_record_lose_update = "UPDATE " + player_table + " SET " + player_losses_key + " = " + player_losses_key + " + 1, " + player_rating_key + " = $1 WHERE " + player_name_key + " = $2;"
	player_record_draw_update = "UPDATE " + player_table + " SET " + player_draws_key + " = " + player_draws_key + " + 1, " + player_rating_key + " = $1 WHERE " + player_name_key + " = $2;"
)

func (db DB) updatePlayerRecords(winner, loser string, draw bool) {
	if debug {
		fmt.Printf("updating record, winner %v loser %v draw %v\n", winner, loser, draw)
	}
	if (winner == easy_computer_player) || (winner == hard_computer_player) || (loser == easy_computer_player) || (loser == hard_computer_player) {
		panicExit("updating record for computer game")
	}
	winnerRating := uint(db.playerRating(winner))
	loserRating := uint(db.playerRating(loser))
	var newWinnerRating, newLoserRating uint
	if draw {
		newWinnerRating = rating.Updated(winnerRating, loserRating, rating.Draw)
		newLoserRating = rating.Updated(loserRating, winnerRating, rating.Draw)
	} else {
		newWinnerRating = rating.Updated(winnerRating, loserRating, rating.Win)
		newLoserRating = rating.Updated(loserRating, winnerRating, rating.Loss)
	}
	if draw {
		_, err := db.Exec(player_record_draw_update, newWinnerRating, winner)
		if err != nil {
			panicExit(err.Error())
		}
		_, err = db.Exec(player_record_draw_update, newLoserRating, loser)
		if err != nil {
			panicExit(err.Error())
		}
	} else {
		_, err := db.Exec(player_record_win_update, newWinnerRating, winner)
		if err != nil {
			panicExit(err.Error())
		}
		_, err = db.Exec(player_record_lose_update, newLoserRating, loser)
		if err != nil {
			panicExit(err.Error())
		}
	}
}

func (db DB) playerCrypt(name string) (bool, string) {
	var c string
	err := database.QueryRow("SELECT "+player_crypt_key+" FROM "+player_table+" WHERE "+player_name_key+"=$1;", name).Scan(&c)
	if err != nil {
		return false, ""
	}
	return true, c
}

func (db DB) newPlayer(name, crypt string) {
	_, err := db.Exec("INSERT INTO "+player_table+"("+player_name_key+", "+player_crypt_key+", "+player_wins_key+", "+player_losses_key+", "+player_draws_key+", "+player_rating_key+", "+player_c5_key+", "+player_c15_key+", "+player_friend_prefix+"0, "+player_friend_prefix+"1, "+player_friend_prefix+"2, "+player_friend_prefix+"3, "+player_friend_prefix+"4, "+player_friend_prefix+"5) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14);", name, crypt, 0, 0, 0, rating.Initial, 0, 0, 0, 0, 0, 0, 0, 0)
	if err != nil {
		panicExit(err.Error())
	}
}

type record struct {
	wins   int
	losses int
	draws  int
}

func (db DB) playerRecord(name string) record {
	var wins, losses, draws int
	err := db.QueryRow("SELECT "+player_wins_key+", "+player_losses_key+", "+player_draws_key+" FROM "+player_table+" WHERE "+player_name_key+"=$1;", name).Scan(&wins, &losses, &draws)
	if err != nil {
		panicExit(err.Error())
	}
	return record{
		wins:   wins,
		losses: losses,
		draws:  draws,
	}
}
