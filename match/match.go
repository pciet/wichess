// Copyright 2017 Matthew Juran
// All Rights Reserved

package match

import (
	"fmt"
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
	// Called when a match is made.
	match func(string, interface{}, string, interface{})
	// The list of available players and their response channels.
	list map[string]player
	// Indicates if this matcher is attempting to match every period.
	matching *bool
	*sync.RWMutex
}

// Period is how many seconds between matching attempt, threshold is how many bad matches before the bad match is made anyway, and pairing is the function that determines if a match is good or bad. Match is a callback when a match is made.
func NewMatcher(period uint, threshold uint, pairing func(int, int) bool, match func(string, interface{}, string, interface{})) Matcher {
	f := false
	return Matcher{period: period,
		threshold: threshold,
		pairing:   pairing,
		list:      make(map[string]player),
		match:     match,
		matching:  &f,
		RWMutex:   &sync.RWMutex{},
	}
}

type player struct {
	rating int
	meta   interface{}
	bad    map[string]uint
}

// This function will panic if a player is already being matched.
func (m Matcher) Match(name string, meta interface{}, rating int) {
	m.Lock()
	defer m.Unlock()
	_, has := m.list[name]
	if has {
		panic(fmt.Sprintf("match: player %v has Match() called multiple times", name))
	}
	m.list[name] = player{rating: rating, meta: meta, bad: make(map[string]uint)}
	if (*m.matching == false) && (len(m.list) > 1) {
		*m.matching = true
		go m.matchmaking()
	}
}

// If this player is matching the meta data provided in Match() will be returned here, nil otherwise.
func (m Matcher) Matching(name string) interface{} {
	m.RLock()
	defer m.RUnlock()
	player, has := m.list[name]
	if has {
		return player.meta
	} else {
		return nil
	}
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
					m.match(name, p1.meta, opponent, p2.meta)
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
					m.match(name, p1.meta, opponent, p2.meta)
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
