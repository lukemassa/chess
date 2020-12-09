package chess

import (
	"fmt"
	"log"
	"strings"
)

// BoardSize number of squares along one edge of the board
const BoardSize = 8

// Board an abstraction for the current setup of the board
type Board struct {
	Pieces []Piece
	Squares
}

// boardMap a description of what squares are filled how
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

// NewLocation get a new location
func NewLocation(locationString string) Location {

	var fileAsRune rune
	var rankAsRune rune
	for i, c := range locationString {
		if i == 0 {
			fileAsRune = c
		} else if i == 1 {
			rankAsRune = c
		} else {
			panic(fmt.Sprintf("Invalid location: %s", locationString))
		}
	}
	fileAsInt := int8(fileAsRune) - 'A'
	rankAsInt := int8(rankAsRune) - '1'
	return Location{
		file: fileAsInt,
		rank: rankAsInt,
	}
}

// Validate check to make sure we are looking at a legal board
func (b *Board) Validate() error {
	foundLocations := make(map[Location]bool)
	for i := 0; i < len(b.Pieces); i++ {
		l := b.Pieces[i].Location
		if _, ok := foundLocations[l]; ok {
			return fmt.Errorf("Found more than one piece at %s", l)
		}
		foundLocations[b.Pieces[i].Location] = true
	}
	return nil
}

// AddPiece add a piece onto the board
func (b *Board) AddPiece(piece Piece) {
	b.Pieces = append(b.Pieces, piece)
	b.Squares[piece.file][piece.rank] = &piece
}

// Potential todo: don't make this a map, have this updated all the time?
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
func (b *Board) Print() {
	err := b.Validate()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", b.Squares)
}

// NewBoard a new board setup for standard play
func NewBoard() *Board {

	b := Board{
		Squares: getEmptySquares(),
	}

	for _, piece := range []Piece{
		{PieceType: Pawn{}, Player: White, Location: NewLocation("A2")},
		{PieceType: Pawn{}, Player: White, Location: NewLocation("B2")},
		{PieceType: Pawn{}, Player: White, Location: NewLocation("C2")},
		{PieceType: Pawn{}, Player: White, Location: NewLocation("D2")},
		{PieceType: Pawn{}, Player: White, Location: NewLocation("E2")},
		{PieceType: Pawn{}, Player: White, Location: NewLocation("F2")},
		{PieceType: Pawn{}, Player: White, Location: NewLocation("G2")},
		{PieceType: Pawn{}, Player: White, Location: NewLocation("H2")},

		{PieceType: Pawn{}, Player: Black, Location: NewLocation("A7")},
		{PieceType: Pawn{}, Player: Black, Location: NewLocation("B7")},
		{PieceType: Pawn{}, Player: Black, Location: NewLocation("C7")},
		{PieceType: Pawn{}, Player: Black, Location: NewLocation("D7")},
		{PieceType: Pawn{}, Player: Black, Location: NewLocation("E7")},
		{PieceType: Pawn{}, Player: Black, Location: NewLocation("F7")},
		{PieceType: Pawn{}, Player: Black, Location: NewLocation("G7")},
		{PieceType: Pawn{}, Player: Black, Location: NewLocation("H7")},

		{PieceType: Rook{}, Player: White, Location: NewLocation("A1")},
		{PieceType: Knight{}, Player: White, Location: NewLocation("B1")},
		{PieceType: Bishop{}, Player: White, Location: NewLocation("C1")},
		{PieceType: Queen{}, Player: White, Location: NewLocation("D1")},
		{PieceType: King{}, Player: White, Location: NewLocation("E1")},
		{PieceType: Bishop{}, Player: White, Location: NewLocation("F1")},
		{PieceType: Knight{}, Player: White, Location: NewLocation("G1")},
		{PieceType: Rook{}, Player: White, Location: NewLocation("H1")},

		{PieceType: Rook{}, Player: Black, Location: NewLocation("A8")},
		{PieceType: Knight{}, Player: Black, Location: NewLocation("B8")},
		{PieceType: Bishop{}, Player: Black, Location: NewLocation("C8")},
		{PieceType: Queen{}, Player: Black, Location: NewLocation("D8")},
		{PieceType: King{}, Player: Black, Location: NewLocation("E8")},
		{PieceType: Bishop{}, Player: Black, Location: NewLocation("F8")},
		{PieceType: Knight{}, Player: Black, Location: NewLocation("G8")},
		{PieceType: Rook{}, Player: Black, Location: NewLocation("H8")},
	} {
		b.AddPiece(piece)
	}

	return &b
}
