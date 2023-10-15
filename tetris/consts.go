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
	SCREEN_WIDTH  = 400
	SCREEN_HEIGHT = 1000
	TILE_SIDE     = 20
	BOARD_WIDTH   = 20
	BOARD_HEIGHT  = 40
	PADDING       = 200
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
	},
}
