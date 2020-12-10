package game

import (
	"errors"
	"fmt"
	"strings"

	"log"
)

// BoardSize number of squares along one edge of the board
const BoardSize = 8

// Board an abstraction for the current setup of the board
type Board struct {
	Pieces []*Piece
	Squares
	validate bool
}

// Squares a description of what squares are filled how
type Squares [][]*Piece

// Location on the board
type Location struct {
	// A-F
	file int8
	// 1-8
	rank int8
}

func (l Location) String() string {
	return fmt.Sprintf("%c%c", l.file+'A', l.rank+'1')
}

// MustParseLocation exits the program if it can't parse the location
// Use sparingly/only for literals
func MustParseLocation(locationString string) Location {
	location, err := ParseLocation(locationString)
	if err != nil {
		log.Fatal(err)
	}
	return location
}

func uppercase(r rune) rune {
	// a little copying is better than a little dependency
	return r + ('A' - 'a')
}

// ParseLocation get a new location based on the location string
func ParseLocation(locationString string) (Location, error) {

	var location Location
	r := []rune(locationString)
	// hm maybe not panic? Not sure what's best here
	if len(r) != 2 {
		return location, fmt.Errorf("Invalid location: %s: Invalid number of characters", locationString)
	}
	if r[0] >= 'a' && r[0] <= 'z' {
		r[0] = uppercase(r[0])
	}
	if r[0] < 'A' || r[0] > 'H' {
		return location, fmt.Errorf("Invalid location: %s: First character is not between A and H", locationString)
	}
	if r[1] < '1' || r[1] > '8' {
		return location, fmt.Errorf("Invalid location: %s: Second character is not betweeen 1 and 8", locationString)
	}
	fileAsInt := int8(r[0]) - 'A'
	rankAsInt := int8(r[1]) - '1'
	return Location{
		file: fileAsInt,
		rank: rankAsInt,
	}, nil
}

// IsValidMove can a given move be made
func (b *Board) IsValidMove(move Move) bool {
	return move.Piece.IsValidMove(move.Destination, b)
}

// MakeMove make a specific move
// returns true if this move won the game
func (b *Board) MakeMove(move *Move, c Color) bool {
	b.FailOnInvalidMove(move, c)
	b.FailOnInvalidBoard()
	// If there's a piece there, remove it
	currentPiece := b.Squares[move.Destination.file][move.Destination.rank]
	if currentPiece != nil {
		b.Squares[move.Destination.file][move.Destination.rank] = nil
		for i := 0; i < len(b.Pieces); i++ {
			if *b.Pieces[i] == *currentPiece {
				// Remove this piece
				b.Pieces[i] = b.Pieces[len(b.Pieces)-1]
				b.Pieces = b.Pieces[:len(b.Pieces)-1]
				break
			}
		}
	}
	// Remove the pointer from the old place, add the pointer at the new place
	b.Squares[move.Piece.Location.file][move.Piece.Location.rank] = nil
	b.Squares[move.Destination.file][move.Destination.rank] = move.Piece

	// Update this piece's location
	move.Piece.Location = move.Destination

	// If validation is set on, make sure we left ourselves with a valid board
	b.FailOnInvalidBoard()

	// TODO: Implement check for end of game
	return false
}

// ValidateMove see if this is a valid move
func (b *Board) ValidateMove(move *Move) error {
	// TODO: Should this code go into IsValidMove?
	if move.Piece == nil {
		return errors.New("Null piece in move")
	}
	if !move.Piece.IsValidMove(move.Destination, b) {
		return errors.New("Is not a legal move")
	}
	return nil
}

// Validate check to make sure we are looking at a legal board
func (b *Board) Validate() error {
	foundLocations := make(map[Location]bool)
	for i := 0; i < len(b.Pieces); i++ {

		l := b.Pieces[i].Location
		if _, ok := foundLocations[l]; ok {
			return fmt.Errorf("Found more than one piece at %s", l)
		}
		foundLocations[l] = true

		inSquare := b.Squares[l.file][l.rank]
		if inSquare == nil {
			return fmt.Errorf("Piece %s is nil in the square", b.Pieces[i])
		}
		if *b.Squares[l.file][l.rank] != *b.Pieces[i] {
			return fmt.Errorf("Piece %s not reflected correctly in squares", b.Pieces[i])
		}
	}

	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			piece := b.Squares[i][j]
			if piece == nil {
				continue
			}
			foundPieces := 0
			for k := 0; k < len(b.Pieces); k++ {
				if *b.Pieces[k] == *piece {
					foundPieces++
				}
			}
			if foundPieces == 0 {
				return fmt.Errorf("There is a piece at %s that is not in piece list", piece.Location)
			}
			if foundPieces > 1 {
				return fmt.Errorf("Piece %s is referenced on more than one square", piece)
			}
		}
	}
	return nil
}

// FailOnInvalidBoard quit if the board is now invalid
func (b *Board) FailOnInvalidBoard() {
	if !b.validate {
		return
	}
	err := b.Validate()
	if err != nil {
		log.Fatalf("Board is in invalid state: %s", err)
	}
}

// FailOnInvalidMove quit if the move is invalid
func (b *Board) FailOnInvalidMove(move *Move, color Color) {
	if !b.validate {
		return
	}
	if move.Piece.Color != color {
		log.Fatalf("Player %s moved piece of color %s", color, move.Piece.Color)
	}
	err := b.ValidateMove(move)
	if err != nil {
		log.Fatalf("Move is in invalid: %s", err)
	}
}

// AddPiece add a piece onto the board
func (b *Board) AddPiece(piece *Piece) {

	// Board will be invalid between these two commands but not at the beginning or end
	b.FailOnInvalidBoard()
	b.Pieces = append(b.Pieces, piece)
	b.Squares[piece.file][piece.rank] = piece
	b.FailOnInvalidBoard()
}

func getEmptySquares() Squares {
	//Start with an empty board
	ret := make([][]*Piece, BoardSize)
	for i := 0; i < BoardSize; i++ {
		ret[i] = make([]*Piece, BoardSize)
	}
	return ret
}

func (s Squares) String() string {
	ret := strings.Builder{}
	ret.WriteRune(' ')
	for i := 0; i < BoardSize; i++ {
		ret.WriteRune(rune(i) + 'A')
	}
	ret.WriteRune('\n')
	for i := BoardSize - 1; i >= 0; i-- {
		ret.WriteString(fmt.Sprintf("%d", i+1))
		for j := 0; j < BoardSize; j++ {
			piece := s[j][i]
			if piece == nil {
				ret.WriteRune(' ')
			} else {
				ret.WriteRune(piece.Symbol())
			}
		}
		ret.WriteString(fmt.Sprintf("%d", i+1))
		ret.WriteRune('\n')
	}
	ret.WriteRune(' ')
	for i := 0; i < BoardSize; i++ {
		ret.WriteRune(rune(i) + 'A')
	}
	ret.WriteRune('\n')
	return ret.String()
}

// Print print the current board setup
func (b *Board) String() string {
	//	b.FailOnInvalidBoard()
	return fmt.Sprintf("%s", b.Squares)
}

// BlankBoard an empty board
func BlankBoard(validate bool) *Board {

	b := Board{
		Squares:  getEmptySquares(),
		validate: validate,
	}
	b.FailOnInvalidBoard()
	return &b
}

// NewBoard a new board setup for standard play
func NewBoard(validate bool) *Board {

	b := BlankBoard(validate)

	pieces := []Piece{
		{PieceType: Pawn{}, Color: White, Location: MustParseLocation("A2")},
		{PieceType: Pawn{}, Color: White, Location: MustParseLocation("B2")},
		{PieceType: Pawn{}, Color: White, Location: MustParseLocation("C2")},
		{PieceType: Pawn{}, Color: White, Location: MustParseLocation("D2")},
		{PieceType: Pawn{}, Color: White, Location: MustParseLocation("E2")},
		{PieceType: Pawn{}, Color: White, Location: MustParseLocation("F2")},
		{PieceType: Pawn{}, Color: White, Location: MustParseLocation("G2")},
		{PieceType: Pawn{}, Color: White, Location: MustParseLocation("H2")},

		{PieceType: Pawn{}, Color: Black, Location: MustParseLocation("A7")},
		{PieceType: Pawn{}, Color: Black, Location: MustParseLocation("B7")},
		{PieceType: Pawn{}, Color: Black, Location: MustParseLocation("C7")},
		{PieceType: Pawn{}, Color: Black, Location: MustParseLocation("D7")},
		{PieceType: Pawn{}, Color: Black, Location: MustParseLocation("E7")},
		{PieceType: Pawn{}, Color: Black, Location: MustParseLocation("F7")},
		{PieceType: Pawn{}, Color: Black, Location: MustParseLocation("G7")},
		{PieceType: Pawn{}, Color: Black, Location: MustParseLocation("H7")},

		{PieceType: Rook{}, Color: White, Location: MustParseLocation("A1")},
		{PieceType: Knight{}, Color: White, Location: MustParseLocation("B1")},
		{PieceType: Bishop{}, Color: White, Location: MustParseLocation("C1")},
		{PieceType: Queen{}, Color: White, Location: MustParseLocation("D1")},
		{PieceType: King{}, Color: White, Location: MustParseLocation("E1")},
		{PieceType: Bishop{}, Color: White, Location: MustParseLocation("F1")},
		{PieceType: Knight{}, Color: White, Location: MustParseLocation("G1")},
		{PieceType: Rook{}, Color: White, Location: MustParseLocation("H1")},

		{PieceType: Rook{}, Color: Black, Location: MustParseLocation("A8")},
		{PieceType: Knight{}, Color: Black, Location: MustParseLocation("B8")},
		{PieceType: Bishop{}, Color: Black, Location: MustParseLocation("C8")},
		{PieceType: Queen{}, Color: Black, Location: MustParseLocation("D8")},
		{PieceType: King{}, Color: Black, Location: MustParseLocation("E8")},
		{PieceType: Bishop{}, Color: Black, Location: MustParseLocation("F8")},
		{PieceType: Knight{}, Color: Black, Location: MustParseLocation("G8")},
		{PieceType: Rook{}, Color: Black, Location: MustParseLocation("H8")},
	}
	for i := 0; i < len(pieces); i++ {
		b.AddPiece(&pieces[i])
	}

	return b
}
