package wichess

import (
	"encoding/json"
	"net/http"

	"github.com/pciet/wichess/game"
	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/piece"
)

type RewardHTMLTemplateData struct {
	memory.PlayerName
	memory.GameIdentifier     // needed to acknowledge game review complete
	Left, Right, Reward   int // piece.Kind templates string, so use int
	piece.Collection
}

func rewardGet(w http.ResponseWriter, r *http.Request, g game.Instance, p *memory.Player) {
	left, right, reward := g.RewardsOf(p.PlayerIdentifier)
	writeHTMLTemplate(w, RewardHTMLTemplate, RewardHTMLTemplateData{
		PlayerName:     p.PlayerName,
		GameIdentifier: g.GameIdentifier,
		Left:           int(left),
		Right:          int(right),
		Reward:         int(reward),
		Collection:     p.Collection,
	})
}

type RewardJSON struct {
	Left   piece.CollectionSlot `json:"l"`
	Right  piece.CollectionSlot `json:"r"`
	Reward piece.CollectionSlot `json:"re"`
}

func rewardPost(w http.ResponseWriter, r *http.Request, g game.Instance, p *memory.Player) {
	rj := handleRewardPostParse(w, r)
	if rj == (RewardJSON{}) {
		return
	}

	left, right, reward := g.RewardsOf(p.PlayerIdentifier)

	write := false

	if (rj.Left > 0) && (rj.Left <= piece.CollectionSize) {
		if left == piece.NoKind {
			debug(RewardPath, "left requested but no piece in game for", p.Name)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		p.Collection[rj.Left-1] = left
		p.Left = p.Left.DifferentSpecialKind()
		write = true
	}

	if (rj.Right > 0) && (rj.Right <= piece.CollectionSize) {
		if right == piece.NoKind {
			debug(RewardPath, "right requested but no piece in game for", p.Name)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		p.Collection[rj.Right-1] = right
		p.Right = p.Right.DifferentSpecialKind()
		write = true
	}

	if (rj.Reward > 0) && (rj.Reward <= piece.CollectionSize) {
		p.Collection[rj.Reward-1] = reward
		write = true
	}

	if write {
		go memory.WritePlayerFile(p.PlayerIdentifier)
	}
}

func handleRewardPostParse(w http.ResponseWriter, r *http.Request) RewardJSON {
	body := handleLimitedBodyRead(w, r)
	if body == nil {
		return RewardJSON{}
	}
	var rj RewardJSON
	err := json.Unmarshal(body, &rj)
	if err != nil {
		debug(RewardPath, "failed to unmarshal JSON:", err)
		w.WriteHeader(http.StatusBadRequest)
		return RewardJSON{}
	}
	if (rj.Left < piece.NotInCollection) || (rj.Left > piece.CollectionSize) ||
		(rj.Right < piece.NotInCollection) || (rj.Right > piece.CollectionSize) ||
		(rj.Reward < piece.NotInCollection) || (rj.Reward > piece.CollectionSize) ||
		((rj.Left == piece.NotInCollection) && (rj.Right == piece.NotInCollection) &&
			(rj.Reward == piece.NotInCollection)) {

		// the case of all being NotInCollection is invalid because the webpage doesn't post
		// to /reward if no reward is requested
		debug(RewardPath, "bad RewardJSON", rj)
		w.WriteHeader(http.StatusBadRequest)
		return RewardJSON{}
	}
	return rj
}
