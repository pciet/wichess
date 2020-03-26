package main

import "fmt"

const (
	PlayerTable = "players"

	PlayerName   = "name"
	PlayerCrypt  = "crypt"
	PlayerWins   = "wins"
	PlayerLosses = "losses"
	PlayerDraws  = "draws"
	PlayerRating = "rating"
	PlayerC5     = "c5"
	PlayerC15    = "c15"

	PlayerFriendGamePrefix = "f"
	PlayerFriendGameCount  = 6
)

var (
	PlayerFriendRows = func() []string {
		s := make([]string, PlayerFriendGameCount)
		for i := 0; i < PlayerFriendGameCount; i++ {
			s[i] = fmt.Sprintf("%s%d", PlayerFriendGamePrefix, i)
		}
		return s
	}()

	PlayerNewInsert = SQLInsert(PlayerTable, func() []string {
		s := []string{
			PlayerName,
			PlayerCrypt,
			PlayerWins,
			PlayerLosses,
			PlayerDraws,
			PlayerRating,
			PlayerC5,
			PlayerC15,
		}
		for _, v := range PlayerFriendRows {
			s = append(s, v)
		}
		return s
	}())

	PlayerRecordQuery = SQLQuery([]string{
		PlayerWins,
		PlayerLosses,
		PlayerDraws,
	}, PlayerTable, PlayerName)

	PlayerFriendGamesQuery = SQLQuery(PlayerFriendRows, PlayerTable, PlayerName)

	PlayerTimedGameQuery = SQLQuery([]string{
		PlayerC5,
		PlayerC15,
	}, PlayerTable, PlayerName)

	PlayerCryptQuery = SQLQuery([]string{PlayerCrypt}, PlayerTable, PlayerName)
)
