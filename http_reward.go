package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pciet/wichess/piece"
)

const (
	RewardPath         = "/reward/"
	RewardHTMLTemplate = "html/reward.tmpl"
)

var RewardHandler = AuthenticRequestHandler{
	Get:  GameIdentifierParsed(PlayerNamed(RewardGet), RewardPath),
	Post: GameIdentifierParsed(PlayerNamed(RewardPost), RewardPath),
}

type RewardHTMLTemplateData struct {
	Name                string
	ID                  GameIdentifier
	Left, Right, Reward int
	Collection          [CollectionCount]piece.Kind
}

func RewardGet(w http.ResponseWriter, r *http.Request, tx *sql.Tx,
	id GameIdentifier, requester Player) {

	left, right, reward := GamePlayersRewards(tx, id,
		GamePlayersOrientation(tx, id, requester.Name))

	coll := PlayerCollection(tx, requester.ID)

	tx.Commit()

	WriteHTMLTemplate(w, RewardHTMLTemplate, RewardHTMLTemplateData{
		requester.Name, id,
		int(left), int(right), int(reward),
		coll.Kinds(),
	})
}

type RewardJSON struct {
	Left   CollectionSlot `json:"l"`
	Right  CollectionSlot `json:"r"`
	Reward CollectionSlot `json:"re"`
}

func RewardPost(w http.ResponseWriter, r *http.Request, tx *sql.Tx,
	id GameIdentifier, requester Player) {

	rj, err := ParseRewardPostBody(w, r)
	if err != nil {
		tx.Commit()
		DebugPrintln(RewardPath, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// if no rewards are put into the collection then the webpage just calls /acknowledge

	if (rj.Left < NotInCollection) || (rj.Right < NotInCollection) ||
		(rj.Reward < NotInCollection) || ((rj.Left == NotInCollection) &&
		(rj.Right == NotInCollection) && (rj.Reward == NotInCollection)) {
		tx.Commit()
		DebugPrintln(RewardPath, "bad RewardJSON", rj)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO: only query the needed pieces
	left, right, reward := GamePlayersRewards(tx, id,
		GamePlayersOrientation(tx, id, requester.Name))

	if rj.Left != NotInCollection {
		if left == piece.NoKind {
			tx.Commit()
			DebugPrintln(RewardPath, requester, "left requested but no piece in game")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		PlayerCollectionAdd(tx, requester.ID, rj.Left, left)
	}

	if rj.Right != NotInCollection {
		if right == piece.NoKind {
			tx.Commit()
			DebugPrintln(RewardPath, requester, "right requested but no piece in game")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		PlayerCollectionAdd(tx, requester.ID, rj.Right, right)
	}

	if rj.Reward != NotInCollection {
		PlayerCollectionAdd(tx, requester.ID, rj.Reward, reward)
	}

	tx.Commit()
}

func ParseRewardPostBody(w http.ResponseWriter, r *http.Request) (RewardJSON, error) {
	// TODO: this json unmarshal is repeated in http_move.go
	var body bytes.Buffer
	_, err := body.ReadFrom(http.MaxBytesReader(w, r.Body, 1024))
	if err != nil {
		return RewardJSON{}, fmt.Errorf("body read failed")
	}

	var rj RewardJSON
	err = json.Unmarshal(body.Bytes(), &rj)
	if err != nil {
		return RewardJSON{}, fmt.Errorf("failed to unmarshal json: %v", err)
	}

	return rj, nil
}

func init() { ParseHTMLTemplate(RewardHTMLTemplate) }
