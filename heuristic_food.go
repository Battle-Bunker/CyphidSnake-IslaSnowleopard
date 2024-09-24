package main

import (
	"github.com/Battle-Bunker/cyphid-snake/agent"
	"github.com/BattlesnakeOfficial/rules"
)

// HeuristicFoodCollection calculates a score based on the amount of food collected and proximity to next food
func HeuristicFoodCollection(snapshot agent.GameSnapshot) float64 {
	var totalScore float64

	for _, allySnake := range snapshot.YourTeam() {
		// Score based on food eaten (snake length - initial length)
		foodEaten := allySnake.Length() - 3 // Assuming initial length is 3
		foodScore := float64(foodEaten) * 10 // Weight food eaten more heavily

		// Score based on proximity to next food
		snakeHead := allySnake.Head()
		closestFoodDistance := float64(snapshot.Width() + snapshot.Height()) // Initialize with max possible distance

		for _, food := range snapshot.Food() {
			distance := manhattanDistance(snakeHead, food)
			if float64(distance) < closestFoodDistance {
				closestFoodDistance = float64(distance)
			}
		}

		// Invert the distance so that closer food gives a higher score
		// Add 1 to avoid division by zero and to give a bonus for being on food
		proximityScore := 1.0 / (closestFoodDistance + 1.0)

		// Combine food eaten score and proximity score
		snakeScore := foodScore + proximityScore
		totalScore += snakeScore
	}

	return totalScore * 100 // Scale the score for better differentiation
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