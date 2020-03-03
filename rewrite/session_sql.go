package main

import ()

const (
	session_key_cookie = "k"
	session_key_length = 64

	session_table = "sessions"
	session_name  = "name"
	session_key   = "key"
)

var (
	session_name_query   = BuildSQLQuery([]string{session_name}, session_table, session_key)
	session_exists_query = BuildSQLForUpdateQuery(nil, session_table, session_name)
	session_key_query    = BuildSQLForUpdateQuery([]string{session_key}, session_table, session_name)
	session_insert       = BuildSQLInsert(session_table, []string{session_name, session_key})
	session_update       = BuildSQLUpdate(session_table, session_key, session_name)
)
