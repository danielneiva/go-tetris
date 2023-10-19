package assets

import (
	_ "embed"
)

var (
	//go:embed tiles.png
	Tiles_png []byte

	//go:embed Poppins-Medium.ttf
	Poppins_Medium_ttf []byte
)
