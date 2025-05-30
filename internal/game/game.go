package game

import (
	"coffee/internal/terminal"
	"time"
)

type CoffeeShop struct {
	coffeePrice float64
	coffeeBeans float64
	milk        float64
	sugar       float64
}

type GameState struct {
	Terminal *terminal.Terminal
}

func (gs *GameState) GameLoop() {
	for {

		time.Sleep(time.Second / 30)
		gs.Terminal.WriteText(gs.Terminal.Screen, 0, 0, "Coffee: ")
	}
}

func (gs *GameState) HandleInput() {

}
