package game

import (
	"fmt"
	"testing"
)

func TestMoveBoard(t *testing.T) {
	testCases := []struct {
		currentLocationString   string
		newLocationString       string
		opponentLocationString  string
		expectedOpponentRemains bool
	}{
		{
			"E4",
			"A4",
			"A4",
			false,
		},
		{
			"E4",
			"A4",
			"H4",
			true,
		},
	}

	for _, tc := range testCases {

		t.Run(fmt.Sprintf("Moving from %s to %s with an opponent piece at %s", tc.currentLocationString, tc.newLocationString, tc.opponentLocationString), func(t *testing.T) {
			board := BlankBoard(false) // Doing our own validation checks
			piece := Piece{
				PieceType: Rook{},
				Color:     White,
				Location:  NewLocation(tc.currentLocationString),
			}
			opponentOriginalLocation := NewLocation(tc.opponentLocationString)
			opponentPiece := Piece{
				PieceType: Queen{},
				Color:     Black,
				Location:  opponentOriginalLocation,
			}
			board.AddPiece(&piece)
			board.AddPiece(&opponentPiece)
			err := board.Validate()
			if err != nil {
				t.Errorf("Found error when validating board before move: %v", err)
			}
			move := Move{
				Piece:       &piece,
				Destination: NewLocation(tc.newLocationString),
			}
			board.MakeMove(&move)

			err = board.Validate()
			if err != nil {
				t.Errorf("Found error when validating board after move: %v", err)
			}
			if tc.expectedOpponentRemains {
				if *board.Squares[opponentOriginalLocation.file][opponentOriginalLocation.rank] != opponentPiece {
					t.Errorf("Expected opponent piece to remain on %s", opponentOriginalLocation)
				}
				if len(board.Pieces) != 2 {
					t.Errorf("Expected both opponent and piece to remain on the board")
				}
			}
			if !tc.expectedOpponentRemains {
				if *board.Squares[opponentOriginalLocation.file][opponentOriginalLocation.rank] != piece {
					t.Errorf("Expected piece to now be where opponent was on %s", opponentOriginalLocation)
				}
				if len(board.Pieces) != 1 {
					t.Errorf("Expected both opponent and piece to remain on the board")
				}
			}
		})
	}

}
