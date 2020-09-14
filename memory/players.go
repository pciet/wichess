package memory

import (
	"sync"
)

var (
	PlayersCache map[PlayerIdentifier]*Player
	PlayersMutex sync.RWMutex
)
