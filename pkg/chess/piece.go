package chess

// PieceType abstraction for a single piece
type PieceType interface {
	Name() string
	SymbolMap() map[Player]rune
}

// Piece an actual piece
type Piece struct {
	PieceType
	Player
	Location
}

// Symbol rune for this piece
func (p Piece) Symbol() rune {
	return p.PieceType.SymbolMap()[p.Player]
}

// Pawn a pawn
type Pawn struct {
}
