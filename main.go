package main

import (
	"coffee/internal/game"
	"coffee/internal/terminal"
	"fmt"
)

func main() {
	term, err := terminal.CreateTermianl(20, 20)
	if err != nil {
		fmt.Errorf("Cannot create terminal: ", err)
	}

	game := game.GameState{Terminal: term}
	game.GameLoop()
}
