package game

import (
	"fmt"
)

// Player player of chess
// This is implemented by whoever wants to "play" the game
type Player interface {
	NextMove(b *Board, c Color) *Move
}

// Color which of the two sides the player is on
type Color string

const (
	// Black the black pieces
	Black Color = "black"
	// White the white pieces
	White Color = "white"
)

// Game current game
type Game struct {
	Board       *Board
	WhitePlayer Player
	BlackPlayer Player
	whitesTurn  bool
	winner      Color
}

func (g Game) String() string {
	color, _ := g.Turn()
	return fmt.Sprintf("%s\n%s to play\n", g.Board, color)
}

// Turn whose turn is it
func (g Game) Turn() (Player, Color) {
	if g.whitesTurn {
		return g.WhitePlayer, White
	}
	return g.BlackPlayer, Black
}

// Play the game, return the color who won
// TODO: Handle draw
func (g *Game) Play() Color {
	var winner Color
	for {
		player, color := g.Turn()

		move := player.NextMove(g.Board, color)
		// Did this move succeed in ending the game?

		if g.Board.MakeMove(move, color) {
			winner = color
			break
		}

		g.whitesTurn = !g.whitesTurn
	}
	return winner
}

// IsOver is the game over
// TODO: Implement
func (g Game) IsOver() bool {
	return false
}

// Winner who is the winner
func (g Game) Winner() Color {
	if !g.IsOver() {
		panic("Called winner before the game is over")
	}
	return g.winner
}

// New get a new game of chess
func New(whitePlayer, blackPlayer Player, validate bool) *Game {
	return &Game{
		Board:       NewBoard(validate),
		whitesTurn:  true,
		WhitePlayer: whitePlayer,
		BlackPlayer: blackPlayer,
	}
}
