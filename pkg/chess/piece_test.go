package chess

import (
	"fmt"
	"testing"
)

func TestBasicPieceMovements(t *testing.T) {

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
			"F3",
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

func TestPawnCapturingMovement(t *testing.T) {

	// Check to make sure the pawn is moving around correctly
	// Use a white pawn starting on E2
	testCases := []struct {
		newLocationString      string
		opponentLocationString string
		expectedIsValidMove    bool
	}{
		{
			"E4",
			"A8",
			true,
		},
		// Cannot capture forward with a pawn
		{
			"E4",
			"E4",
			false,
		},
		// Cannot move diagonally unless capturing
		{
			"D3",
			"E4",
			false,
		},
		// Can capture diagonally
		{
			"D3",
			"D3",
			true,
		},
	}

	for _, tc := range testCases {

		prefix := "Can move"
		if !tc.expectedIsValidMove {
			prefix = "Can't move"
		}
		testName := fmt.Sprintf("%s pawn to %s with opponent on %s", prefix, tc.newLocationString, tc.opponentLocationString)
		t.Run(testName, func(t *testing.T) {
			board := BlankBoard()
			opponentPiece := Piece{
				PieceType: Queen{},
				Player:    Black,
				Location:  NewLocation(tc.opponentLocationString),
			}
			board.AddPiece(opponentPiece)

			piece := Piece{
				PieceType: Pawn{},
				Player:    White,
				Location:  NewLocation("E2"),
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
