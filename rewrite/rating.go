package main

// See http://www.glicko.net/research/acjpaper.pdf for a description of the Elo rating system implemented here.

import (
	"log"
	"math"
)

const (
	// The attenuation factor determines the weight this game has over all previous games.
	// For example, for k=32 pre 94.7% and this 5.3%, k=24 96.2%/3.8%, k=16 97.5%/2.5%.
	RatingAttenuation = 30
	// The value a new player should start at.
	InitialRating = 1000

	LossRatingScore float64 = 0
	DrawRatingScore         = 0.5
	WinRatingScore          = 1
)

func UpdatedRating(previous, opponent uint, score float64) uint {
	if (score < 0) || (score > 1) {
		log.Panicln("rating score", score, "out of bounds")
	}
	return previous + uint(RatingAttenuation*(score-ExpectedRatingScore(previous, opponent)))
}

func ExpectedRatingScore(rating, opponent uint) float64 {
	return 1 / (1 + math.Pow(10, -(float64(rating)-float64(opponent))/400))
}
