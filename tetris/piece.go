package gotetris

type Point struct {
	X int
	Y int
}

type Piece struct {
	body  [4][4]int
	point *Point
	game  *Game
	sprite Sprite
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

func (p *Piece) SetSprite() {
	p.sprite
}
