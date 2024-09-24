package main

import (
	"math"

	"github.com/Battle-Bunker/cyphid-snake/agent"
	"github.com/BattlesnakeOfficial/rules"
)

// HeuristicAvoidSnakes calculates a score that rewards positions farther from other snakes
func HeuristicAvoidSnakes(snapshot agent.GameSnapshot) float64 {
	var totalScore float64
	allSnakes := snapshot.Snakes()

	for _, allySnake := range snapshot.YourTeam() {
		head := allySnake.Head()
		snakeScore := 0.0

		for _, otherSnake := range allSnakes {
			if otherSnake.ID() == allySnake.ID() {
				continue // Skip self
			}

			for _, bodyPart := range otherSnake.Body() {
				distance := manhattanDistance(head, bodyPart)

				// Add to score based on distance. Closer snakes have more impact.
				snakeScore += 1.0 / math.Max(float64(distance), 1.0)
			}
		}

		// Invert the score so that being far from snakes gives a higher score
		snakeScore = 1.0 / (snakeScore + 1.0)
		totalScore += snakeScore
	}

	return totalScore * 1000 // Scale the score for better differentiation
}

// manhattanDistance calculates the Manhattan distance between two points
func manhattanDistance(p1, p2 rules.Point) int {
	return abs(p1.X-p2.X) + abs(p1.Y-p2.Y)
}

// abs returns the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}