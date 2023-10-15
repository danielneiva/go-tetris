package gotetris

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func isKeyJustPressedOrBeingHeld(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)

	pressDuration := inpututil.KeyPressDuration(key)
	if pressDuration == 1 {
		return true
	}
	if pressDuration >= delay && (pressDuration-delay)%interval == 0 {
		return true
	}
	return false
}

func (g *Game) ExecCommand(key ebiten.Key) {
	movementKeys := map[ebiten.Key]int{
		ebiten.KeyLeft:  Left,
		ebiten.KeyRight: Right,
		ebiten.KeyDown:  Down,
	}
	for key, direction := range movementKeys {
		if isKeyJustPressedOrBeingHeld(key) {
			if !g.piece.WillTouch(direction) {
				g.piece.Move(direction)
			}
		}
	}

	if isKeyJustPressedOrBeingHeld(ebiten.KeyUp) {
		if g.piece.CanRotate() {
			g.piece.Rotate()
		}
	}
}

func (g *Game) HandleInput() {

	var keys []ebiten.Key
	keys = inpututil.AppendPressedKeys(keys)

	for _, key := range keys {
		g.ExecCommand(key)
	}

}
