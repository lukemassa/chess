package chess

import (
	"fmt"
	"testing"
)

func TestPieceMovements(t *testing.T) {

	testCases := []struct {
		pieceType             PieceType
		player                Player
		currentLocationString string
		newLocationString     string
		expectedIsValidMove   bool
	}{
		{
			King{},
			White,
			"E4",
			"E5",
			true,
		},
		{
			King{},
			White,
			"E4",
			"E6",
			false,
		},
		{
			Queen{},
			White,
			"E4",
			"E5",
			true,
		},
		{
			Queen{},
			White,
			"E4",
			"E6",
			true,
		},
		{
			Queen{},
			White,
			"E4",
			"C4",
			true,
		},
		{
			Queen{},
			White,
			"E4",
			"F5",
			true,
		},
		{
			Queen{},
			White,
			"E4",
			"F6",
			false,
		},
		{
			Bishop{},
			White,
			"E4",
			"F5",
			true,
		},
		{
			Bishop{},
			White,
			"E4",
			"E5",
			false,
		},
		{
			Pawn{},
			White,
			"E2",
			"E4",
			true,
		},
		{
			Pawn{},
			White,
			"E2",
			"D4",
			false,
		},
		{
			Pawn{},
			Black,
			"E2",
			"E4",
			false,
		},
		{
			Pawn{},
			White,
			"E3",
			"E4",
			true,
		},
		{
			Pawn{},
			White,
			"E3",
			"E5",
			false,
		},
		{
			Knight{},
			White,
			"E5",
			"E7",
			false,
		},
		{
			Knight{},
			White,
			"E5",
			"D7",
			true,
		},
		{
			Knight{},
			White,
			"E5",
			"C7",
			false,
		},
		{
			Knight{},
			White,
			"E5",
			"C6",
			true,
		},
	}

	board := BlankBoard()
	for _, tc := range testCases {
		prefix := "Can move"
		if !tc.expectedIsValidMove {
			prefix = "Can't move"
		}
		testName := fmt.Sprintf("%s %s %s from %s -> %s", prefix, tc.player, tc.pieceType.Name(), tc.currentLocationString, tc.newLocationString)
		t.Run(testName, func(t *testing.T) {
			piece := Piece{
				PieceType: tc.pieceType,
				Player:    tc.player,
				Location:  NewLocation(tc.currentLocationString),
			}
			actualIsValidMove := piece.IsValidMove(NewLocation(tc.newLocationString), board)

			if actualIsValidMove && !tc.expectedIsValidMove {
				t.Errorf("Expected invalid move")
			}
			if !actualIsValidMove && tc.expectedIsValidMove {
				t.Errorf("Expected valid move")
			}

		})
	}
}
