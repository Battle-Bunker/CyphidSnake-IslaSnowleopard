package main

import (
	"github.com/Battle-Bunker/cyphid-snake/agent"
)


// TODO implement a heuristic that returns higher values when closer to food
func HeuristicFood(snapshot agent.GameSnapshot) float64 {
	headx:=snapshot.You().Head().X
	heady:=snapshot.You().Head().Y
	foodx:=snapshot.Food()[1].X
	foody:=snapshot.Food()[1].Y
	distance := 300-  ((headx-foodx)*(headx-foodx)+(heady-foody)*(heady-foody))
	return float64(distance)

}
