package chess

// Boring class that just lists the name and symbol of all the pieces

// Name of pawn
func (p Pawn) Name() string {
	return "Pawn"
}

// SymbolMap for pawn
func (p Pawn) SymbolMap() map[Player]rune {
	return map[Player]rune{
		Black: '♟',
		White: '♙',
	}
}

// Name of rook
func (r Rook) Name() string {
	return "Rook"
}

// SymbolMap for rook
func (r Rook) SymbolMap() map[Player]rune {
	return map[Player]rune{
		Black: '♜',
		White: '♖',
	}
}

// Name of knight
func (k Knight) Name() string {
	return "Knight"
}

// SymbolMap for knight
func (k Knight) SymbolMap() map[Player]rune {
	return map[Player]rune{
		Black: '♞',
		White: '♘',
	}
}

// Name of queen
func (q Queen) Name() string {
	return "Queen"
}

// SymbolMap for knight
func (q Queen) SymbolMap() map[Player]rune {
	return map[Player]rune{
		Black: '♛',
		White: '♕',
	}
}

// Name of king
func (k King) Name() string {
	return "King"
}

// SymbolMap for knight
func (k King) SymbolMap() map[Player]rune {
	return map[Player]rune{
		Black: '♚',
		White: '♔',
	}
}

// Name of Bishop
func (b Bishop) Name() string {
	return "Bishop"
}

// SymbolMap for knight
func (b Bishop) SymbolMap() map[Player]rune {
	return map[Player]rune{
		Black: '♝',
		White: '♗',
	}
}
