package main

import (
	"github.com/jlhags/tutorials/lesson_08/trickytriangle"
)

func main() {
	game := trickytriangle.Game{}
	game.Start(14)
	solveGame(&game)
}

func solveGame(g *trickytriangle.Game) bool {
	if g.GameOver() {
		if g.Score()==1 {
			g.PrintMoves()
			return true
		}
		return false
	}
	for to := 0; to < 15; to++ {
		pm := g.PotentialMoves(to)
		for _, from := range pm {
			g2 := *g
			g2.Move(from, to)
			if solveGame(&g2) {
				return true
			}
		}
	}
	return false
}
