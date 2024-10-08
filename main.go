package main

import (
	"github.com/Battle-Bunker/cyphid-snake/agent"
	"github.com/Battle-Bunker/cyphid-snake/server"
	"github.com/BattlesnakeOfficial/rules/client"
)

func main() {

	metadata := client.SnakeMetadataResponse{
		APIVersion: "1",
		Author:     "",
		Color:      "#FFD700",
		Head:       "default",
		Tail:       "default",
	}

	portfolio := agent.NewPortfolio(
		agent.NewHeuristic(1.0, "team-health", HeuristicHealth),
		agent.NewHeuristic(1.0, "centre proximity", HeuristicCenterProximity),
		agent.NewHeuristic(1.0, "food", HeuristicFoodCollection),
	)

	snakeAgent := agent.NewSnakeAgent(portfolio, metadata)
	server := server.NewServer(snakeAgent)

	server.Start()
}
