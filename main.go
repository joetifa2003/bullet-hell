package main

import (
	"github.com/joetifa2003/bullet-hell/pkg/game"
)

const DEBUG_COLLISON = false

func main() {
	g := game.NewGame()
	g.Loop()
}
