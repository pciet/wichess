package piece

// CollectionSize is the number of pieces a player's collection can have.
const CollectionSize = 21

// These three special CollectionSlot values indicate either that a piece isn't in the collection
// or is one of the random picks.
const (
	NotInCollection = 0
	LeftPick        = -1
	RightPick       = -2
)

type (
	// The CollectionSlot is a location in a Collection starting at one going up to CollectionSize.
	// Special values are NotInCollection, LeftPick, and RightPick.
	CollectionSlot int

	// A Collection is a set of special pieces a player owns that are used to construct armies.
	Collection [CollectionSize]Kind

	// A player has two picks that are the source of new special pieces. These are random and
	// replaced when used in an army. If used then they can be optionally added to the collection
	// when a game is completed without being conceded.
	RandomPicks struct {
		Left, Right Kind
	}
)

func (a CollectionSlot) Int() int { return int(a) }
