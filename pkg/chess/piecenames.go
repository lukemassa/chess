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
