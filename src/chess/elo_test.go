package chess

import (
	"github.com/Jacobbrewer1/chess-boards/src/custom"
	"testing"
)

func TestNewElo(t *testing.T) {
	tests := []struct {
		name                string
		currentPlayerRating float64
		opponentRating      float64
		playerWon           bool
		expected            float64
	}{
		{"Two new players, winner: current", EloInitialValue, EloInitialValue, true, 1515},
		{"Two new players, winner: opponent", EloInitialValue, EloInitialValue, false, 1485},
		{"New vs established strong player, winner: current", EloInitialValue, 1800, true, 1525},
		{"New vs established strong player, winner: opponent", EloInitialValue, 1800, false, 1495},
		{"Established strong player vs new, winner: current", 1800, EloInitialValue, true, 1805},
		{"Established strong player vs new, winner: opponent", 1800, EloInitialValue, false, 1775},
		{"New vs established player, winner: current", EloInitialValue, 1800, true, 1525},
		{"New vs established player, winner: opponent", EloInitialValue, 1800, false, 1495},
		{"Established weaker player vs new, winner: opponent", 900, EloInitialValue, true, 929},
		{"Established weaker player vs new, winner: opponent", 900, EloInitialValue, false, 899},
		{"New vs established weaker player, winner: opponent", EloInitialValue, 900, true, 1501},
		{"New vs established weaker player, winner: opponent", EloInitialValue, 900, false, 1471},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotA := NewElo(tt.currentPlayerRating, tt.opponentRating, custom.Bool(tt.playerWon))
			if gotA != tt.expected {
				t.Errorf("got = %f, expected %f", gotA, tt.expected)
				return
			}
		})
	}
}

func TestNewEloWithConstant(t *testing.T) {
	tests := []struct {
		name                string
		currentPlayerRating float64
		opponentRating      float64
		constant            int
		playerWon           bool
		expected            float64
	}{
		// constant of 20
		{"current: 1200, opponent: 1000, const: 200, Winner: current", 1200, 1000, 20, true, 1205},
		{"current: 1200, opponent: 1000, const: 200, Winner: opponent", 1200, 1000, 20, false, 1185},
		{"current: 1000, opponent: 1200, const: 200, Winner: current", 1000, 1200, 20, true, 1015},
		{"current: 1000, opponent: 1200, const: 200, Winner: opponent", 1000, 1200, 20, false, 995},
		// constant of 30
		{"current: 1200, opponent: 1000, const: 30, Winner: current", 1200, 1000, 30, true, 1207},
		{"current: 1200, opponent: 1000, const: 30, Winner: opponent", 1200, 1000, 30, false, 1177},
		{"current: 1000, opponent: 1200, const: 30, Winner: current", 1000, 1200, 30, true, 1023},
		{"current: 1000, opponent: 1200, const: 30, Winner: opponent", 1000, 1200, 30, false, 993},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotA := NewEloWithConstant(tt.currentPlayerRating, tt.opponentRating, tt.constant, custom.Bool(tt.playerWon))
			if gotA != tt.expected {
				t.Errorf("got = %f, expected %f", gotA, tt.expected)
				return
			}
		})
	}
}
