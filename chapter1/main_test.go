package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMainFunction(t *testing.T) {
	var plays = map[string]play{
		"hamlet": {
			name:     "Hamlet",
			playType: "tragedy",
		},
		"as-like": {
			name:     "As You Like It",
			playType: "comedy",
		},
		"othello": {
			name:     "Othello",
			playType: "tragedy",
		},
	}

	var performances = []performance{
		{
			playID:   "hamlet",
			audience: 55,
		},
		{
			playID:   "as-like",
			audience: 35,
		},
		{
			playID:   "othello",
			audience: 40,
		},
	}

	invoice := invoice{
		customer:     "BigCo",
		performances: performances,
	}
	result := statement(invoice, plays)
	expected := `Statement for BigCo
 Hamlet: $650 (55) seats
 As You Like It: $580 (35) seats
 Othello: $500 (40) seats
Amount owed is $1730
You earned  47 credits
`
	assert.Equal(t, expected, result)
}
