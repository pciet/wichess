// Copyright 2017 Matthew Juran
// All Rights Reserved

package match

import (
	"sync"
	"time"
)

type Matcher struct {
	// The rate at which a matching is triggered in seconds.
	period uint
	// The number of bad pairings without good pairings after which a bad match will be made anyway.
	threshold uint
	// The function used to determine if a pairing is good or bad.
	pairing func(int, int) bool
	// The list of available players and their response channels.
	list map[string]player
	// Indicates if this matcher is attempting to match every period.
	matching *bool
	*sync.Mutex
}

// Period is how many seconds between matching attempt, threshold is how many bad matches before the bad match is made anyway, and pairing is the function that determines if a match is good or bad.
func NewMatcher(period uint, threshold uint, pairing func(int, int) bool) Matcher {
	f := false
	return Matcher{period: period,
		threshold: threshold,
		pairing:   pairing,
		list:      make(map[string]player),
		matching:  &f,
		Mutex:     &sync.Mutex{},
	}
}

type player struct {
	rating int
	match  chan<- string
	bad    map[string]uint
}

// Returns nil if this player is already being matched.
func (m Matcher) Match(name string, rating int) <-chan string {
	m.Lock()
	defer m.Unlock()
	_, has := m.list[name]
	if has {
		return nil
	}
	c := make(chan string)
	m.list[name] = player{rating: rating, match: c, bad: make(map[string]uint)}
	if (*m.matching == false) && (len(m.list) > 1) {
		*m.matching = true
		go m.matchmaking()
	}
	return c
}

// This is the only place where players are removed from the matcher list.
func (m Matcher) matchmaking() {
	for {
		<-time.After(time.Second * time.Duration(m.period))
		m.Lock()
	OUTER:
		for name, p1 := range m.list {
			for opponent, p2 := range m.list {
				if name == opponent {
					continue
				}
				if m.pairing(p1.rating, p2.rating) == false {
					p1.bad[opponent]++
				} else {
					p1.match <- opponent
					close(p1.match)
					p2.match <- name
					close(p2.match)
					delete(m.list, name)
					delete(m.list, opponent)
					continue OUTER
				}
			}
		}
		if len(m.list) < 2 {
			*m.matching = false
			m.Unlock()
			return
		}
		// if any bad pairings are left that meet the threshold then make them anyway
	OUTER2:
		for name, p1 := range m.list {
			for opponent, p2 := range m.list {
				if name == opponent {
					continue
				}
				if p1.bad[opponent] >= m.threshold {
					p1.match <- opponent
					close(p1.match)
					p2.match <- name
					close(p2.match)
					delete(m.list, name)
					delete(m.list, opponent)
					continue OUTER2
				}
			}
		}
		if len(m.list) < 2 {
			*m.matching = false
			m.Unlock()
			return
		}
		m.Unlock()
	}
}
