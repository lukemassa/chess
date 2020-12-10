package game

// Move encapsulation of a move
type Move struct {
	Piece       *Piece
	Destination Location
}

// IsValidMove can this piece be moved to new square
func (p Piece) IsValidMove(newLocation Location, b *Board) bool {
	targetPiece := b.Squares[newLocation.file][newLocation.rank]
	// Target is empty
	if targetPiece == nil {
		return p.PieceType.IsValidMove(p.Location, newLocation, p.Color)
	}

	// TODO: Make sure there are no pieces in the way

	// One of our pieces is already there
	if targetPiece.Color == p.Color {
		return false
	}
	return p.PieceType.CanCapture(p.Location, newLocation, p.Color)
}

// IsValidMove is this move valid on this board
func (m *Move) IsValidMove(b *Board) bool {
	return m.Piece.IsValidMove(m.Destination, b)
}
