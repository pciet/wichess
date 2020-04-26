package main

const (
	PlayerTable = "players"

	PlayerName  = "name"
	PlayerCrypt = "crypt"
)

var (
	PlayerNewInsert = SQLInsert(PlayerTable, []string{
		PlayerName,
		PlayerCrypt,
	})

	PlayerCryptQuery = SQLQuery([]string{PlayerCrypt}, PlayerTable, PlayerName)
)
