package main

import (
	"coffee/internal/game"
	"coffee/internal/terminal"
	"fmt"
)

func main() {

	gameTerm, err := terminal.CreateTerminal(50, 50)
	coffeeGame := game.GameState{
		Terminal: gameTerm,
	}

	coffeeGame.GameLoop()

	if err != nil {
		fmt.Errorf("Cannot create terminal window", err)
	}

}
