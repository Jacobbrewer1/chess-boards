package chess

import (
	"github.com/Jacobbrewer1/chess-boards/src/custom"
	"math"
)

const EloInitialValue = 1500
const EloDefaultConstant = 30

// probability returns the probability of rating1 winning
func probability(ratingA, ratingB float64) float64 {
	return float64(1.0) * 1.0 / (1 + 1.0*math.Pow(10, 1.0*(ratingA-ratingB)/400))
}

// NewElo is a wrapper func for the NewEloWithConstant func, however we use the EloDefaultConstant value
func NewElo(currentRating, opponentRating float64, playerWon custom.Bool) float64 {
	return NewEloWithConstant(currentRating, opponentRating, EloDefaultConstant, playerWon)
}

func NewEloWithConstant(currentRating, opponentRating float64, constant int, playerWon custom.Bool) float64 {
	var newRating float64

	winningProbability := probability(opponentRating, currentRating)

	if playerWon {
		newRating = currentRating + float64(constant)*(1-winningProbability)
	} else {
		newRating = currentRating + float64(constant)*(0-winningProbability)
	}

	return math.Round(newRating*1.0) / 1.0
}
