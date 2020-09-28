package game

import (
	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/piece"
)

func selectedCollectionPieces(id memory.PlayerIdentifier,
	requests []piece.CollectionSlot) ([]piece.Kind, piece.Kind, piece.Kind) {

	p := memory.RLockPlayer(id)
	if p == nil {
		return nil, piece.NoKind, piece.NoKind
	}
	defer p.RUnlock()

	out := make([]piece.Kind, 0, 8)
	for _, r := range requests {
		if (r <= 0) || (r > piece.CollectionSize) {
			continue
		}
		out = append(out, p.Collection[r-1])
	}

	return out, p.Left, p.Right
}
