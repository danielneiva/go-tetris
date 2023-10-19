package gotetris

import (
	"bytes"
	"image"

	"github.com/danielneiva/go-tetris/assets"

	"github.com/danielneiva/go-tetris/exceptions"
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	imageWidth  int
	imageHeight int
	offsetX     int
	offsetY     int
	image       *ebiten.Image
}

func (s *Sprite) LoadSprite(offsetX, offsetY int) {
	img, _, err := image.Decode(bytes.NewReader(assets.Tiles_png))
	exceptions.Handle(err)
	s.image = ebiten.NewImageFromImage(img)
	s.offsetX = offsetX
	s.offsetY = offsetY
}

func (s *Sprite) DrawSprite(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(s.image.SubImage(image.Rect(s.offsetX, s.offsetY, s.imageWidth+s.offsetX, s.imageHeight+s.offsetY)).(*ebiten.Image), op)
}
