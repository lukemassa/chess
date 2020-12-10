package game

// Boring file that just lists the name and symbol of all the pieces

// Symbol rune for this piece
func (p Piece) Symbol() rune {
	if p.Color == White {
		return p.WhiteSymbol()
	}
	return p.BlackSymbol()
}

////////////////////////////////////////////////////////////////////////////////////
// Name of pawn
func (p Pawn) Name() string {
	return "Pawn"
}

// WhiteSymbol for pawn
func (p Pawn) WhiteSymbol() rune {
	return '♙'
}

// BlackSymbol for pawn
func (p Pawn) BlackSymbol() rune {
	return '♟'
}

////////////////////////////////////////////////////////////////////////////////////
// Name of rook
func (r Rook) Name() string {
	return "Rook"
}

// WhiteSymbol for rook
func (r Rook) WhiteSymbol() rune {
	return '♖'
}

// BlackSymbol for rook
func (r Rook) BlackSymbol() rune {
	return '♜'
}

////////////////////////////////////////////////////////////////////////////////////
// Name of knight
func (k Knight) Name() string {
	return "Knight"
}

// WhiteSymbol for knight
func (k Knight) WhiteSymbol() rune {
	return '♘'
}

// BlackSymbol for knight
func (k Knight) BlackSymbol() rune {
	return '♞'
}

////////////////////////////////////////////////////////////////////////////////////
// Name of Bishop
func (b Bishop) Name() string {
	return "Bishop"
}

// WhiteSymbol for knight
func (b Bishop) WhiteSymbol() rune {
	return '♗'
}

// BlackSymbol for knight
func (b Bishop) BlackSymbol() rune {
	return '♝'
}

////////////////////////////////////////////////////////////////////////////////////
// Name of queen
func (q Queen) Name() string {
	return "Queen"
}

// WhiteSymbol for queen
func (q Queen) WhiteSymbol() rune {
	return '♕'
}

// BlackSymbol for queen
func (q Queen) BlackSymbol() rune {
	return '♛'
}

////////////////////////////////////////////////////////////////////////////////////
// Name of king
func (k King) Name() string {
	return "King"
}

// WhiteSymbol for knight
func (k King) WhiteSymbol() rune {
	return '♔'
}

// BlackSymbol for knight
func (k King) BlackSymbol() rune {
	return '♚'
}
