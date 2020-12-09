package chess

// PieceType abstraction for a single piece
type PieceType interface {
	Name() string
	SymbolMap() map[Player]rune
	IsValidMove(currentLocation, newLocation Location) bool
}

// Piece an actual piece
type Piece struct {
	PieceType
	Player
	Location
}

// IsValidMove can this piece be moved to new square
func (p Piece) IsValidMove(newLocation Location) bool {
	return p.PieceType.IsValidMove(p.Location, newLocation)
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
func (p Pawn) IsValidMove(currentLocation, newLocation Location) bool {
	return false
}

// IsValidMove for a rook
func (r Rook) IsValidMove(currentLocation, newLocation Location) bool {
	return false
}

// IsValidMove for a knight
func (k Knight) IsValidMove(currentLocation, newLocation Location) bool {
	return false
}

// IsValidMove for a bishop
func (b Bishop) IsValidMove(currentLocation, newLocation Location) bool {
	return false
}

// IsValidMove for a queen
func (q Queen) IsValidMove(currentLocation, newLocation Location) bool {
	return false
}

// IsValidMove for a king
func (k King) IsValidMove(currentLocation, newLocation Location) bool {
	return false
}
