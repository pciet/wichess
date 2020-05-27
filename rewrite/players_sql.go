package main

const (
	PlayersTable = "players"

	PlayersIdentifier = "id"
	PlayersName       = "name"

	PlayersCrypt   = "crypt"
	PlayersSession = "session"

	PlayersRecentOpponents = "recent_opponents"

	PlayersLeftKind   = "left_kind"
	PlayersRightKind  = "right_kind"
	PlayersCollection = "collection"

	PlayersComputerStreak     = "computer_streak"
	PlayersBestComputerStreak = "best_computer_streak"
)

var (
	PlayersNewInsert = SQLInsertReturning(PlayersTable, []string{
		PlayersName,
		PlayersCrypt,
		PlayersLeftKind,
		PlayersRightKind,
		PlayersCollection,
		PlayersComputerStreak,
		PlayersBestComputerStreak,
	}, PlayersIdentifier)

	PlayersNameQuery       = SQLQuery([]string{PlayersName}, PlayersTable, PlayersIdentifier)
	PlayersIdentifierQuery = SQLQuery([]string{PlayersIdentifier}, PlayersTable, PlayersName)

	PlayersCryptQuery = SQLQuery([]string{PlayersIdentifier, PlayersCrypt},
		PlayersTable, PlayersName)

	PlayersSessionUpdate = SQLUpdate(PlayersTable, PlayersSession, PlayersIdentifier)
	PlayersSessionQuery  = SQLQuery([]string{PlayersSession}, PlayersTable, PlayersIdentifier)

	PlayersPiecePicksQuery = SQLQuery([]string{
		PlayersLeftKind, PlayersRightKind}, PlayersTable, PlayersName)

	PlayersCollectionQuery = SQLQuery([]string{PlayersCollection}, PlayersTable, PlayersIdentifier)

	PlayersRecentOpponentsQuery = SQLQuery([]string{PlayersRecentOpponents},
		PlayersTable, PlayersIdentifier)
)
