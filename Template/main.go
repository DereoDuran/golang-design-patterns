package main

import "fmt"

type Game interface {
	Start()
	TakeTurn()
	HaveWinner() bool
}

func PlayGame(game Game) {
	game.Start()

	for !game.HaveWinner() {
		game.TakeTurn()
	}

	fmt.Println("Game Over!")
}

type chess struct {
	turn, maxTurns, currentPlayer int
}

func (c *chess) Start() {
	fmt.Printf("Starting a game of chess with %d turns", c.maxTurns)
	c.turn = 1
	c.maxTurns = 10
	c.currentPlayer = 1
}

func (c *chess) TakeTurn() {
	c.turn++
	c.currentPlayer = (c.currentPlayer + 1) % 2
}

func (c *chess) HaveWinner() bool {
	return c.turn == c.maxTurns
}

func NewGameOfChess() Game {
	return &chess{}
}

func main() {

}
