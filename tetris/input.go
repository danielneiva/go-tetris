package gotetris

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)

	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
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
		if repeatingKeyPressed(key) {
			if !g.piece.WillTouch(direction) {
				g.piece.Move(direction)
			}
		}
	}

	if repeatingKeyPressed(ebiten.KeyUp) {
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
