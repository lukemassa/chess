package game

// Player player of chess
type Player interface {
	NextMove(b *Board) Move
}

// Play the game
func (g *Game) Play() {
	for {
		player := g.Turn()
		g.whitesTurn = !g.whitesTurn
		move := player.NextMove(g.Board)
		g.Board.MakeMove(move)
		break
	}
}
