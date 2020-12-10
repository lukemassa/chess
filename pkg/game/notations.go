package game

import (
	"fmt"
	"strings"
)

// Notation a chess
type Notation interface {
	ConvertToMove(b *Board, notation string, color Color) (*Move, error)
}

// AlgebraicNotation for chess, like "NH1"
type AlgebraicNotation struct {
}

// CoordinateNotation for chess, like "E2-E3"
type CoordinateNotation struct {
}

// ConvertToMove convert an coordinate string (like "D2-D3") into a move
// If poorly formed or not a legal move, return an error
func (c CoordinateNotation) ConvertToMove(b *Board, notation string, color Color) (*Move, error) {

	coordinates := strings.Split(notation, "-")
	if len(coordinates) != 2 {
		return nil, fmt.Errorf("Unexpected number of dashes: %d", len(coordinates)-1)
	}
	current, err := ParseLocation(coordinates[0])
	if err != nil {
		return nil, fmt.Errorf("Error parsing first coordinate %s: %s", coordinates[0], err)
	}
	destination, err := ParseLocation(coordinates[1])
	if err != nil {
		return nil, fmt.Errorf("Error parsing second coordinate %s: %s", coordinates[1], err)
	}
	piece := b.Squares[current.file][current.rank]
	if piece == nil {
		return nil, fmt.Errorf("There is no piece in position %s", current)
	}
	move := Move{
		Piece:       piece,
		Destination: destination,
	}
	if !move.IsValidMove(b) {
		return nil, fmt.Errorf("%s is not a legal move", move)
	}
	if move.Piece.Color != color {
		return nil, fmt.Errorf("%s is not owned by %s", piece, color)
	}
	return &move, nil
}
