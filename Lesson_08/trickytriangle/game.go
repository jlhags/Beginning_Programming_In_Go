package trickytriangle

import (
	"errors"
	"fmt"
)

var potential_moves [15][]int = [15][]int{
	{3, 5},
	{6, 8},
	{7, 9},
	{5, 10, 12, 0},
	{11, 13},
	{0, 3, 12, 14},
	{1, 8},
	{2, 9},
	{6, 1},
	{2, 7},
	{3, 12},
	{13, 4},
	{10, 3, 14, 5},
	{11, 4},
	{12, 5},
}

type Move struct {
	To   int
	From int
}

type Game struct {
	board Board
	moves []Move
}

func (g *Game) Start(openSpot int) {
	for i := 0; i < len(g.board.Holes); i++ {
		if i == openSpot {
			g.board.Holes[i].RemovePeg()
		} else {
			g.board.Holes[i].AddPeg()
		}
	}
}

func (g *Game) Move(from int, to int) error {
	if g.isValidMove(from, to) {
		center := (from + to) / 2
		g.board.Holes[from].RemovePeg()
		g.board.Holes[center].RemovePeg()
		g.board.Holes[to].AddPeg()
		g.moves = append(g.moves, Move{From: from, To: to})
		return nil
	}
	return errors.New("not a valid move")
}

func (g Game) DisplayBoard() {
	fmt.Println(g.board)
}

func (g Game) Score() int {
	return g.board.Remaining()
}

func (g Game) isValidMove(from int, to int) bool {
	if g.board.Holes[from].Empty() || !g.board.Holes[to].Empty() {
		return false
	}
	for _, s := range potential_moves[to] {
		center := (from + to) / 2
		if s == from && !g.board.Holes[center].Empty() {
			return true
		}
	}
	return false
}

func (g Game) GameOver() bool {
	for i, h := range g.board.Holes {
		if h.Empty() {
			for _, s := range potential_moves[i] {
				if !g.board.Holes[(s+i)/2].Empty() && !g.board.Holes[s].Empty() {
					return false
				}
			}
		}
	}
	return true
}

func (g Game) PotentialMoves(to int) []int {
	var ret []int
	for _, from := range potential_moves[to] {
		if g.isValidMove(from, to) {
			ret = append(ret, from)
		}
	}
	return ret
}
