package game

import (
	"fmt"
	"strings"
)

// BoardSize number of squares along one edge of the board
const BoardSize = 8

// Board an abstraction for the current setup of the board
type Board struct {
	Pieces []*Piece
	Squares
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

// IsValidMove can a given move be made
func (b *Board) IsValidMove(move Move) bool {
	return move.Piece.IsValidMove(move.Destination, b)
}

// MakeMove make a specific move
// returns true if this move won the game
func (b *Board) MakeMove(move *Move) bool {

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
	fmt.Printf("%s", b)
	// Remove the pointer from the old place, add the pointer at the new place
	b.Squares[move.Piece.Location.file][move.Piece.Location.rank] = nil

	fmt.Printf("Now its null %v\n", move.Piece.Location)
	fmt.Printf("%s", b)
	b.Squares[move.Destination.file][move.Destination.rank] = move.Piece

	// Update this piece's location
	move.Piece.Location = move.Destination

	// TODO: Implement check for end of game
	return false
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
		if b.Squares[l.file][l.rank] == nil {
			return fmt.Errorf("Piece not marked in square at %s", l)
		}
	}
	foundPieces := 0
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			piece := b.Squares[i][j]
			if piece == nil {
				continue
			}
			foundPieces++
			fmt.Printf("Found non-nil: %s\n", piece)
			foundPiece := false
			for k := 0; k < len(b.Pieces); k++ {
				if *b.Pieces[k] == *piece {
					foundPiece = true
					break
				}
			}
			if !foundPiece {
				return fmt.Errorf("There is a piece at %s that is not in piece list", piece.Location)
			}
		}
	}
	if foundPieces != len(b.Pieces) {
		return fmt.Errorf("Found %d in the board, but %d pieces in the set", foundPieces, len(b.Pieces))
	}
	return nil
}

// AddPiece add a piece onto the board
func (b *Board) AddPiece(piece *Piece) {
	b.Pieces = append(b.Pieces, piece)
	b.Squares[piece.file][piece.rank] = piece
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
	//err := b.Validate()
	//if err != nil {
	//log.Fatal(err)
	//}
	return fmt.Sprintf("%s", b.Squares)
}

// BlankBoard an empty board
func BlankBoard() *Board {

	return &Board{
		Squares: getEmptySquares(),
	}
}

// NewBoard a new board setup for standard play
func NewBoard() *Board {

	b := BlankBoard()

	pieces := []Piece{
		{PieceType: Pawn{}, Color: White, Location: NewLocation("A2")},
		{PieceType: Pawn{}, Color: White, Location: NewLocation("B2")},
		{PieceType: Pawn{}, Color: White, Location: NewLocation("C2")},
		{PieceType: Pawn{}, Color: White, Location: NewLocation("D2")},
		{PieceType: Pawn{}, Color: White, Location: NewLocation("E2")},
		{PieceType: Pawn{}, Color: White, Location: NewLocation("F2")},
		{PieceType: Pawn{}, Color: White, Location: NewLocation("G2")},
		{PieceType: Pawn{}, Color: White, Location: NewLocation("H2")},

		{PieceType: Pawn{}, Color: Black, Location: NewLocation("A7")},
		{PieceType: Pawn{}, Color: Black, Location: NewLocation("B7")},
		{PieceType: Pawn{}, Color: Black, Location: NewLocation("C7")},
		{PieceType: Pawn{}, Color: Black, Location: NewLocation("D7")},
		{PieceType: Pawn{}, Color: Black, Location: NewLocation("E7")},
		{PieceType: Pawn{}, Color: Black, Location: NewLocation("F7")},
		{PieceType: Pawn{}, Color: Black, Location: NewLocation("G7")},
		{PieceType: Pawn{}, Color: Black, Location: NewLocation("H7")},

		{PieceType: Rook{}, Color: White, Location: NewLocation("A1")},
		{PieceType: Knight{}, Color: White, Location: NewLocation("B1")},
		{PieceType: Bishop{}, Color: White, Location: NewLocation("C1")},
		{PieceType: Queen{}, Color: White, Location: NewLocation("D1")},
		{PieceType: King{}, Color: White, Location: NewLocation("E1")},
		{PieceType: Bishop{}, Color: White, Location: NewLocation("F1")},
		{PieceType: Knight{}, Color: White, Location: NewLocation("G1")},
		{PieceType: Rook{}, Color: White, Location: NewLocation("H1")},

		{PieceType: Rook{}, Color: Black, Location: NewLocation("A8")},
		{PieceType: Knight{}, Color: Black, Location: NewLocation("B8")},
		{PieceType: Bishop{}, Color: Black, Location: NewLocation("C8")},
		{PieceType: Queen{}, Color: Black, Location: NewLocation("D8")},
		{PieceType: King{}, Color: Black, Location: NewLocation("E8")},
		{PieceType: Bishop{}, Color: Black, Location: NewLocation("F8")},
		{PieceType: Knight{}, Color: Black, Location: NewLocation("G8")},
		{PieceType: Rook{}, Color: Black, Location: NewLocation("H8")},
	}
	for i := 0; i < len(pieces); i++ {
		b.AddPiece(&pieces[i])
	}

	return b
}
