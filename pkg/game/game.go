package game

import "fmt"

// Player players of the game
type Player string

const (
	// Black player controlling the black pieces
	Black Player = "black"
	// White player controlling the white pieces
	White Player = "white"
)

// Game current game
type Game struct {
	Board     *Board
	whiteTurn bool
}

// Print print the game
func (g *Game) Print() {
	g.Board.Print()
	fmt.Printf("%s to play\n", g.Turn())
}

// Turn whose turn is it
func (g *Game) Turn() Player {
	if g.whiteTurn {
		return White
	}
	return Black
}

// New get a new game of chess
func New() *Game {
	return &Game{
		Board:     NewBoard(),
		whiteTurn: true,
	}
}
