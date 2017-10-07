// Copyright 2017 Matthew Juran
// All Rights Reserved

package match

import (
	"math"
	"testing"
	"time"
)

const (
	p1 = "player1"
	p2 = "player2"
	p3 = "player3"
	p4 = "player4"
	p5 = "player5"

	p1r = 100
	p2r = 110
	p3r = 250
	p4r = 275
	p5r = 500
)

func TestMultipleGoodMatching(t *testing.T) {
	m := NewMatcher(1, 2, pairing)
	// player 5 should be left unmatched
	c1 := m.Match(p1, p1r)
	c2 := m.Match(p2, p2r)
	c3 := m.Match(p3, p3r)
	c4 := m.Match(p4, p4r)
	c5 := m.Match(p5, p5r)
	count := make(chan struct{})
	go listen(t, true, p2, c1, count)
	go listen(t, true, p1, c2, count)
	go listen(t, true, p4, c3, count)
	go listen(t, true, p3, c4, count)
	go listen(t, true, p5, c5, count)
	var i int
OUTER:
	for {
		select {
		case <-count:
			i++
			continue
		case <-time.After(time.Second * 2):
			break OUTER
		}
	}
	if i != 4 {
		t.Fatalf("%v matchings indicated instead of 4", i)
	}
}

func listen(t *testing.T, matched bool, opponent string, match <-chan string, done chan<- struct{}) {
	opp := <-match
	if matched {
		if opponent != opp {
			t.Fatalf("recieved %v instead of opponent %v", opp, opponent)
			return
		}
	} else {
		t.Fatalf("recieved %v instead of no match", opp)
	}
	done <- struct{}{}
}

func TestGoodMatching(t *testing.T) {
	m := NewMatcher(1, 3, pairing)
	c1 := m.Match(p1, p1r)
	c2 := m.Match(p2, p2r)
	var m1, m2 string
	select {
	case m1 = <-c1:
		if m1 != p2 {
			t.Fatalf("p1 recieved match with %v, not %v", m1, p2)
		}
	case m2 = <-c2:
		if m2 != p1 {
			t.Fatalf("p2 recieved match with %v, not %v", m2, p1)
		}
	case <-time.After(time.Second * 2):
		t.Fatal("no match made within two seconds (bad match threshold is three seconds but the match is a good one)")
	}
	if m2 == "" {
		select {
		case m2 = <-c2:
			if m2 != p1 {
				t.Fatalf("p2 recieved match with %v, not %v", m2, p1)
			}
		case <-time.After(time.Second * 2):
			t.Fatal("no match made within two seconds")
		}
	} else {
		select {
		case m1 = <-c1:
			if m1 != p2 {
				t.Fatalf("p1 recieved match with %v, not %v", m1, p2)
			}
		case <-time.After(time.Second * 2):
			t.Fatal("no match made within two seconds")
		}
	}
}

func TestBadMatching(t *testing.T) {
	m := NewMatcher(1, 2, pairing)
	start := time.Now()
	c1 := m.Match(p1, p1r)
	c3 := m.Match(p3, p3r)
	var m1, m3 string
	select {
	case m1 = <-c1:
		if m1 != p3 {
			t.Fatalf("p1 recieved match with %v, not %v", m1, p3)
		}
	case m3 = <-c3:
		if m3 != p1 {
			t.Fatalf("p3 recieved match with %v, not %v", m3, p1)
		}
	case <-time.After(time.Second * 3):
		t.Fatal("no match made within three seconds (bad match threshold is three seconds)")
	}
	if time.Now().Sub(start).Seconds() < 1.5 {
		t.Fatal("bad match made early")
	}
	if m3 == "" {
		select {
		case m3 = <-c3:
			if m3 != p1 {
				t.Fatalf("p3 recieved match with %v, not %v", m3, p1)
			}
		case <-time.After(time.Second * 2):
			t.Fatal("no match made within two seconds")
		}
	} else {
		select {
		case m1 = <-c1:
			if m1 != p3 {
				t.Fatalf("p1 recieved match with %v, not %v", m1, p3)
			}
		case <-time.After(time.Second * 2):
			t.Fatal("no match made within two seconds")
		}
	}
}

func pairing(a int, b int) bool {
	if math.Abs(float64(a)-float64(b)) < 100 {
		return true
	} else {
		return false
	}
}
