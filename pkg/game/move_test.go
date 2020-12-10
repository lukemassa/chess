package game

import (
	"fmt"
	"testing"
)

func TestCapture(t *testing.T) {
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
				Location:  MustParseLocation(tc.currentLocationString),
			}
			opponentOriginalLocation := MustParseLocation(tc.opponentLocationString)
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
				Destination: MustParseLocation(tc.newLocationString),
			}
			board.MakeMove(&move, White)

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

func TestPiecesBetween(t *testing.T) {
	testCases := []struct {
		location1            string
		location2            string
		pieceLocation        string
		expectedPieceBetween bool
	}{
		{
			"A1",
			"A8",
			"A4",
			true,
		},
		{
			"A8",
			"A1",
			"A4",
			true,
		},
		{
			"A1",
			"A8",
			"B4",
			false,
		},
		{
			"A1",
			"H1",
			"D1",
			true,
		},
		{
			"H1",
			"A1",
			"D1",
			true,
		},
		{
			"A1",
			"H8",
			"D4",
			true,
		},
		{
			"H8",
			"A1",
			"D4",
			true,
		},
		{
			"A1",
			"H8",
			"D2",
			false,
		},
		{
			"A8",
			"H1",
			"D5",
			true,
		},
		{
			"H1",
			"A8",
			"D5",
			true,
		},
		{
			"A1",
			"H8",
			"D1",
			false,
		},
	}

	for _, tc := range testCases {

		prefix := "Do not expect"
		if tc.expectedPieceBetween {
			prefix = "Expect"
		}
		t.Run(fmt.Sprintf("%s piece %s is between %s and %s", prefix, tc.pieceLocation, tc.location1, tc.location2), func(t *testing.T) {
			board := BlankBoard(true)
			piece := Piece{
				PieceType: Rook{},
				Color:     White,
				Location:  MustParseLocation(tc.pieceLocation),
			}
			board.AddPiece(&piece)

			actualPieceBetween := board.PieceBetween(MustParseLocation(tc.location1), MustParseLocation(tc.location2))
			if actualPieceBetween != tc.expectedPieceBetween {
				t.Errorf("Expected actual piece between %v, found %v", tc.expectedPieceBetween, actualPieceBetween)
			}
		})
	}

}
