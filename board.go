package main

type Board struct {
	Width  int
	Height int
}

func CreateBoard(width int, height int) *Board {
	return &Board{Width: width, Height: height}
}
