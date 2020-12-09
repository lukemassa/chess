package game

import (
	"fmt"
	"testing"
)

func TestMoveBoard(t *testing.T) {
	testCases := []struct {
		currentLocationString string

		newLocationString      string
		opponentLocationString string
		expectedNumberOfPieces int
	}{
		{
			"E4",
			"A4",
			"A4",
			1,
		},
		{
			"E4",
			"A4",
			"H4",
			2,
		},
	}

	for _, tc := range testCases {

		t.Run(fmt.Sprintf("Moving from %s to %s with an opponent piece at %s", tc.currentLocationString, tc.newLocationString, tc.opponentLocationString), func(t *testing.T) {
			board := BlankBoard()
			piece := Piece{
				PieceType: Rook{},
				Color:     White,
				Location:  NewLocation(tc.currentLocationString),
			}
			opponentPiece := Piece{
				PieceType: Queen{},
				Color:     Black,
				Location:  NewLocation(tc.opponentLocationString),
			}
			board.AddPiece(piece)
			board.AddPiece(opponentPiece)
			err := board.Validate()
			if err != nil {
				t.Errorf("Found error when validating board: %v", err)
			}
			move := Move{
				Piece:       piece,
				Destination: NewLocation(tc.newLocationString),
			}
			board.MakeMove(move)
			if len(board.Pieces) != tc.expectedNumberOfPieces {
				t.Errorf("Expected %d pieces after the move, found %d", tc.expectedNumberOfPieces, len(board.Pieces))
			}
			err = board.Validate()
			if err != nil {
				t.Errorf("Found error when validating board: %v", err)
			}
		})
	}

}
