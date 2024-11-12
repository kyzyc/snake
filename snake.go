package main

import (
	"github.com/eiannone/keyboard"
)

type Point struct {
	X int
	Y int
}

type Direction int

const (
	LEFT Direction = iota
	RIGHT
	UP
	DOWN
)

type Snake struct {
	Body      []Point   // points of snake body, Body[0] is the head
	Length    int       // the length of snake
	Direction Direction // the snake move direction
}

func CreateSnake(startX, startY int) *Snake {
	body := make([]Point, 1)
	body[0] = Point{X: startX, Y: startY}

	snake := &Snake{
		Body:      body,
		Direction: RIGHT,
	}

	return snake
}

func (snake *Snake) IsInside(x, y int, body []Point) bool {
	for _, p := range body {
		if p.X == x && p.Y == y {
			return true
		}
	}
	return false
}

func (snake *Snake) ChangeDirection(key keyboard.Key) {
	switch key {
	case keyboard.KeyArrowLeft:
		if snake.Direction != RIGHT {
			snake.Direction = LEFT
		}
	case keyboard.KeyArrowRight:
		if snake.Direction != LEFT {
			snake.Direction = RIGHT
		}
	case keyboard.KeyArrowUp:
		if snake.Direction != DOWN {
			snake.Direction = UP
		}
	case keyboard.KeyArrowDown:
		if snake.Direction != UP {
			snake.Direction = DOWN
		}
	default:
		panic("unhandled default case")
	}
}

func (snake *Snake) MoveSnake() {
	head := snake.Body[0]

	switch snake.Direction {
	case LEFT:
		head.Y--
	case RIGHT:
		head.Y++
	case UP:
		head.X--
	case DOWN:
		head.X++
	}

	// 更新蛇的身体
	// 蛇身体的移动是将蛇尾去掉，蛇头插入新的位
	snake.Body = append([]Point{head}, snake.Body[:len(snake.Body)]...)
}
