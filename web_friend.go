// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type FriendForm struct {
	Name        string `json:"friend"`
	Assignments []int  `json:"assignments"`
}

// post: requesting a match with friend
// get: load the game for the specified slot
func friendHandler(w http.ResponseWriter, r *http.Request) {
	key, name := database.validSession(r)
	if key == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if name == "" {
		clearClientSession(w)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	slot, err := strconv.ParseInt(r.URL.Path[8:len(r.URL.Path)], 10, 0)
	if err != nil {
		if debug {
			fmt.Println(err.Error())
		}
		http.NotFound(w, r)
		return
	}
	if (slot > 5) || (slot < 0) {
		if debug {
			fmt.Println(slot, "out of range")
		}
		http.NotFound(w, r)
		return
	}
	id := database.playersGameFromFriendSlot(name, uint8(slot))
	if r.Method == "POST" {
		if id != 0 {
			http.Redirect(w, r, r.URL.Path, http.StatusSeeOther)
			return
		}
		var form FriendForm
		err := json.NewDecoder(r.Body).Decode(&form)
		if err != nil {
			if debug {
				fmt.Println(err.Error())
			}
			http.NotFound(w, r)
			return
		}
		if form.Name == "" {
			if debug {
				fmt.Println("no name defined in friend match request")
			}
			http.NotFound(w, r)
		}
		defer r.Body.Close()
		setup, err := gameSetupFromRequest(form.Assignments)
		if err != nil {
			if debug {
				fmt.Println(err.Error())
			}
			http.NotFound(w, r)
			return
		}
		// no validation for the target friend. both players have to specify each other.
		opponentSlot := database.friendRequest(name, setup, form.Name, uint8(slot))
		if opponentSlot != -1 {
			notifyFriendMatch(form.Name, uint8(opponentSlot))
			json, err := json.Marshal(true)
			if err != nil {
				panic(err.Error())
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(json)
		}
		return
	}
	if r.Method != "GET" {
		if debug {
			fmt.Println("friend request not GET or POST")
		}
		http.NotFound(w, r)
		return
	}
	if id == 0 {
		if debug {
			fmt.Println("friend game request where no game is defined")
		}
		http.NotFound(w, r)
		return
	}
	info := database.gameInfo(id)
	if (info.White != name) && (info.Black != name) {
		if debug {
			fmt.Println(name, "not black or white")
		}
		http.NotFound(w, r)
		return
	}
	executeWebTemplate(w, game_template, gameTemplate{
		GameInfo: info,
		Name:     name,
	})
}

func friendNotificationWebsocketHandler(w http.ResponseWriter, r *http.Request) {
	key, name := database.validSession(r)
	if key == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if name == "" {
		clearClientSession(w)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if debug {
			fmt.Println(err.Error())
		}
		http.NotFound(w, r)
		return
	}
	listeningForFriendMatches(name, conn)
}

func friendCancelHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		if debug {
			fmt.Println("request not POST")
		}
		http.NotFound(w, r)
		return
	}
	key, name := database.validSession(r)
	if key == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if name == "" {
		clearClientSession(w)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	slot, err := strconv.ParseInt(r.URL.Path[14:len(r.URL.Path)], 10, 0)
	if err != nil {
		if debug {
			fmt.Println(err.Error())
		}
		http.NotFound(w, r)
		return
	}
	if (slot > 5) || (slot < 0) {
		if debug {
			fmt.Println(slot, " out of range")
		}
		http.NotFound(w, r)
		return
	}
	if database.cancelFriendRequest(name, uint8(slot)) == false {
		if debug {
			fmt.Println("cancel friend request returned false")
		}
		http.NotFound(w, r)
		return
	}
}
