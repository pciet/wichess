package main

import ()

const (
	key_cookie = "k"

	key_length = 64

	session_table = "sessions"
	session_name  = "name"
	session_key   = "key"
)

var session_player_name_query = BuildSQLQuery(session_name, session_table, session_key)
