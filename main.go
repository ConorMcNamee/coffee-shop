package main

import (
	"coffee/internal/terminal"
	"fmt"
)

func main() {
	term, err := terminal.CreateTermianl(20, 20)
	if err != nil {
		fmt.Errorf("Cannot create terminal: ", err)
	}

	term.WriteText(term.Screen, 0, 0, "Hello World")
}
