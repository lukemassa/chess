package game

import "fmt"

// Color players of the game
type Color string

const (
	// Black player controlling the black pieces
	Black Color = "black"
	// White player controlling the white pieces
	White Color = "white"
)

// Game current game
type Game struct {
	Board       *Board
	whitesTurn  bool
	WhitePlayer Player
	BlackPlayer Player
}

// Print print the game
func (g *Game) Print() {
	g.Board.Print()
	fmt.Printf("%s to play\n", g.Turn())
}

// Turn whose turn is it
func (g *Game) Turn() Player {
	if g.whitesTurn {
		return g.WhitePlayer
	}
	return g.BlackPlayer
}

// New get a new game of chess
func New() *Game {
	return &Game{
		Board:      NewBoard(),
		whitesTurn: true,
	}
}
