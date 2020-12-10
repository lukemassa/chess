package players

// An interactive chess player
import (
	"fmt"

	"github.com/lukemassa/chess/pkg/game"
)

// InteractivePlayer a player of chess
type InteractivePlayer struct {
	Notation game.Notation
}

// NextMove implementation
// Given a board, it presents the board to the user and asks them to enter a move
func (i *InteractivePlayer) NextMove(board *game.Board, color game.Color) *game.Move {
	fmt.Printf("%s\n", board)
	var moveString string
	for {
		fmt.Printf("%s move: ", color)
		fmt.Scanln(&moveString)
		if moveString == "" {
			continue
		}
		move, err := i.Notation.ConvertToMove(board, moveString)
		if err == nil {
			return move
		}
		fmt.Printf("Invalid move: %s\n", err)
	}
}
