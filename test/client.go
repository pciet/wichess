// Copyright 2017 Matthew Juran
// All Rights Reserved

// wichess/test is an automated client for testing the wichess server.
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/pciet/wichess/wichessing"
)

const (
	server        = "http://localhost:8080"
	ws            = "ws://localhost:8080"
	index         = server + "/"
	login         = server + "/login"
	pieces        = server + "/pieces"
	competitive15 = server + "/competitive15"
	games         = server + "/games"
	moves         = server + "/moves"
	moven         = ws + "/moven"
	makeMove      = server + "/move"
	acknowledge   = server + "/acknowledge"

	nameKey     = "name"
	passwordKey = "password"

	fromKey = "From"
	toKey   = "To"
	kindKey = "Kind"

	playerKey = "player"
	idKey     = "gameid"

	keyCookie = "k"
)

type client struct {
	name     string
	password string
}

type board [64]piece

func (b board) String() string {
	var bs string
	for rank := 7; rank >= 0; rank-- {
		for file := 0; file < 8; file++ {
			piece := b[wichessing.IndexFromFileAndRank(uint8(file), uint8(rank))]
			if piece.Kind == 0 {
				bs += fmt.Sprintf("[%v]", wichessing.AbsPoint{uint8(file), uint8(rank)})
			} else {
				bs += fmt.Sprintf("[%v %v]", wichessing.AbsPoint{uint8(file), uint8(rank)}, piece)
			}
		}
		bs += "\n"
	}
	return bs
}

type game struct {
	White  string
	Black  string
	Active string
	Points board
}

type piece struct {
	wichessing.Piece
}

func (p piece) String() string {
	return fmt.Sprintf("%v", p.Piece)
}

type encodedMoves map[string]map[string]struct{}

type movesMap map[wichessing.AbsPoint]wichessing.AbsPointSet

var clients = []client{
	{"client1", "passwordclient1"},
	{"client2", "passwordclient2"},
	{"client3", "passwordclient3"},
	{"client4", "passwordclient4"},
	{"client5", "passwordclient5"},
	{"client6", "passwordclient6"},
	{"client7", "passwordclient7"},
	{"client8", "passwordclient8"},
	{"client9", "passwordclient9"},
	{"client10", "passwordclient10"},
}

func main() {
	wait := sync.WaitGroup{}
	for _, cl := range clients {
		jar, err := cookiejar.New(nil)
		if err != nil {
			panic(err.Error())
		}
		wait.Add(1)
		go func(client *http.Client, meta client) {
			r, err := client.Get(server)
			if err != nil {
				panic(err.Error())
			}
			err = r.Body.Close()
			if err != nil {
				panic(err.Error())
			}
			if r.StatusCode != http.StatusOK {
				panic(fmt.Sprintf("unexpected error code on /: %v", r.StatusCode))
			}
			l := url.Values{}
			l.Add(nameKey, meta.name)
			l.Add(passwordKey, meta.password)
			r, err = client.PostForm(login, l)
			if err != nil {
				panic(err.Error())
			}
			err = r.Body.Close()
			if err != nil {
				panic(err.Error())
			}
			if r.StatusCode != http.StatusOK {
				panic(fmt.Sprintf("unexpected error code on /login: %v", r.StatusCode))
			}
			s, err := url.Parse(server)
			if err != nil {
				panic(err.Error())
			}
			cookies := client.Jar.Cookies(s)
			if len(cookies) != 1 {
				panic(fmt.Sprintf("%v cookies found after login\n", len(cookies)))
			}
			for _, cookie := range cookies {
				if cookie.Name != keyCookie {
					panic(fmt.Sprintf("unexpected cookie %v\n", cookie.Name))
				}
				if cookie.Value == "" {
					panic(fmt.Sprintf("%v cookie set to empty string\n", keyCookie))
				}
			}
		NEWGAME:
			for {
				// TODO: open ws://localhost:8080/competitive48n
				// get available pieces
				r, err = client.Get(pieces)
				if err != nil {
					panic(err.Error())
				}
				bodyBytes, err := ioutil.ReadAll(r.Body)
				if err != nil {
					panic(err.Error())
				}
				err = r.Body.Close()
				if err != nil {
					panic(err.Error())
				}
				// TODO: assign pieces here
				if r.StatusCode != http.StatusOK {
					panic(fmt.Sprintf("unexpected error code on GET /pieces: %v", r.StatusCode))
				}
				// TODO: test cancel
				// TODO: test competitive48, competitive5, friend, easy/hard computer, and hotseat
				// request a match
				// this post returns when a match is made
				r, err = client.Post(competitive15, "application/json", bytes.NewBuffer(assignRandomPieces(bodyBytes)))
				if err != nil {
					panic(err.Error())
				}
				err = r.Body.Close()
				if err != nil {
					panic(err.Error())
				}
				if r.StatusCode != http.StatusOK {
					panic(fmt.Sprintf("unexpected error code on POST /competitive15: %v", r.StatusCode))
				}
				r, err = client.Get(competitive15)
				if err != nil {
					panic(err.Error())
				}
				var gameid int
				scanner := bufio.NewScanner(r.Body)
				for scanner.Scan() {
					t := scanner.Text()
					if strings.Contains(t, "var game_id = ") {
						_, err = fmt.Sscanf(t, "	var game_id = %d", &gameid)
						if err != nil {
							panic(err.Error())
						}
						break
					}
				}
				err = scanner.Err()
				if err != nil {
					panic(err.Error())
				}
				if gameid == 0 {
					panic("zero value game identifier")
				}
				err = r.Body.Close()
				if err != nil {
					panic(err.Error())
				}
				if r.StatusCode != http.StatusOK {
					panic(fmt.Sprintf("unexpected error code on GET /competitive15: %v", r.StatusCode))
				}
				conn, _, err := (&websocket.Dialer{
					Jar: client.Jar,
				}).Dial(moven+"/"+fmt.Sprintf("%v", gameid), nil)
				if err != nil {
					panic(err.Error())
				}
				r, err = client.Get(games + "/" + fmt.Sprintf("%d", gameid))
				if err != nil {
					panic(err.Error())
				}
				var g game
				err = json.NewDecoder(r.Body).Decode(&g)
				if err != nil {
					panic(err.Error())
				}
				err = r.Body.Close()
				if err != nil {
					panic(err.Error())
				}
				if r.StatusCode != http.StatusOK {
					panic(fmt.Sprintf("unexpected error code on GET /games/%v: %v", gameid, r.StatusCode))
				}
				if (meta.name != g.White) && (meta.name != g.Black) {
					panic(fmt.Sprintf("%v is not White (%v) or Black (%v)", meta.name, g.White, g.Black))
				}
				var orientation wichessing.Orientation
				if meta.name == g.White {
					orientation = wichessing.White
				} else {
					orientation = wichessing.Black
				}
			PLAYGAME:
				for {
					r, err = client.Get(moves + "/" + fmt.Sprintf("%d", gameid))
					if err != nil {
						panic(err.Error())
					}
					if r.StatusCode != http.StatusOK {
						panic(fmt.Sprintf("unexpected error code on GET /moves/%v: %v", gameid, r.StatusCode))
					}
					em := encodedMoves{}
					err = json.NewDecoder(r.Body).Decode(&em)
					if err != nil {
						panic(err.Error())
					}
					err = r.Body.Close()
					if err != nil {
						panic(err.Error())
					}
					availableMoves := movesMap{}
					for point, mvs := range em {
						if (point == "checkmate") || (point == "draw") || (point == "time") {
							l = url.Values{}
							l.Add(playerKey, meta.name)
							l.Add(idKey, fmt.Sprintf("%d", gameid))
							r, err = client.PostForm(acknowledge, l)
							if err != nil {
								panic(err.Error())
							}
							if r.StatusCode != http.StatusOK {
								panic(fmt.Sprintf("unexpected error code on POST /acknowledge: %v", r.StatusCode))
							}
							err = r.Body.Close()
							if err != nil {
								panic(err.Error())
							}
							err = conn.Close()
							if err != nil {
								panic(err.Error())
							}
							continue NEWGAME
						}
						if point == "promote" {
							if g.Active == meta.name {
								if orientation == wichessing.White {
									g.Active = g.Black
								} else {
									g.Active = g.White
								}
								break // wait for opponent to promote
							}
							from := 0
							if orientation == wichessing.White {
								for i := 56; i < 64; i++ {
									if (g.Points[i].Base == wichessing.Pawn) && (g.Points[i].Orientation == wichessing.White) {
										from = i
										break
									}
								}
							} else {
								for i := 0; i < 8; i++ {
									if (g.Points[i].Base == wichessing.Pawn) && (g.Points[i].Orientation == wichessing.Black) {
										from = i
										break
									}
								}
							}
							l = url.Values{}
							l.Add(fromKey, fmt.Sprintf("%d", from))
							l.Add(kindKey, fmt.Sprintf("%d", wichessing.Queen))
							r, err = client.PostForm(makeMove+"/"+fmt.Sprintf("%v", gameid), l)
							if err != nil {
								panic(err.Error())
							}
							if r.StatusCode != http.StatusOK {
								panic(fmt.Sprintf("unexpected error code on POST /move/%v: %v", gameid, r.StatusCode))
							}
							diff := map[string]piece{}
							err = json.NewDecoder(r.Body).Decode(&diff)
							if err != nil {
								panic(err.Error())
							}
							err = r.Body.Close()
							if err != nil {
								panic(err.Error())
							}
							for point, p := range diff {
								g.Points[wichessing.IndexFromAddressString(point)] = p
							}
							if orientation == wichessing.White {
								g.Active = g.Black
							} else {
								g.Active = g.White
							}
							continue PLAYGAME
						}
						if point == "check" {
							continue
						}
						set := wichessing.AbsPointSet{}
						for mv, _ := range mvs {
							pt := wichessing.AbsPointFromAddressString(mv)
							set[&pt] = struct{}{}
						}
						availableMoves[wichessing.AbsPointFromAddressString(point)] = set
					}
					if g.Active != meta.name {
						diff := map[string]piece{}
						err = conn.ReadJSON(&diff)
						if err != nil {
							panic(err.Error())
						}
						for point, p := range diff {
							g.Points[wichessing.IndexFromAddressString(point)] = p
						}
						g.Active = meta.name
						continue
					}
				MAKEMOVE:
					for pt, mvs := range availableMoves {
						if g.Points[pt.Index()].Orientation != orientation {
							continue
						}
						for mv, _ := range mvs {
							l = url.Values{}
							l.Add(fromKey, fmt.Sprintf("%d", pt.Index()))
							l.Add(toKey, fmt.Sprintf("%d", mv.Index()))
							r, err = client.PostForm(makeMove+"/"+fmt.Sprintf("%v", gameid), l)
							if err != nil {
								panic(err.Error())
							}
							if r.StatusCode != http.StatusOK {
								fmt.Printf("%v (%v) moving %v to %v failed\n", meta.name, orientation, pt, mv)
								fmt.Println(availableMoves)
								fmt.Println(g.Points)
								panic(fmt.Sprintf("unexpected error code on POST /move/%v: %v", gameid, r.StatusCode))
							}
							diff := map[string]piece{}
							err = json.NewDecoder(r.Body).Decode(&diff)
							if err != nil {
								panic(err.Error())
							}
							err = r.Body.Close()
							if err != nil {
								panic(err.Error())
							}
							for point, p := range diff {
								g.Points[wichessing.IndexFromAddressString(point)] = p
							}
							if orientation == wichessing.White {
								g.Active = g.Black
							} else {
								g.Active = g.White
							}
							break MAKEMOVE
						}
					}
				}
			}
			wait.Done()
		}(&http.Client{
			Jar: jar,
		}, cl)
	}
	wait.Wait()
}

// returns a JSON string with the piece assignments
func assignRandomPieces(piecesBody []byte) []byte {
	var pieces = map[string][]int{
		"assignments": []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	jsonm, err := json.Marshal(pieces)
	if err != nil {
		panic(err.Error())
	}
	return jsonm
}
