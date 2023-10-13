package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"

	// "image/color"
	"log"
)

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
	SCREEN_WIDTH  = 640
	SCREEN_HEIGHT = 480
	TILE_SIDE     = 20
	BOARD_WIDTH   = 20
	BOARD_HEIGHT  = 40
	PADDING       = 200
	TILES_PATH    = "assets/tiles.png"
)

var BASE_PIECES = map[string]Piece{
	"T": {
		[4][4]int{
			{0, 1, 0, 0},
			{1, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 0, 0},
		},
		&Point{10, 0},
	},
	"L": {
		[4][4]int{
			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 0},
		},
		&Point{10, 0},
	},
	"J": {
		[4][4]int{
			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{1, 1, 0, 0},
			{0, 0, 0, 0},
		},
		&Point{10, 0},
	},
	"_": {
		[4][4]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{1, 1, 1, 1},
		},
		&Point{10, 0},
	},
	"O": {
		[4][4]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 1, 1, 0},
			{0, 1, 1, 0},
		},
		&Point{10, 0},
	},
	"Z": {
		[4][4]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{1, 1, 0, 0},
			{0, 1, 1, 0},
		},
		&Point{10, 0},
	},
	"S": {
		[4][4]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 1, 1, 0},
			{1, 1, 0, 0},
		},
		&Point{10, 0},
	},
}

type Point struct {
	X int
	Y int
}

type Piece struct {
	body  [4][4]int
	point *Point
}

func (p *Piece) SetFigure(basePiece *Piece) {

}
func (p *Piece) PrintFigure() {

}
func (p *Piece) CanRotate() bool {
	return true
}
func (p *Piece) Rotate() {

}
func (p *Piece) WontTouch(direction int) bool {
	return true
}
func (p *Piece) Move(direction int) {
	switch direction {
	case Left:
		p.point.X -= 1
	case Right:
		p.point.X += 1
	case Down:
		p.point.Y += 1
	}
}
func (p *Piece) Descend() {
	p.point.Y = p.point.Y + 1
}

type Game struct {
	board         [BOARD_WIDTH][BOARD_HEIGHT]int
	piece         Piece
	score         int
	ticks         int
	speed         int
	gameOver      bool
	sprite        *ebiten.Image
	updateCounter int
}

func NewGame() *Game {
	var g = &Game{}
	g.LoadSprites()
	g.InitBoard()
	g.GetNewPiece()
	g.score = 0
	g.ticks = 0
	g.speed = 25
	g.gameOver = false
	return g
}

func (g *Game) InitBoard() {
	for i := 0; i < BOARD_WIDTH; i++ {
		for j := 0; j < BOARD_HEIGHT; j++ {
			if i == 0 || i == BOARD_WIDTH-1 || j == BOARD_HEIGHT-1 {
				g.board[i][j] = 2 //field limits
			} else {
				g.board[i][j] = 0
			}
		}
	}
}

func (g *Game) restart() {
	g = NewGame()
}

func (g *Game) GetNewPiece() {
	index := rand.Intn(len(BASE_PIECES))
	for _, newPiece := range BASE_PIECES {
		if index == 0 {
			g.piece = newPiece
		}
		index--
	}
	g.piece.point = &Point{10, 0}
}

func (g *Game) LoadSprites() {
	g.sprite, _, err = ebitenutil.NewImageFromFile(TILES_PATH)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	if g.gameOver {
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			g.restart()
		}
		return nil
	}
	g.updateCounter++
	//Control speed
	if g.updateCounter < g.speed {
		return nil
	}
	g.updateCounter = 0
	g.piece.Descend()

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.piece.Move(Left)
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.piece.Rotate()
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.piece.Move(Right)
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.piece.Move(Down)
	}

	return nil
}

func drawRec(screen *ebiten.Image, x, y, w, h int, clr color.Color) {
	vector.DrawFilledRect(screen, float32(x*TILE_SIDE+PADDING), float32(y*TILE_SIDE+PADDING), float32(w), float32(h), clr, true)
	vector.StrokeRect(screen, float32(x*TILE_SIDE+PADDING), float32(y*TILE_SIDE+PADDING), float32(w), float32(h), float32(2), color.RGBA{0, 0, 0, 255}, true)
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	for i := 0; i < BOARD_WIDTH; i++ {
		for j := 0; j < BOARD_HEIGHT; j++ {
			if g.board[i][j] != 0 {
				drawRec(screen, i, j, TILE_SIDE, TILE_SIDE, color.RGBA{255, 0, 0, 255})
			}
		}
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if g.piece.body[i][j] != 0 {
				drawRec(screen, g.piece.point.X+i, g.piece.point.Y+j, TILE_SIDE, TILE_SIDE, color.RGBA{0, 0, 255, 255})
			}
		}
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	//screen.DrawImage(g.sprite, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 720, 1280
}

func main() {
	game := NewGame()
	ebiten.SetWindowSize(SCREEN_WIDTH*2, SCREEN_HEIGHT*2)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
