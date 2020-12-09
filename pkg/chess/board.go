package chess

import "fmt"

// BoardSize number of squares along one edge of the board
const BoardSize = 8

// Board an abstraction for the current setup of the board
type Board struct {
	Pieces []Piece
}

// boardMap a description of what squares are filled how
type boardMap [][]*Piece

// Location on the board
type Location struct {

	// A-F
	file uint8
	// 1-8
	rank uint8
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
	fileAsInt := uint8(fileAsRune) - 'A'
	rankAsInt := uint8(rankAsRune) - '1'
	return Location{
		file: fileAsInt,
		rank: rankAsInt,
	}
}

// Potential todo: don't make this a map, have this updated all the time?
func (b *Board) getBoardMap() boardMap {
	//Start with an empty board
	ret := make([][]*Piece, BoardSize)
	for i := 0; i < BoardSize; i++ {
		ret[i] = make([]*Piece, BoardSize)
	}
	for i := 0; i < len(b.Pieces); i++ {
		piece := b.Pieces[i]
		fmt.Printf("Putting a piece at %dx%d", piece.Location.file, piece.Location.rank)
		ret[piece.Location.file][piece.Location.rank] = &piece
	}
	return ret
}

func (b boardMap) print() {
	for i := BoardSize - 1; i >= 0; i-- {
		fmt.Printf("%d", i+1)
		for j := 0; j < BoardSize; j++ {
			piece := b[j][i]
			if piece == nil {
				fmt.Print(" ")
			} else {
				fmt.Printf("%c", piece.Symbol())
			}
		}
		fmt.Println()
	}
	for i := 0; i < BoardSize; i++ {
		fmt.Printf("%c", rune(i)+'A')
	}
	fmt.Println()
}

// Print print the current board setup
func (b *Board) Print() {
	bm := b.getBoardMap()
	bm.print()
}

// NewBoard a new board setup for standard play
func NewBoard() *Board {

	return &Board{
		Pieces: []Piece{
			{PieceType: Pawn{}, Player: White, Location: NewLocation("C2")},
			{PieceType: Pawn{}, Player: Black, Location: NewLocation("C7")},
		},
	}
}
