package game

import (
	"fmt"
	"testing"
)

func TestBasicPieceMovements(t *testing.T) {

	testCases := []struct {
		pieceType             PieceType
		color                 Color
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
			Rook{},
			White,
			"E4",
			"E5",
			true,
		},
		{
			Rook{},
			White,
			"E4",
			"E6",
			true,
		},
		{
			Rook{},
			White,
			"E4",
			"C4",
			true,
		},
		{
			Rook{},
			White,
			"E4",
			"F5",
			false,
		},
		{
			Rook{},
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
			"E2",
			"E6",
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

	board := BlankBoard(true)
	for _, tc := range testCases {
		prefix := "Can move"
		if !tc.expectedIsValidMove {
			prefix = "Can't move"
		}
		testName := fmt.Sprintf("%s %s %s from %s -> %s", prefix, tc.color, tc.pieceType.Name(), tc.currentLocationString, tc.newLocationString)
		t.Run(testName, func(t *testing.T) {
			piece := Piece{
				PieceType: tc.pieceType,
				Color:     tc.color,
				Location:  MustParseLocation(tc.currentLocationString),
			}
			actualIsValidMove := piece.IsValidMove(MustParseLocation(tc.newLocationString), board)

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
		name                   string
		newLocationString      string
		opponentLocationString string
		expectedIsValidMove    bool
	}{
		{
			"normal pawn move",
			"E4",
			"A8",
			true,
		},
		{
			"Cannot capture forward with a pawn",
			"E4",
			"E4",
			false,
		},
		{
			"Cannot move diagonally unless capturing",
			"D3",
			"E4",
			false,
		},
		{
			"Can capture diagonally",
			"D3",
			"D3",
			true,
		},
		{
			"Can only capture diagonally one ahead",
			"D5",
			"D5",
			false,
		},
	}

	for _, tc := range testCases {

		prefix := "Can move"
		if !tc.expectedIsValidMove {
			prefix = "Can't move"
		}
		testName := fmt.Sprintf("%s pawn to %s with opponent on %s: %s", prefix, tc.newLocationString, tc.opponentLocationString, tc.name)
		t.Run(testName, func(t *testing.T) {
			board := BlankBoard(false) // Doing
			opponentPiece := Piece{
				PieceType: Queen{},
				Color:     Black,
				Location:  MustParseLocation(tc.opponentLocationString),
			}
			board.AddPiece(&opponentPiece)

			piece := Piece{
				PieceType: Pawn{},
				Color:     White,
				Location:  MustParseLocation("E2"),
			}
			actualIsValidMove := piece.IsValidMove(MustParseLocation(tc.newLocationString), board)

			if actualIsValidMove && !tc.expectedIsValidMove {
				t.Errorf("Expected invalid move")
			}
			if !actualIsValidMove && tc.expectedIsValidMove {
				t.Errorf("Expected valid move")
			}

		})
	}
}
