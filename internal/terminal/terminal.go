package terminal

import (
	"fmt"
	"os"
	"sync"

	"golang.org/x/term"
	"unicode/utf8"
)

type Terminal struct {
	ScreenHeight int
	ScreenWidth  int
	Screen       [][]rune
	Frame        []byte
	Lock         sync.Mutex
}

func CreateTerminal(height, width int) (*Terminal, error) {

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Errorf("Couldnt make term raw: ", err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	screen := make([][]rune, height)
	for i := range screen {
		screen[i] = make([]rune, width)
	}

	FrameBuffer := make([]byte, height*width*4)

	terminal := &Terminal{
		ScreenHeight: height,
		ScreenWidth:  width,
		Screen:       screen,
		Frame:        FrameBuffer,
	}

	return terminal, nil
}

func (t *Terminal) Render() []byte {
	t.Frame = t.Frame[:0]
	tmp := make([]byte, 4)

	for x, _ := range t.Screen {
		for y, _ := range t.Screen[x] {
			n := utf8.EncodeRune(tmp, t.Screen[x][y])
			t.Frame = append(t.Frame, tmp[:n]...)
		}
		t.Frame = append(t.Frame, '\n')
	}
	return t.Frame
}

func (t *Terminal) Flush() {
	for y := range t.Screen {
		for x := range t.Screen {
			t.Screen[x][y] = ' '
		}
	}

	os.Stdout.Write([]byte("\033[H\033[2J"))
}

func (t *Terminal) WriteText(screen [][]rune, x, y int, text string) error {

	t.Flush()
	for i, r := range text {
		screen[x][y+i] = r
	}

	t.Screen = screen
	bytes := t.Render()
	os.Stdout.Write(bytes)

	return nil
}
