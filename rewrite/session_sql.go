package main

const (
	SessionTable = "sessions"
	SessionName  = "name"
	SessionKey   = "key"
)

var (
	SessionNameQuery   = SQLQuery([]string{SessionName}, SessionTable, SessionKey)
	SessionExistsQuery = SQLForUpdateQuery(nil, SessionTable, SessionName)
	SessionKeyQuery    = SQLForUpdateQuery([]string{SessionKey}, SessionTable, SessionName)
	SessionInsert      = SQLInsert(SessionTable, []string{SessionName, SessionKey})
	SessionUpdate      = SQLUpdate(SessionTable, SessionKey, SessionName)
)
