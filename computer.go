package wichess

import "database/sql"

const (
	ComputerPlayerName = "Computer Player"
	ComputerPlayerID   = 0
)

var ComputerGameIdentifierQuery = SQLGeneralizedWhereQuery([]string{GamesIdentifier}, GamesTable,
	"("+GamesWhite+"=$1 AND "+GamesBlack+"='"+ComputerPlayerName+"') OR ("+
		GamesWhite+"='"+ComputerPlayerName+"' AND "+GamesBlack+"=$1)")

// ComputerGameIdentifier gets the game identifier for the player's one computer opponent game.
func ComputerGameIdentifier(tx *sql.Tx, player string) GameIdentifier {
	var id GameIdentifier
	err := tx.QueryRow(ComputerGameIdentifierQuery, player).Scan(&id)
	if err == sql.ErrNoRows {
		return 0
	} else if err != nil {
		Panic("tx.QueryRow failed:", err)
	}
	return id
}
