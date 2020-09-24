package memory

// RecentOpponentCount is the number of recent opponents saved for a player's people mode matching.
const RecentOpponentCount = 5

// AddRecentOpponent inserts the opponent at the beginning of the RecentOpponents list. The rest
// of the players in the list are shifted back and any past RecentOpponentCount are lost.
func (a *Player) AddRecentOpponent(id PlayerIdentifier) {
	// remove possible one duplicate of this opponent then condense the list
	for i, opp := range a.RecentOpponents {
		if opp != id {
			continue
		}
		for j := i; j < RecentOpponentCount; j++ {
			if j == (RecentOpponentCount - 1) {
				a.RecentOpponents[j] = NoPlayer
				break
			}
			a.RecentOpponents[j] = a.RecentOpponents[j+1]
		}
		break
	}

	// insert the opponent at the start of the list
	for i := (RecentOpponentCount - 1); i > 0; i-- {
		a.RecentOpponents[i] = a.RecentOpponents[i-1]
	}
	a.RecentOpponents[0] = id
}
