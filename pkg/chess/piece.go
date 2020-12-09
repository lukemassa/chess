package chess

// PieceType abstraction for a single piece
type PieceType interface {
	Name() string
	SymbolMap() map[Player]rune
	IsValidMove(currentLocation, newLocation Location, player Player) bool
}

// Piece an actual piece
type Piece struct {
	PieceType
	Player
	Location
}

// IsValidMove can this piece be moved to new square
func (p Piece) IsValidMove(newLocation Location) bool {
	// TODO: Is one of my pieces there/etc
	return p.PieceType.IsValidMove(p.Location, newLocation, p.Player)
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
func (p Pawn) IsValidMove(currentLocation, newLocation Location, player Player) bool {
	// Has to go in a straight line
	if currentLocation.file != newLocation.file {
		return false
	}
	// Switch to "white's" perspective
	currentRank := currentLocation.rank
	newRank := newLocation.rank
	if player == Black {
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
func (r Rook) IsValidMove(currentLocation, newLocation Location, player Player) bool {
	return onSameRankOrFile(currentLocation, newLocation)
}

// IsValidMove for a knight
func (k Knight) IsValidMove(currentLocation, newLocation Location, player Player) bool {
	rankDiff := abs(currentLocation.rank - newLocation.rank)
	fileDiff := abs(currentLocation.file - newLocation.file)
	//fmt.Printf("Going from %s to %s: rankdiff is %d, filediff is %d", currentLocation, newLocation, rankDiff, fileDiff)
	if (rankDiff == 2 && fileDiff == 1) || (rankDiff == 1 && fileDiff == 2) {
		return true
	}
	return false
}

// IsValidMove for a bishop
func (b Bishop) IsValidMove(currentLocation, newLocation Location, player Player) bool {
	return onSameDiagonal(currentLocation, newLocation)
}

// IsValidMove for a queen
func (q Queen) IsValidMove(currentLocation, newLocation Location, player Player) bool {
	return onSameRankOrFile(currentLocation, newLocation) || onSameDiagonal(currentLocation, newLocation)
}

// IsValidMove for a king
func (k King) IsValidMove(currentLocation, newLocation Location, player Player) bool {
	rankDiff := abs(currentLocation.rank - newLocation.rank)
	fileDiff := abs(currentLocation.file - currentLocation.file)
	if rankDiff <= 1 && fileDiff <= 1 {
		return true
	}
	return false
}
