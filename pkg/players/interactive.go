package players

// An interactive chess player
import (
	"fmt"

	"github.com/lukemassa/chess/pkg/game"
)

// InteractivePlayer a player of chess
type InteractivePlayer struct {
}

// NextMove implementation
func (i *InteractivePlayer) NextMove(board *game.Board, color game.Color) game.Move {
	fmt.Printf("%s\n", board)
	var moveString string
	for {
		fmt.Printf("%s move: ", color)
		fmt.Scanln(&moveString)
		move, err := board.ConvertAlgebraicToMove(moveString)
		if err == nil {
			return move
		}
		fmt.Printf("Invalid move: %s\n", err)
	}
}
