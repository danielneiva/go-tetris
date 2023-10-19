package gotetris

import (
	"assets"
	"image"
	"bytes"
	"exceptions"
)

type Sprite struct {
	imageWidth int
	imageHeight int
	image *ebiten.Image
}


func (s *Sprite) LoadSprite(offsetX, offsetY int, imageName string) {
	img, _, err = image.Decode(bytes.NewReader(assets.Tiles_png))
	exceptions.handle(err)
	s.image = ebiten.NewImageFromImage(img)
}