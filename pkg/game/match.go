package game

// Player player of chess
type Player interface {
	MakeMove(b *Board) Move
}
