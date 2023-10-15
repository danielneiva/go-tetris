package main

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"log"
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

type Point struct {
	X int
	Y int
}

type Piece struct {
	body  [4][4]int
	point *Point
	game  *Game
}

func (p *Piece) CanRotate() bool {
	rtmp := *p
	rtmp.Rotate()
	return !rtmp.IsTouching()
}

func (p *Piece) IsTouching() bool {
	if p.point.Y < 0 || p.point.X < 0 {
		return true
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if p.game.board[i+p.point.X][j+p.point.Y] != 0 && p.body[i][j] != 0 {
				return true
			}
		}
	}
	return false
}

func (p *Piece) Rotate() {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			tmp := p.body[3-j][i]
			p.body[3-j][i] = p.body[3-i][3-j]
			p.body[3-i][3-j] = p.body[j][3-i]
			p.body[j][3-i] = p.body[i][j]
			p.body[i][j] = tmp
		}
	}
}
func (p *Piece) WillTouch(direction int) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			isTileFilled := p.body[i][j] != 0
			switch direction {
			case Left:
				if isTileFilled && p.game.board[i+p.point.X-1][j+p.point.Y] != 0 {
					return true
				}
			case Right:
				if isTileFilled && p.game.board[i+p.point.X+1][j+p.point.Y] != 0 {
					return true
				}
			case Down:
				if isTileFilled && p.game.board[i+p.point.X][j+p.point.Y+1] != 0 {
					return true
				}
			}
		}
	}
	return false
}

func (p *Piece) Move(direction int) {
	// fmt.Println("Moving...")
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
	// fmt.Println("Descending...")
	if !p.WillTouch(Down) {
		p.point.Y = p.point.Y + 1
	}
}

func (p *Piece) AddToGameBoard() {
	// fmt.Println("Adding to game board...")
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if p.body[i][j] != 0 {
				p.game.board[i+p.point.X][j+p.point.Y] = 1
			}
		}
	}
	p.game.CheckForCompleteLines()
	p.game.GetNewPiece()
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
	font          font.Face
}

func (g *Game) NewGame() {
	// fmt.Println("Starting game...")
	g.LoadSprites()
	g.LoadFonts()
	g.InitBoard()
	g.GetNewPiece()
	g.score = 0
	g.ticks = 25
	g.speed = 25
	g.gameOver = false
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

func (g *Game) LoadFonts() {

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	g.font, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
}

func (g *Game) ClearRow(y int) {
	for i := 1; i < BOARD_WIDTH-1; i++ {
		g.board[i][y] = 0
		g.score += TILE_POINT
	}
}

func (g *Game) DescendRows(origin int) {
	rowIsEmpty := true
	var row []int
	for i := 1; i < BOARD_WIDTH-1; i++ {
		row = append(row, g.board[i][origin])
		if g.board[i][origin] != 0 {
			rowIsEmpty = false
		}
	}
	if rowIsEmpty {
		return
	}
	for i := 1; i < BOARD_WIDTH-1; i++ {
		g.board[i][origin+1] = row[i-1]
		g.board[i][origin] = 0
	}
	g.DescendRows(origin - 1)
}

func (g *Game) CheckForCompleteLines() {
	// fmt.Println("Checking for completed lines...")
	for j := BOARD_HEIGHT - 2; j > 0; j-- {
		rowIsComplete := true
		for i := 1; i < BOARD_WIDTH-1; i++ {
			if g.board[i][j] == 0 {
				rowIsComplete = false
				break
			}
		}
		if rowIsComplete {
			g.ClearRow(j)
			g.DescendRows(j - 1)
			j++
		}
	}
}

func (g *Game) PrintBoard() {
	for i := 0; i < BOARD_WIDTH; i++ {
		for j := 0; j < BOARD_HEIGHT; j++ {
			fmt.Printf("%d ", g.board[i][j])
		}
		fmt.Println()
	}
	fmt.Println("==============================================================================================================================")
}

func (g *Game) Restart() {
	g.NewGame()
}

func (g *Game) DrawScore(screen *ebiten.Image) {
	msg := fmt.Sprintf("SCORE: %d", g.score)
	text.Draw(screen, msg, g.font, 0, 100, color.Black)
}

func (g *Game) CheckFirstLine() {
	// fmt.Println("Checking first line...")
	for i := 1; i < BOARD_WIDTH-1; i++ {
		if g.board[i][1] != 0 {
			g.gameOver = true
			return
		}
	}
}

func (g *Game) GetNewPiece() {
	// fmt.Println("Getting new Piece...")
	g.CheckFirstLine()
	index := rand.Intn(len(BASE_PIECES))
	newPiece := &Piece{}
	for _, *newPiece = range BASE_PIECES {
		if index == 0 {
			g.piece = *newPiece
		}
		index--
	}
	g.piece.point = &Point{BOARD_WIDTH / 2, 0}
	g.piece.game = g
}

func (g *Game) LoadSprites() {
	// fmt.Println("Loading Sprites...")
	g.sprite, _, err = ebitenutil.NewImageFromFile(TILES_PATH)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	if g.gameOver {
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			g.Restart()
		}
		return nil
	}

	if repeatingKeyPressed(ebiten.KeyLeft) {
		if !g.piece.WillTouch(Left) {
			g.piece.Move(Left)
		}
	}
	if repeatingKeyPressed(ebiten.KeyUp) {
		if g.piece.CanRotate() {
			g.piece.Rotate()
		}
	}
	if repeatingKeyPressed(ebiten.KeyRight) {
		if !g.piece.WillTouch(Right) {
			g.piece.Move(Right)
		}
	}
	if repeatingKeyPressed(ebiten.KeyDown) {
		if !g.piece.WillTouch(Down) {
			g.piece.Move(Down)
		}
	}

	g.updateCounter++
	//Control speed
	if g.updateCounter < g.speed {
		return nil
	}
	g.updateCounter = 0
	if g.piece.WillTouch(Down) {
		g.piece.AddToGameBoard()
		return nil
	}
	g.piece.Descend()
	return nil
}

func (g *Game) drawRec(screen *ebiten.Image, x, y, w, h int, clr color.Color, val int) {
	vector.DrawFilledRect(screen, float32(x*TILE_SIDE), float32(y*TILE_SIDE+PADDING), float32(w), float32(h), clr, true)
	vector.StrokeRect(screen, float32(x*TILE_SIDE), float32(y*TILE_SIDE+PADDING), float32(w), float32(h), float32(2), color.RGBA{0, 0, 0, 255}, true)
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	for i := 0; i < BOARD_WIDTH; i++ {
		for j := 0; j < BOARD_HEIGHT; j++ {
			switch g.board[i][j] {
			case 1:
				g.drawRec(screen, i, j, TILE_SIDE, TILE_SIDE, color.RGBA{0, 255, 0, 255}, g.board[i][j])
			case 2:
				g.drawRec(screen, i, j, TILE_SIDE, TILE_SIDE, color.RGBA{255, 0, 0, 255}, g.board[i][j])
			}
		}
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if g.piece.body[i][j] != 0 {
				g.drawRec(screen, g.piece.point.X+i, g.piece.point.Y+j, TILE_SIDE, TILE_SIDE, color.RGBA{0, 0, 255, 255}, g.piece.body[i][j])
			}
		}
	}

	g.DrawScore(screen)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	//screen.DrawImage(g.sprite, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func main() {
	game := &Game{}
	game.NewGame()
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
