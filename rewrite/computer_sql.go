package main

var computer_game_id_query = BuildSQLGeneralizedWhereQuery([]string{games_identifier}, games_table,
	games_white+" = $1 AND "+games_black+" = '"+computer_player_name+"'")
