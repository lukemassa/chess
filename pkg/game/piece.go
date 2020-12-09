package game

// PieceType abstraction for a single piece
type PieceType interface {
	Name() string
	SymbolMap() map[Color]rune
	IsValidMove(currentLocation, newLocation Location, color Color) bool
	CanCapture(currentLocation, opponentLocation Location, color Color) bool
}

// Piece an actual piece
type Piece struct {
	PieceType
	Color
	Location
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

func abs(x int8) int8 {
	if x < 0 {
		return -x
	}
	return x
}

func onSameRankOrFile(locationA, locationB Location) bool {
	return locationA.rank == locationB.rank || locationA.file == locationB.file
}

func onSameDiagonal(locationA, locationB Location) bool {
	return locationA.rank-locationB.rank == locationA.file-locationB.file ||
		locationA.rank-locationB.rank == locationB.file-locationA.file
}

// Pawn a pawn
type Pawn struct {
}

// Rook a rook
type Rook struct {
}

// Bishop a bishop
type Bishop struct {
}

// Knight a rook
type Knight struct {
}

// Queen a rook
type Queen struct {
}

// King a rook
type King struct {
}

// IsValidMove for a pawn
func (p Pawn) IsValidMove(currentLocation, newLocation Location, color Color) bool {
	// Has to go in a straight line
	if currentLocation.file != newLocation.file {
		return false
	}
	// Switch to "white's" perspective
	currentRank := currentLocation.rank
	newRank := newLocation.rank
	if color == Black {
		currentRank = (BoardSize - 1) - currentRank
		newRank = (BoardSize - 1) - newRank
	}

	// Has to go forward
	if currentRank > newRank {
		return false
	}
	// Can never move more than 2
	if newRank-currentRank > 2 {
		return false
	}
	// Can only move two if starting at square '1'
	if newRank-currentRank == 2 && currentRank != 1 {
		return false
	}
	return true
}

// IsValidMove for a rook
func (r Rook) IsValidMove(currentLocation, newLocation Location, color Color) bool {
	return onSameRankOrFile(currentLocation, newLocation)
}

// IsValidMove for a knight
func (k Knight) IsValidMove(currentLocation, newLocation Location, color Color) bool {
	rankDiff := abs(currentLocation.rank - newLocation.rank)
	fileDiff := abs(currentLocation.file - newLocation.file)
	if (rankDiff == 2 && fileDiff == 1) || (rankDiff == 1 && fileDiff == 2) {
		return true
	}
	return false
}

// IsValidMove for a bishop
func (b Bishop) IsValidMove(currentLocation, newLocation Location, color Color) bool {
	return onSameDiagonal(currentLocation, newLocation)
}

// IsValidMove for a queen
func (q Queen) IsValidMove(currentLocation, newLocation Location, color Color) bool {
	return onSameRankOrFile(currentLocation, newLocation) || onSameDiagonal(currentLocation, newLocation)
}

// IsValidMove for a king
func (k King) IsValidMove(currentLocation, newLocation Location, color Color) bool {
	rankDiff := abs(currentLocation.rank - newLocation.rank)
	fileDiff := abs(currentLocation.file - newLocation.file)
	if rankDiff <= 1 && fileDiff <= 1 {
		return true
	}
	return false
}

// CanCapture from a pawn's perspective
func (p Pawn) CanCapture(currentLocation, opponentLocation Location, color Color) bool {
	fileDiff := abs(currentLocation.file - opponentLocation.file)

	// Must be one file to the left or right
	if fileDiff != 1 {
		return false
	}

	// Switch to "white's" perspective
	currentRank := currentLocation.rank
	opponentRank := opponentLocation.rank
	if color == Black {
		currentRank = (BoardSize - 1) - currentRank
		opponentRank = (BoardSize - 1) - opponentRank
	}
	// Must be the very next rank
	if currentRank+1 != opponentRank {
		return false
	}
	return true
}

// CanCapture from a rook's perspective
func (r Rook) CanCapture(currentLocation, opponentLocation Location, color Color) bool {
	return r.IsValidMove(currentLocation, opponentLocation, color)
}

// CanCapture from a knight's perspective
func (k Knight) CanCapture(currentLocation, opponentLocation Location, color Color) bool {
	return k.IsValidMove(currentLocation, opponentLocation, color)
}

// CanCapture from a bishop's perspective
func (b Bishop) CanCapture(currentLocation, opponentLocation Location, color Color) bool {
	return b.IsValidMove(currentLocation, opponentLocation, color)
}

// CanCapture from a queen's perspective
func (q Queen) CanCapture(currentLocation, opponentLocation Location, color Color) bool {
	return q.IsValidMove(currentLocation, opponentLocation, color)
}

// CanCapture from a king's perspective
func (k King) CanCapture(currentLocation, opponentLocation Location, color Color) bool {
	return k.IsValidMove(currentLocation, opponentLocation, color)
}
