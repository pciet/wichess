package wichess

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/piece"
)

type RewardHTMLTemplateData struct {
	memory.PlayerName
	memory.GameIdentifier // needed to acknowledge game review complete
	Left, Right, Reward   piece.Kind
	piece.Collection
}

func rewardGet(w http.ResponseWriter, r *http.Request, g game.Instance, p player.Instance) {
	left, right, reward := g.RewardsOf(p.PlayerIdentifier)
	writeHTMLTemplate(w, RewardHTMLTemplate, RewardHTMLTemplateData{
		PlayerName:     p.Name,
		GameIdentifier: g.GameIdentifier,
		Left:           left,
		Right:          right,
		Reward:         reward,
		Collection:     p.Collection,
	})
}

type RewardJSON struct {
	Left   piece.CollectionSlot `json:"l"`
	Right  piece.CollectionSlot `json:"r"`
	Reward piece.CollectionSlot `json:"re"`
}

func rewardPost(w http.ResponseWriter, r *http.Request, g game.Instance, p player.Instance) {
	rj := handleRewardPostParse(w, r)
	if rj == (RewardJSON{}) {
		return
	}

	left, right, reward := g.RewardsOf(p.PlayerIdentifier)

	if rj.Left != piece.NotInCollection {
		if left == piece.NoKind {
			debug(RewardPath, "left requested but no piece in game for", p.Name)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		p.CollectionAdd(rj.Left, left)
	}

	if rj.Right != piece.NotInCollection {
		if right == piece.NoKind {
			debug(RewardPath, "right requested but no piece in game for", p.Name)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		p.CollectionAdd(rj.Right, right)
	}

	if rj.Reward != piece.NotInCollection {
		p.CollectionAdd(rj.Reward, reward)
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
