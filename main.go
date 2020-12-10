package main

import (
	"github.com/lukemassa/chess/pkg/game"
	"github.com/lukemassa/chess/pkg/players"
)

func main() {
	g := game.New(&players.InteractivePlayer{
		Notation: game.CoordinateNotation{},
	}, &players.InteractivePlayer{
		Notation: game.CoordinateNotation{},
	}, true)
	g.Play()
}
