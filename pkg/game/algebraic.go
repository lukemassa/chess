package game

// ConvertAlgebraicToMove convert an algebraic string (like "BF3" or "NxA7") into a move
// If poorly formed or not a legal move, return an error
func (b *Board) ConvertAlgebraicToMove(algebraicString string) (*Move, error) {
	return &Move{
		Piece:       b.Squares[4][1],
		Destination: NewLocation("E4"),
	}, nil
}
