// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"github.com/pciet/wichess/rating"
)

const (
	player_table = "players"

	player_name_key   = "name"
	player_crypt_key  = "crypt"
	player_wins_key   = "wins"
	player_losses_key = "losses"
	player_draws_key  = "draws"
	player_rating_key = "rating"
	player_c5_key     = "c5"
	player_c15_key    = "c15"
)

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
	_, err := db.Exec("INSERT INTO "+player_table+"("+player_name_key+", "+player_crypt_key+", "+player_wins_key+", "+player_losses_key+", "+player_draws_key+", "+player_rating_key+", "+player_c5_key+", "+player_c15_key+") VALUES ($1, $2, $3, $4, $5, $6, $7, $8);", name, crypt, 0, 0, 0, rating.Initial, 0, 0)
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
