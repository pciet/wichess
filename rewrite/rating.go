package main

import "math"

// Adapted from the Elo rating system as described in http://www.glicko.net/research/acjpaper.pdf

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
		Panic("rating score", score, "out of bounds")
	}
	return previous + uint(RatingAttenuation*(score-ExpectedRatingScore(previous, opponent)))
}

func ExpectedRatingScore(rating, opponent uint) float64 {
	return 1 / (1 + math.Pow(10, -(float64(rating)-float64(opponent))/400))
}
