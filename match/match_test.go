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
	var i int
	c := make(chan struct{})
	m := NewMatcher(1, 2, pairing, func(a string, b string) {
		i++
		switch a {
		case p1:
			if b != p2 {
				t.Fatalf("%v matched with %v instead of %v", b, a, p2)
			}
		case p2:
			if b != p1 {
				t.Fatalf("%v matched with %v instead of %v", b, a, p1)
			}
		case p3:
			if b != p4 {
				t.Fatalf("%v matched with %v instead of %v", b, a, p4)
			}
		case p4:
			if b != p3 {
				t.Fatalf("%v matched with %v instead of %v", b, a, p3)
			}
		case p5:
			t.Fatalf("%v unexpectedly matched with %v", a, b)
		default:
			t.Fatalf("%v unexpectedly matched with %v", a, b)
		}
		if i == 2 {
			go func() {
				// wait in case of an unexpected third match
				time.After(time.Second)
				c <- struct{}{}
			}()
		}
	})
	// player 5 should be left unmatched
	m.Match(p1, p1r)
	m.Match(p2, p2r)
	m.Match(p3, p3r)
	m.Match(p4, p4r)
	m.Match(p5, p5r)
	select {
	case <-time.After(time.Second * 3):
		t.Fatal("matches not made after three seconds")
	case <-c:
	}
}

func TestGoodMatching(t *testing.T) {
	c := make(chan struct{})
	m := NewMatcher(1, 3, pairing, func(a string, b string) {
		if a == p1 {
			if b == p2 {
				c <- struct{}{}
				return
			}
		}
		if a == p2 {
			if b == p1 {
				c <- struct{}{}
				return
			}
		}
		t.Fatalf("unexpected match %v with %v", a, b)
	})
	m.Match(p1, p1r)
	m.Match(p2, p2r)
	select {
	case <-time.After(time.Second * 4):
		t.Fatal("no match after four seconds")
	case <-c:
	}
}

func TestBadMatching(t *testing.T) {
	c := make(chan struct{})
	m := NewMatcher(1, 2, pairing, func(a string, b string) {
		if a == p1 {
			if b == p3 {
				c <- struct{}{}
				return
			}
		}
		if a == p3 {
			if b == p1 {
				c <- struct{}{}
				return
			}
		}
		t.Fatalf("unexpected match %v with %v", a, b)
	})
	start := time.Now()
	m.Match(p1, p1r)
	m.Match(p3, p3r)
	<-c
	if time.Now().Sub(start).Seconds() < 1.5 {
		t.Fatal("bad match made early")
	}

}

func pairing(a int, b int) bool {
	if math.Abs(float64(a)-float64(b)) < 100 {
		return true
	} else {
		return false
	}
}
