package wichess

const (
	PlayersTable = "players"

	PlayersIdentifier = "id"
	PlayersName       = "name"

	PlayersCrypt   = "crypt"
	PlayersSession = "session"

	PlayersRecentOpponents = "recent_opponents"
	PlayersPeopleGame      = "people_game"

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
	PlayersRecentOpponentsUpdate = SQLUpdate(PlayersTable,
		PlayersRecentOpponents, PlayersIdentifier)

	PlayersPeopleGameQuery  = SQLQuery([]string{PlayersPeopleGame}, PlayersTable, PlayersIdentifier)
	PlayersPeopleGameUpdate = SQLUpdate(PlayersTable, PlayersPeopleGame, PlayersIdentifier)

	PlayersComputerStreakQuery = SQLQuery([]string{PlayersComputerStreak,
		PlayersBestComputerStreak}, PlayersTable, PlayersIdentifier)
	PlayersComputerStreakUpdate = SQLUpdate(PlayersTable,
		PlayersComputerStreak, PlayersIdentifier)
	PlayersBestComputerStreakUpdate = SQLUpdate(PlayersTable,
		PlayersBestComputerStreak, PlayersIdentifier)
)
