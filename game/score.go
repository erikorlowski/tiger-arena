// Copyright 2023 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Model representing the instantaneous score of a match.

package game

type Score struct { // TIGER_TODO
	HighCones       int
	LowCones        int
	EndgameStatuses [1]EndgameStatus
	Fouls           []Foul
	PlayoffDq       bool
}

// Represents the state of a robot at the end of the match.
type EndgameStatus int

const (
	EndgameNone EndgameStatus = iota // TIGER_TODO
	EndgameParked
)

// Calculates and returns the summary fields used for ranking and display.
func (score *Score) Summarize(opponentScore *Score) *ScoreSummary { // TIGER_TODO
	summary := new(ScoreSummary)

	// Leave the score at zero if the alliance was disqualified.
	if score.PlayoffDq {
		return summary
	}

	// Calculate teleoperated period points.
	highConePoints := score.HighCones * 3
	lowConePoints := score.LowCones
	for i := 0; i < 1; i++ {
		switch score.EndgameStatuses[i] {
		case EndgameParked:
			summary.ParkPoints += 2
		}
	}

	summary.HighConePoints = highConePoints
	summary.LowConePoints = lowConePoints
	summary.MatchPoints = summary.HighConePoints + summary.LowConePoints + summary.ParkPoints

	// Calculate penalty points.
	for _, foul := range opponentScore.Fouls {
		summary.FoulPoints += foul.PointValue()
	}

	summary.Score = summary.MatchPoints + summary.FoulPoints

	return summary
}

// Returns true if and only if all fields of the two scores are equal.
func (score *Score) Equals(other *Score) bool { // TIGER_TODO
	if score.HighCones != other.HighCones ||
		score.LowCones != other.LowCones ||
		score.EndgameStatuses != other.EndgameStatuses ||
		score.PlayoffDq != other.PlayoffDq ||
		len(score.Fouls) != len(other.Fouls) {
		return false
	}

	for i, foul := range score.Fouls {
		if foul != other.Fouls[i] {
			return false
		}
	}

	return true
}
