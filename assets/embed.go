package assets

import (
	_ "embed"
)

var (
	//go:embed tiles.png
	var Tiles_png []byte

	//go:embed Poppins-Medium.ttf
	var Poppins_Medium_ttf []byte
)