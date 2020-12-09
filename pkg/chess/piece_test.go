package chess

import (
	"fmt"
	"testing"
)

func TestPieceMovements(t *testing.T) {

	testCases := []struct {
		pieceType             PieceType
		currentLocationString string
		newLocationString     string
		expectedIsValidMove   bool
	}{
		{
			King{},
			"E4",
			"E5",
			true,
		},
		{
			King{},
			"E4",
			"E6",
			false,
		},
	}

	for _, tc := range testCases {
		prefix := "Can move"
		if !tc.expectedIsValidMove {
			prefix = "Can't move"
		}
		testName := fmt.Sprintf("%s %s from %s -> %s", prefix, tc.pieceType.Name(), tc.currentLocationString, tc.newLocationString)
		t.Run(testName, func(t *testing.T) {
			piece := Piece{
				PieceType: tc.pieceType,
				Location:  NewLocation(tc.currentLocationString),
			}
			actualIsValidMove := piece.IsValidMove(NewLocation(tc.newLocationString))

			if actualIsValidMove && !tc.expectedIsValidMove {
				t.Errorf("Expected invalid move")
			}
			if !actualIsValidMove && tc.expectedIsValidMove {
				t.Errorf("Expected valid move")
			}

		})
	}
}
