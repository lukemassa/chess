package game

import "fmt"

// Player player of chess
// This is implemented by whoever wants to "play" the game
type Player interface {
	NextMove(b *Board) Move
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
	whitesTurn  bool
	WhitePlayer Player
	BlackPlayer Player
}

func (g *Game) String() string {
	return fmt.Sprintf("%s\n%s to play\n", g.Board, g.Turn())
}

// Turn whose turn is it
func (g *Game) Turn() Player {
	if g.whitesTurn {
		return g.WhitePlayer
	}
	return g.BlackPlayer
}

// Play the game
func (g *Game) Play() {
	for {
		player := g.Turn()
		g.whitesTurn = !g.whitesTurn
		move := player.NextMove(g.Board)
		g.Board.MakeMove(move)
		err := g.Board.Validate()
		if err != nil {
			//log.Fatalf("Error validating %v", err)
		}
		fmt.Printf("%s", g)
		break
	}
}

// New get a new game of chess
func New(whitePlayer, blackPlayer Player) *Game {
	return &Game{
		Board:       NewBoard(),
		whitesTurn:  true,
		WhitePlayer: whitePlayer,
		BlackPlayer: blackPlayer,
	}
}
