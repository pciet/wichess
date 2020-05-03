package main

const (
	PlayerTable = "players"

	PlayerName       = "name"
	PlayerCrypt      = "crypt"
	PlayerLeftPiece  = "left_piece"
	PlayerRightPiece = "right_piece"
)

var (
	PlayerNewInsert = SQLInsert(PlayerTable, []string{
		PlayerName,
		PlayerCrypt,
		PlayerLeftPiece,
		PlayerRightPiece,
	})

	PlayerCryptQuery = SQLQuery([]string{PlayerCrypt}, PlayerTable, PlayerName)

	PlayerPiecePicksQuery = SQLQuery([]string{
		PlayerLeftPiece, PlayerRightPiece}, PlayerTable, PlayerName)
)
