// Copyright 2017 Matthew Juran
// All Rights Reserved

// See http://www.glicko.net/research/acjpaper.pdf for theory describing the foundation of the rating system implemented here.
package rating

import (
	"fmt"
	"math"
)

const (
	// The attenuation factor determines the weight this game has over all previous games.
	// For example, for k=32 pre 94.7% and this 5.3%, k=24 96.2%/3.8%, k=16 97.5%/2.5%.
	k = 30
	// The value a new player should start at.
	Initial = 1000

	Loss float64 = 0
	Draw         = 0.5
	Win          = 1
)

func Updated(previous, opponent uint, score float64) uint {
	if (score < 0) || (score > 1) {
		panic(fmt.Sprintf("rating: score %v out of bounds", score))
	}
	return previous + uint(k*(score-ExpectedScore(previous, opponent)))
}

func ExpectedScore(rating, opponent uint) float64 {
	return 1 / (1 + math.Pow(10, -(float64(rating)-float64(opponent))/400))
}
