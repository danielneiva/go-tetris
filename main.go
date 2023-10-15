package main

import (
	"github.com/danielneiva/tp1-es2/game"
	"github.com/hajimehoshi/ebiten"
	"log"
)

func main() {
	game := &Game{}
	game.NewGame()
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
