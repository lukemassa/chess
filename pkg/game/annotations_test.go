package game

import (
	"fmt"
	"testing"
)

func TestCoordinateNotation(t *testing.T) {
	// White pawn at E2, black Queen at E4, white to move
	testCases := []struct {
		notation            string
		expectedMove        bool
		expectedErrorString string
	}{
		{
			"Foobar",
			false,
			"Unexpected number of dashes: 0",
		},
		{
			"a-b-c-d",
			false,
			"Unexpected number of dashes: 3",
		},
		{
			"a-b",
			false,
			"Error parsing first coordinate a: Invalid location: a: Invalid number of characters",
		},
		{
			"a1-b",
			false,
			"Error parsing second coordinate b: Invalid location: b: Invalid number of characters",
		},
		{
			"a1-b1",
			false,
			"There is no piece in position A1",
		},
		{
			"e2-e8",
			false,
			"{E2 E8} is not a legal move",
		},
		{
			"e4-e5",
			false,
			"E4 is not owned by white",
		},
		{
			"e2-e3",
			true,
			"",
		},
	}
	c := CoordinateNotation{}
	b := BlankBoard(true)
	b.AddPiece(&Piece{
		PieceType: Pawn{},
		Color:     White,
		Location:  MustParseLocation("E2"),
	})
	b.AddPiece(&Piece{
		PieceType: Queen{},
		Color:     Black,
		Location:  MustParseLocation("E4"),
	})
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("test %s", tc.notation), func(t *testing.T) {
			actualErrorString := ""
			actualMove, err := c.ConvertToMove(b, tc.notation, White)
			if err != nil {
				actualErrorString = err.Error()
			}
			if actualErrorString != tc.expectedErrorString {
				t.Errorf("Expected error string %s, got %s", tc.expectedErrorString, actualErrorString)
			}
			if tc.expectedMove && actualMove == nil {
				t.Error("Expected a move but move is nil")
			}
			if !tc.expectedMove && actualMove != nil {
				t.Error("Expected no move but move is not nil")
			}

		})
	}
}
