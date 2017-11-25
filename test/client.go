// Copyright 2017 Matthew Juran
// All Rights Reserved

// wichess/test is an automated client for testing the wichess server.
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pciet/wichess/wichessing"
)

const (
	debug         = true
	client_count  = 40
	delay_seconds = 0
	// [0,1), if random number in that range is less than this value then no special pieces will be used for this round
	special_piece_probability = 0.5
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
	Identifier int
	wichessing.Piece
}

func (p piece) String() string {
	return fmt.Sprintf("%v", p.Piece)
}

type encodedMoves map[string]map[string]struct{}

type movesMap map[wichessing.AbsPoint]wichessing.AbsPointSet

func (g game) promoting() (bool, wichessing.Orientation) {
	for i := 56; i < 64; i++ {
		if (g.Points[i].Base == wichessing.Pawn) && (g.Points[i].Orientation == wichessing.White) {
			return true, wichessing.White
		}
	}
	for i := 0; i < 8; i++ {
		if (g.Points[i].Base == wichessing.Pawn) && (g.Points[i].Orientation == wichessing.Black) {
			return true, wichessing.Black
		}
	}
	return false, wichessing.White
}

func main() {
	rand.Seed(42)
	clients := make([]client, client_count)
	for i := 0; i < client_count; i++ {
		clients[i] = client{
			name:     fmt.Sprintf("client%v", i),
			password: fmt.Sprintf("passwordclient%v", i),
		}
	}
	wait := sync.WaitGroup{}
	for _, cl := range clients {
		jar, err := cookiejar.New(nil)
		if err != nil {
			panic(err.Error())
		}
		// let the client hold onto more connections
		// http://tleyden.github.io/blog/2016/11/21/tuning-the-go-http-client-library-for-load-testing/
		dtp, ok := http.DefaultTransport.(*http.Transport)
		if ok == false {
			panic("failed to assert http.DefaultTransport to *http.Transport")
		}
		dt := *dtp
		dt.MaxIdleConns = client_count
		dt.MaxIdleConnsPerHost = client_count
		// dt set as client Transport on the go func call below
		wait.Add(1)
		go func(client *http.Client, meta client) {
			dialer := &websocket.Dialer{
				Jar: client.Jar,
			}
			// set an initial time offset
			<-time.After(time.Second * time.Duration(20*rand.Float64()))
			if debug {
				fmt.Printf("(%v) GET %v\n", meta.name, server)
			}
			r, err := client.Get(server)
			if err != nil {
				panic(err.Error())
			}
			_, err = io.Copy(ioutil.Discard, r.Body)
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
			<-time.After(time.Second * delay_seconds)
			l := url.Values{}
			l.Add(nameKey, meta.name)
			l.Add(passwordKey, meta.password)
			if debug {
				fmt.Printf("(%v) POST %v: %v %v\n", meta.name, login, meta.name, meta.password)
			}
			r, err = client.PostForm(login, l)
			if err != nil {
				panic(err.Error())
			}
			_, err = io.Copy(ioutil.Discard, r.Body)
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
				<-time.After(time.Second * delay_seconds)
				// TODO: open ws://localhost:8080/competitive48n
				if debug {
					fmt.Printf("(%v) GET %v\n", meta.name, pieces)
				}
				r, err = client.Get(pieces)
				if err != nil {
					panic(err.Error())
				}
				var pcs []piece
				err = json.NewDecoder(r.Body).Decode(&pcs)
				if err != nil {
					panic(err.Error())
				}
				err = r.Body.Close()
				if err != nil {
					panic(err.Error())
				}
				if r.StatusCode != http.StatusOK {
					panic(fmt.Sprintf("unexpected error code on GET /pieces: %v", r.StatusCode))
				}
				assignments := assignRandomPieces(pcs)
				<-time.After(time.Second * delay_seconds)
				// TODO: test cancel
				// TODO: test competitive48, competitive5, friend, easy/hard computer, and hotseat
				// request a match
				// this post returns when a match is made
				if debug {
					fmt.Printf("(%v) POST %v: %v\n", meta.name, competitive15, string(assignments))
				}
				r, err = client.Post(competitive15, "application/json", bytes.NewBuffer(assignments))
				if err != nil {
					panic(err.Error())
				}
				_, err = io.Copy(ioutil.Discard, r.Body)
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
				if debug {
					fmt.Printf("(%v) GET %v\n", meta.name, competitive15)
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
				_, err = io.Copy(ioutil.Discard, r.Body)
				if err != nil {
					panic(err.Error())
				}
				err = r.Body.Close()
				if err != nil {
					panic(err.Error())
				}
				if r.StatusCode != http.StatusOK {
					panic(fmt.Sprintf("unexpected error code on GET /competitive15: %v", r.StatusCode))
				}
				if debug {
					fmt.Printf("(%v) DIAL %v\n", meta.name, moven+"/"+fmt.Sprintf("%v", gameid))
				}
				conn, r, err := dialer.Dial(moven+"/"+fmt.Sprintf("%v", gameid), nil)
				if err != nil {
					fmt.Println(r)
					panic(fmt.Sprintf("%v: error code %v\n", err.Error(), r.StatusCode))
				}
				if debug {
					fmt.Printf("(%v) GET %v\n", meta.name, games+"/"+fmt.Sprintf("%d", gameid))
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
				} else if meta.name == g.Black {
					orientation = wichessing.Black
				} else {
					panic(fmt.Sprintf("game expects white (%v) or black (%v) but player is %v\n", g.White, g.Black, meta.name))
				}
				timeLoss := false
			PLAYGAME:
				for {
					if debug {
						fmt.Printf("(%v) GET %v\n", meta.name, moves+"/"+fmt.Sprintf("%d", gameid))
					}
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
					if len(em) == 0 {
						panic(fmt.Sprintf("(%v) recieved zero length response to GET %v\n", meta.name, moves+"/"+fmt.Sprintf("%d", gameid)))
					}
					availableMoves := movesMap{}
					for point, _ := range em {
						if (point == "checkmate") || (point == "draw") || (point == "time") {
							<-time.After(time.Second * delay_seconds)
							l = url.Values{}
							l.Add(playerKey, meta.name)
							l.Add(idKey, fmt.Sprintf("%d", gameid))
							if debug {
								fmt.Printf("(%v) POST %v: %v %v\n", meta.name, acknowledge, meta.name, gameid)
							}
							r, err = client.PostForm(acknowledge, l)
							if err != nil {
								panic(err.Error())
							}
							if r.StatusCode != http.StatusOK {
								panic(fmt.Sprintf("unexpected error code on POST /acknowledge: %v", r.StatusCode))
							}
							_, err = io.Copy(ioutil.Discard, r.Body)
							if err != nil {
								panic(err.Error())
							}
							err = r.Body.Close()
							if err != nil {
								panic(err.Error())
							}
							if debug {
								fmt.Printf("(%v) CLOSE %v\n", meta.name, moven+"/"+fmt.Sprintf("%v", gameid))
							}
							err = conn.Close()
							if err != nil {
								panic(err.Error())
							}
							continue NEWGAME
						}
					}
					for point, mvs := range em {
						if point == "promote" {
							if g.Active != meta.name {
								break // wait for opponent to promote
							}
							from := 0
							if orientation == wichessing.White {
							WHITEPROMOTE:
								for {
									for i := 56; i < 64; i++ {
										if (g.Points[i].Base == wichessing.Pawn) && (g.Points[i].Orientation == wichessing.White) {
											from = i
											break WHITEPROMOTE
										}
									}
									fmt.Println(g.Points)
									panic(fmt.Sprintf("(%v) found no promotion (white)\n", meta.name))
								}
							} else {
							BLACKPROMOTE:
								for {
									for i := 0; i < 8; i++ {
										if (g.Points[i].Base == wichessing.Pawn) && (g.Points[i].Orientation == wichessing.Black) {
											from = i
											break BLACKPROMOTE
										}
									}
									fmt.Println(g.Points)
									panic(fmt.Sprintf("(%v) found no promotion (black)\n", meta.name))
								}
							}
							<-time.After(time.Second * delay_seconds)
							l = url.Values{}
							l.Add(fromKey, fmt.Sprintf("%d", from))
							l.Add(kindKey, fmt.Sprintf("%d", wichessing.Queen))
							if debug {
								fmt.Printf("(%v) POST %v %v %v\n", meta.name, wichessing.AbsPointFromIndex(uint8(from)), "Queen", makeMove+"/"+fmt.Sprintf("%v", gameid))
							}
							r, err = client.PostForm(makeMove+"/"+fmt.Sprintf("%v", gameid), l)
							if err != nil {
								panic(err.Error())
							}
							if r.StatusCode != http.StatusOK {
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
					if timeLoss == true {
						fmt.Println(em)
						fmt.Println(g.Points)
						panic(fmt.Sprintf("%v recieved zero length diff but the game is not a time loss", meta.name))
					}
					if g.Active != meta.name {
						diff := map[string]piece{}
						if debug {
							fmt.Printf("(%v) WAIT DIFF\n", meta.name)
						}
						err = conn.ReadJSON(&diff)
						if err != nil {
							panic(err.Error())
						}
						if debug {
							fmt.Printf("(%v) RECV DIFF\n", meta.name)
						}
						if len(diff) == 0 {
							timeLoss = true
							continue
						}
						for point, p := range diff {
							g.Points[wichessing.IndexFromAddressString(point)] = p
						}
						promoting, promotingOrientation := g.promoting()
						if (promoting == false) || (promoting && (promotingOrientation == orientation)) {
							g.Active = meta.name
						}
						<-time.After(time.Second * delay_seconds)
						continue
					}
					for pt, mvs := range availableMoves {
						if g.Points[pt.Index()].Orientation != orientation {
							continue
						}
						for mv, _ := range mvs {
							<-time.After(time.Second * delay_seconds)
							l = url.Values{}
							l.Add(fromKey, fmt.Sprintf("%d", pt.Index()))
							l.Add(toKey, fmt.Sprintf("%d", mv.Index()))
							if debug {
								fmt.Printf("(%v) POST %v %v %v\n", meta.name, pt, mv, makeMove+"/"+fmt.Sprintf("%v", gameid))
							}
							r, err = client.PostForm(makeMove+"/"+fmt.Sprintf("%v", gameid), l)
							if err != nil {
								fmt.Println(g.Points)
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
							if len(diff) == 0 {
								timeLoss = true
								continue PLAYGAME
							}
							for point, p := range diff {
								g.Points[wichessing.IndexFromAddressString(point)] = p
							}
							promoting, promotingOrientation := g.promoting()
							if (promoting == false) || (promoting && (promotingOrientation != orientation)) {
								if orientation == wichessing.White {
									g.Active = g.Black
								} else {
									g.Active = g.White
								}
							}
							continue PLAYGAME
						}
						fmt.Println(availableMoves)
						panic(fmt.Sprintf("(%v) has no move at %v\n", meta.name, pt))
					}
					fmt.Println(availableMoves)
					fmt.Println(g.Points)
					panic(fmt.Sprintf("(%v) has no available moves\n", meta.name))
				}
			}
			wait.Done()
		}(&http.Client{
			Jar:       jar,
			Transport: &dt,
		}, cl)
	}
	wait.Wait()
}

func pieceSliceHas(slice []piece, p piece) bool {
	for _, pc := range slice {
		if pc == p {
			return true
		}
	}
	return false
}

// returns a JSON string with the piece assignments
func assignRandomPieces(pieces []piece) []byte {
	used := make([]piece, 0, 16)
	assignments := [16]int{}
	for i := 0; i < 16; i++ {
		if (i == 8) || (i == 15) {
			for _, pc := range pieces {
				if pieceSliceHas(used, pc) {
					continue
				}
				if pc.Base == wichessing.Rook {
					assignments[i] = pc.Identifier
					used = append(used, pc)
					break
				}
			}
		} else if (i == 9) || (i == 14) {
			for _, pc := range pieces {
				if pieceSliceHas(used, pc) {
					continue
				}
				if pc.Base == wichessing.Knight {
					assignments[i] = pc.Identifier
					used = append(used, pc)
					break
				}
			}
		} else if (i == 10) || (i == 13) {
			for _, pc := range pieces {
				if pieceSliceHas(used, pc) {
					continue
				}
				if pc.Base == wichessing.Bishop {
					assignments[i] = pc.Identifier
					used = append(used, pc)
					break
				}
			}
			// no kings or queens for now
		} else if i < 8 {
			for _, pc := range pieces {
				if pieceSliceHas(used, pc) {
					continue
				}
				if pc.Base == wichessing.Pawn {
					assignments[i] = pc.Identifier
					used = append(used, pc)
					break
				}
			}
		}
	}
	if rand.Float32() < special_piece_probability {
		assignments = [16]int{}
	}
	jsonm, err := json.Marshal(map[string][16]int{
		"assignments": assignments,
	})
	if err != nil {
		panic(err.Error())
	}
	return jsonm
}
