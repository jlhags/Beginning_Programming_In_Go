package main

import (
	"github.com/jlhags/tutorials/lesson_07/trickytriangle"
)

func main() {
	game := trickytriangle.Game{}
	game.Start(14)
	solveGame(&game)
	game.DisplayBoard()
}

func solveGame(g *trickytriangle.Game) bool {
	if g.GameOver() {
		return g.Score() == 1
	}
	for to := 0; to < 15; to++ {
		pm := g.PotentialMoves(to)
		for _, from := range pm {
			g2 := *g
			g2.Move(from, to)
			if solveGame(&g2) {
				g2.DisplayBoard()
				return true
			}
		}
	}
	return false
}
