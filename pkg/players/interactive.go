package players

// An interactive chess player
import (
	"github.com/lukemassa/chess/pkg/game"
)

// InteractivePlayer a player of chess
type InteractivePlayer struct {
}

// NextMove implemenation
func (i *InteractivePlayer) NextMove(b *game.Board) game.Move {
	return game.Move{
		Piece:       *b.Squares[4][3],
		Destination: game.NewLocation("E4"),
	}
}
