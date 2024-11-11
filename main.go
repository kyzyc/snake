package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
)

func main() {
	g := Game{}
	err := keyboard.Open()
	if err != nil {
		fmt.Println("open keyboard err:", err)
		return
	}
	defer func() {
		err := keyboard.Close()
		if err != nil {
			fmt.Println("close keyboard err:", err)
			return
		}
	}()
	g.InitGame()
	g.StartGame()
}
