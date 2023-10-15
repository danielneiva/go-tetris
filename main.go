package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"tp1-es2/tetris"
)

func main() {
	game := &gotetris.Game{}
	game.NewGame()
	ebiten.SetWindowSize(gotetris.SCREEN_WIDTH, gotetris.SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
