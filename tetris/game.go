package gotetris

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

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
	g.HandleInput()

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

func (g *Game) drawRec(screen *ebiten.Image, x, y, w, h int, bgColor color.Color, strkColor color.Color, val int) {
	vector.DrawFilledRect(screen, float32(x*TILE_SIDE), float32(y*TILE_SIDE+PADDING), float32(w), float32(h), bgColor, true)
	vector.StrokeRect(screen, float32(x*TILE_SIDE), float32(y*TILE_SIDE+PADDING), float32(w), float32(h), float32(2), strkColor, true)
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	for i := 0; i < BOARD_WIDTH; i++ {
		for j := 0; j < BOARD_HEIGHT; j++ {
			var recColor color.RGBA
			strkColor := color.RGBA{0, 0, 0, 255}
			switch g.board[i][j] {
			case 1:
				recColor = color.RGBA{0, 255, 0, 255}
			case 2:
				recColor = color.RGBA{255, 0, 0, 255}
			case 0:
				strkColor = color.RGBA{0, 0, 0, 0}
			}
			g.drawRec(screen, i, j, TILE_SIDE, TILE_SIDE, recColor, strkColor, g.board[i][j])
		}
	}

	blueColor := color.RGBA{0, 0, 255, 255}
	blackColor := color.RGBA{0, 0, 0, 255}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if g.piece.body[i][j] != 0 {
				g.drawRec(screen, g.piece.point.X+i, g.piece.point.Y+j, TILE_SIDE, TILE_SIDE, blueColor, blackColor, g.piece.body[i][j])
			}
		}
	}

	g.piece.sprite.DrawSprite(screen)

	g.DrawScore(screen)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	//screen.DrawImage(g.sprite, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}
