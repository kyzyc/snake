package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"math/rand"
	"time"
)

type Game struct {
	Board  *Board
	Snake  *Snake
	Food   *Food
	IsEnd  bool
	FPS    int
	Scores int

	keyCh chan keyboard.Key
}

func (g *Game) InitGame() {
	width := 20
	height := 10
	g.Board = CreateBoard(width, height)
	g.Snake = CreateSnake(height/2, width/2)
	g.Food = CreateFood(2, 4)
	g.IsEnd = false
	g.FPS = 5
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

func randomInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // 创建新的随机数生成器并设置种子
	return r.Intn(max-min+1) + min
}

func (g *Game) UpdateLogic() {
	g.Snake.MoveSnake()
	if g.CheckWallCollision() || g.CheckSelfCollision() {
		g.IsEnd = true
	} else if g.CheckFoodCollision() {
		g.Scores += g.Food.Value
		newX := randomInt(0, g.Board.Height)
		newY := randomInt(0, g.Board.Width)

		for g.Snake.IsInside(newX, newY, g.Snake.Body) {
			newX = randomInt(0, g.Board.Height)
			newY = randomInt(0, g.Board.Width)
		}
		g.Food.Location.X = newX
		g.Food.Location.Y = newY
		return
	}
	g.Snake.Body = g.Snake.Body[:len(g.Snake.Body)-1]
}

func (g *Game) DrawFrame() {
	// 刷新界面
	fmt.Print("\033[H\033[2J")

	// 绘制顶部边界
	for y := 0; y < g.Board.Width+2; y++ {
		fmt.Printf("*")
	}
	fmt.Println()

	for x := 0; x < g.Board.Height; x++ {
		row := "*"
		for y := 0; y < g.Board.Width; y++ {
			if g.Snake.IsInside(x, y, g.Snake.Body) {
				row += "+"
			} else if g.Food.Location.X == x && g.Food.Location.Y == y {
				row += "."
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

func (g *Game) CheckFoodCollision() bool {
	head := g.Snake.Body[0]

	return g.Food.Location.X == head.X && g.Food.Location.Y == head.Y
}

func (g *Game) CheckSelfCollision() bool {
	head := g.Snake.Body[0]

	return g.Snake.IsInside(head.X, head.Y, g.Snake.Body[1:])
}

func (g *Game) StartGame() {

	frameDuration := time.Second / time.Duration(g.FPS)
	ticker := time.NewTicker(frameDuration)

	defer ticker.Stop()

	go g.KeyBoardHandle()

	for !g.IsEnd {
		select {
		case <-ticker.C:
			g.ReadKeyBoard()
			g.UpdateLogic()
			g.DrawFrame()
		}
	}

	fmt.Println("game end! score:", g.Scores)
}
