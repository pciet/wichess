package wichess

import (
	"encoding/json"
	"net/http"

	"github.com/pciet/wichess/game"
	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/piece"
	"github.com/pciet/wichess/rules"
)

// The webpage sends a move request in the MoveJSON format as the POST body to /move/[game id].
// If a promotion is requested then p is nonzero and f/t is ignored.
type MoveJSON struct {
	From      rules.AddressIndex `json:"f"`
	To        rules.AddressIndex `json:"t"`
	Promotion piece.Kind         `json:"p"`
}

func movePost(w http.ResponseWriter, r *http.Request,
	gid memory.GameIdentifier, pid memory.PlayerIdentifier) {

	move, promotion := handleMovePostParse(w, r)
	if (move == rules.NoMove) && (promotion == piece.NoKind) {
		debug(MovePath, "no move and no promotion by", pid)
		return
	}

	// game is locked here instead of the auth handler to finely control the amount of lock time
	g := game.Lock(gid)
	if g.Nil() {
		debug(MovePath, "no game with ID", gid, "for", pid)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if g.PlayerActive(pid) == false {
		g.Unlock()
		debug(MovePath, "player", pid, "not active in game", gid,
			"\n", move, promotion, "\n", g.String())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	previousActive := g.PreviousActive

	var changes []rules.Square
	var captures []rules.Piece
	var promotionNeeded bool
	if promotion != piece.NoKind {
		changes = g.Promote(promotion)
	} else {
		changes, captures, promotionNeeded = g.Move(move)
	}

	if (changes == nil) || (len(changes) == 0) {
		debug(MovePath, "bad move from", pid, "in", gid, ":", move, promotion, "\n", g.String())
		g.Unlock()
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	opponentID := g.OpponentOf(pid)
	opponentOrientation := g.OrientationOf(opponentID)

	var promotionWasReverse bool
	if (promotion != piece.NoKind) && (previousActive != g.OrientationOf(pid)) {
		promotionWasReverse = true
	}

	g.Unlock()

	alertUpdate := game.Update{
		Squares:  game.SquaresFromRules(changes),
		Captures: game.PiecesFromRules(captures),
		FromMove: move,
	}
	if promotionNeeded || promotionWasReverse {
		alertUpdate.UpdateState = game.WaitUpdate
	}
	go game.Alert(gid, opponentOrientation, opponentID, alertUpdate)

	responseUpdate := game.Update{
		Squares:  game.SquaresFromRules(changes),
		Captures: game.PiecesFromRules(captures),
	}
	if promotionNeeded {
		responseUpdate.UpdateState = game.PromotionNeededUpdate
	} else if promotionWasReverse {
		responseUpdate.UpdateState = game.ContinueUpdate
	}

	jsonResponse(w, responseUpdate)
}

// handleMovePostParse parses a move or promotion from the request body. If rules.NoMove and
// piece.NoKind are returned then error handling was done and the calling handler just returns.
func handleMovePostParse(w http.ResponseWriter, r *http.Request) (rules.Move, piece.Kind) {
	body := handleLimitedBodyRead(w, r)
	if body == nil {
		return rules.NoMove, piece.NoKind
	}

	var mj MoveJSON
	err := json.Unmarshal(body, &mj)
	if err != nil {
		debug(MovePath, "failed to read body JSON:", err)
		w.WriteHeader(http.StatusBadRequest)
		return rules.NoMove, piece.NoKind
	}

	if mj == (MoveJSON{}) {
		debug(MovePath, "all zero move request")
		w.WriteHeader(http.StatusBadRequest)
		return rules.NoMove, piece.NoKind
	}

	var to rules.Address
	if mj.Promotion != piece.NoKind {
		if mj.Promotion.IsBasic() == false {
			debug("promotion request", mj.Promotion, "not basic kind")
			w.WriteHeader(http.StatusBadRequest)
			return rules.NoMove, piece.NoKind
		}
		to = rules.NoAddress
	} else {
		to = rules.AddressIndex(mj.To).Address()
	}

	return rules.Move{
		rules.AddressIndex(mj.From).Address(),
		to,
	}, mj.Promotion
}
