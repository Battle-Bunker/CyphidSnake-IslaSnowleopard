package main

import (
  _"github.com/BattlesnakeOfficial/rules"
  "github.com/Battle-Bunker/cyphid-snake/agent"
  "math"
)

// HeuristicCenterProximity calculates a score based on how close the snakes are to the center of the board
func HeuristicCenterProximity(snapshot agent.GameSnapshot) float64 {
  centerX := float64(snapshot.Width()) / 2
  centerY := float64(snapshot.Height()) / 2
  maxDistance := math.Sqrt(centerX*centerX + centerY*centerY)

  totalScore := 0.0

  for _, allySnake := range snapshot.YourTeam() {
    if !allySnake.Alive() {
      continue
    }

    head := allySnake.Head()
    distanceFromCenter := math.Sqrt(math.Pow(float64(head.X)-centerX, 2) + math.Pow(float64(head.Y)-centerY, 2))

    // Calculate a score that's higher when closer to the center
    // We use maxDistance - distanceFromCenter so that being closer to the center yields a higher score
    snakeScore := maxDistance - distanceFromCenter

    // Scale the score by the snake's health to account for survival potential
    snakeScore *= float64(allySnake.Health())

    totalScore += snakeScore
  }

  return totalScore
}