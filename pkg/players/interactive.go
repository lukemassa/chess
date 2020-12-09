package players

// An interactive chess player
import (
	"github.com/lukemassa/chess/pkg/game"
)

// InteractivePlayer a player of chess
type InteractivePlayer struct {
}

// NextMove implementation
func (i *InteractivePlayer) NextMove(b *game.Board, c game.Color) game.Move {
	return game.Move{
		Piece:       *b.Squares[4][1],
		Destination: game.NewLocation("E4"),
	}
}
