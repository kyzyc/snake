package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"time"
)

type Game struct {
	Board     *Board
	Snake     *Snake
	IsEnd     bool
	FrameRate int
	Scores    int

	keyCh chan keyboard.Key
}

func (g *Game) InitGame() {
	width := 20
	height := 10
	g.Board = CreateBoard(width, height)
	g.Snake = CreateSnake(height/2, width/2)
	g.IsEnd = false
	g.FrameRate = 5
	g.Scores = 0

	g.keyCh = make(chan keyboard.Key, 3)
}

func (g *Game) KeyBoardHandle() {
	// 捕获按键
	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			close(g.keyCh)
			fmt.Println("Error reading key:", err)
			panic("Error reading key")
		}
		g.keyCh <- key
	}
}

func (g *Game) ReadKeyBoard() {
	select {
	case key := <-g.keyCh:
		g.Snake.ChangeDirection(key)
	default:
		return
	}
}

func (g *Game) UpdateLogic() {
	g.Snake.MoveSnake()
	if g.CheckWallCollision() {
		g.IsEnd = true
	}
}

func (g *Game) DrawFrame() {
	// 绘制顶部边界
	for y := 0; y < g.Board.Width+2; y++ {
		fmt.Printf("*")
	}
	fmt.Println()

	for x := 0; x < g.Board.Height; x++ {
		row := "*"
		for y := 0; y < g.Board.Width; y++ {
			if g.Snake.IsInside(x, y) {
				row += "+"
			} else {
				row += " "
			}
		}
		row += "*"
		fmt.Println(row)
	}

	for y := 0; y < g.Board.Width+2; y++ {
		fmt.Printf("*")
	}
	fmt.Println()
}

func (g *Game) CheckWallCollision() bool {
	head := g.Snake.Body[0]

	return head.X < 0 || head.X >= g.Board.Height || head.Y < 0 || head.Y >= g.Board.Width
}

func (g *Game) StartGame() {
	frameDuration := time.Second / time.Duration(g.FrameRate)
	ticker := time.NewTicker(frameDuration)

	defer ticker.Stop()

	go g.KeyBoardHandle()

	for !g.IsEnd {
		select {
		case <-ticker.C:
			fmt.Print("\033[H\033[2J")
			g.ReadKeyBoard()
			g.UpdateLogic()
			g.DrawFrame()
		}
	}

	fmt.Println("game end! score:", g.Scores)
}
