package chess

import (
	"github.com/Jacobbrewer1/chess-boards/src/custom"
	"math"
)

func probability(rating1, rating2 float64) float64 {
	return float64(1.0) * 1.0 / (1 + 1.0*math.Pow(10, 1.0*(rating1-rating2)/400))
}

// CalculateNew returns the new player a rating and the new rating for player b
func CalculateNew(ratingA, ratingB float64, constant int, playerAWon custom.Bool) (float64, float64) {
	var newRatingA, newRatingB float64

	probabilityPlayerB := probability(ratingA, ratingB)
	probabilityPlayerA := probability(ratingB, ratingA)

	if playerAWon {
		newRatingA = ratingA + float64(constant)*(1-probabilityPlayerA)
		newRatingB = ratingB + float64(constant)*(0-probabilityPlayerB)
	} else {
		newRatingA = ratingA + float64(constant)*(0-probabilityPlayerA)
		newRatingB = ratingB + float64(constant)*(1-probabilityPlayerB)
	}

	return math.Round(newRatingA*1.0) / 1.0, math.Round(newRatingB*1.0) / 1.0
}
