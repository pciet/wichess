package memory

// The easy artificial computer player "punching bag" opponent has a reserved name and id.
const (
	ComputerPlayerName       = "Computer Player"
	ComputerPlayerIdentifier = 0
)

func (a PlayerIdentifier) IsComputerPlayer() bool { return a == ComputerPlayerIdentifier }

// IncrementComputerStreak adds one to the ComputerStreak field and changes BestComputerStreak if
// it has been surpassed.
func (a *Player) IncrementComputerStreak() {
	a.ComputerStreak++
	if a.ComputerStreak > a.BestComputerStreak {
		a.BestComputerStreak = a.ComputerStreak
	}
}

// ResetComputerStreak sets the ComputerStreak field to zero.
func (a *Player) ResetComputerStreak() { a.ComputerStreak = 0 }
