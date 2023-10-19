package gotetris

var (
	err error
)

const (
	Left  = 0
	Up    = 1
	Right = 2
	Down  = 3
)

const (
	BOARD_WIDTH   = 10
	BOARD_HEIGHT  = 20
	TILE_SIDE     = 40
	PADDING       = (TILE_SIDE * BOARD_HEIGHT) / 5
	SCREEN_WIDTH  = TILE_SIDE * BOARD_WIDTH
	SCREEN_HEIGHT = TILE_SIDE*BOARD_HEIGHT + PADDING
	TILE_POINT    = 100
	TILES_PATH    = "assets/tiles.png"
)

var BASE_PIECES = map[string]Piece{
	"T": {
		[4][4]int{
			{0, 1, 0, 0},
			{0, 1, 1, 0},
			{0, 1, 0, 0},
			{0, 0, 0, 0},
		},
		&Point{BOARD_WIDTH / 2, 0},
		&Game{},
		&Sprite{},
	},
	"L": {
		[4][4]int{
			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 0},
		},
		&Point{BOARD_WIDTH / 2, 0},
		&Game{},
		&Sprite{},
	},
	"J": {
		[4][4]int{
			{0, 0, 1, 0},
			{0, 0, 1, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 0},
		},
		&Point{BOARD_WIDTH / 2, 0},
		&Game{},
		&Sprite{},
	},
	"_": {
		[4][4]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{1, 1, 1, 1},
			{0, 0, 0, 0},
		},
		&Point{BOARD_WIDTH / 2, 0},
		&Game{},
		&Sprite{},
	},
	"O": {
		[4][4]int{
			{0, 0, 0, 0},
			{0, 1, 1, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 0},
		},
		&Point{BOARD_WIDTH / 2, 0},
		&Game{},
		&Sprite{},
	},
	"Z": {
		[4][4]int{
			{0, 0, 0, 0},
			{1, 1, 0, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 0},
		},
		&Point{BOARD_WIDTH / 2, 0},
		&Game{},
		&Sprite{},
	},
	"S": {
		[4][4]int{
			{0, 0, 0, 0},
			{0, 0, 1, 1},
			{0, 1, 1, 0},
			{0, 0, 0, 0},
		},
		&Point{BOARD_WIDTH / 2, 0},
		&Game{},
		&Sprite{},
	},
}
