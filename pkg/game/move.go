package game

// Move encapsulation of a move
type Move struct {
	Piece       *Piece
	Destination Location
}

// IsValidMove can this piece be moved to new square
func (p Piece) IsValidMove(newLocation Location, b *Board) bool {

	// Cannot move if there are intervening pieces
	if p.Name() != "knight" && b.PieceBetween(p.Location, newLocation) {
		return false
	}

	targetPiece := b.Squares[newLocation.file][newLocation.rank]

	// Target square is empty
	if targetPiece == nil {
		return p.PieceType.IsValidMove(p.Location, newLocation, p.Color)
	}

	// It's occupied but by one of our pieces
	if targetPiece.Color == p.Color {
		return false
	}
	// It's occupied by an opponent piece, can it be taken?
	return p.PieceType.CanCapture(p.Location, newLocation, p.Color)
}

// PieceBetween found a piece between two locations
func (b *Board) PieceBetween(location1, location2 Location) bool {

	// Find if they are on the same rank/file/diagonal
	// if they are, walk between them and check each square
	var filedelta int8
	var rankdelta int8

	// Same file
	if location1.file == location2.file {
		if location1.rank < location2.rank {
			rankdelta = 1
		} else {
			rankdelta = -1
		}
		return b.pieceBetween(location1, location2, filedelta, rankdelta)
	}
	// Same rank
	if location1.rank == location2.rank {
		if location1.file < location2.file {
			filedelta = 1
		} else {
			filedelta = -1
		}
		return b.pieceBetween(location1, location2, filedelta, rankdelta)
	}
	// Diagonal upward
	if location1.rank-location2.rank == location1.file-location2.file {
		if location1.file < location2.file {
			filedelta = 1
			rankdelta = 1
		} else {
			filedelta = -1
			rankdelta = -1
		}
		return b.pieceBetween(location1, location2, filedelta, rankdelta)
	}
	// Diagonal downward
	if location1.rank-location2.rank == location2.file-location1.file {
		if location1.file < location2.file {
			filedelta = 1
			rankdelta = -1
		} else {
			filedelta = -1
			rankdelta = 1
		}
		return b.pieceBetween(location1, location2, filedelta, rankdelta)
	}
	return false
}

// Walks between location1 and location2 by the deltas
// No bounds checking is done, so make sure you know the deltas
// will get you to the other side
func (b *Board) pieceBetween(location1, location2 Location, filedelta, rankdelta int8) bool {

	file := location1.file + filedelta
	rank := location1.rank + rankdelta
	for rank != location2.rank || file != location2.file {

		if b.Squares[file][rank] != nil {
			return true
		}
		file += filedelta
		rank += rankdelta
	}
	return false
}

// IsValidMove is this move valid on this board
func (m *Move) IsValidMove(b *Board) bool {
	return m.Piece.IsValidMove(m.Destination, b)
}
