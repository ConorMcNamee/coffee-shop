package terminal

import (
	"os"
	"sync"
	"unicode/utf8"
)

type Terminal struct {
	ScreenHeight int
	ScreenWidth  int
	Screen       [][]rune
	Frame        []byte
	Lock         sync.Mutex
}

func CreateTermianl(height, width int) (*Terminal, error) {

	screen := make([][]rune, height)
	for i := range screen {
		screen[i] = make([]rune, width)
	}

	FrameBuffer := make([]byte, 1024)

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

	for y, _ := range t.Screen {
		for x, _ := range t.Screen[y] {
			n := utf8.EncodeRune(tmp, t.Screen[x][y])
			t.Frame = append(t.Frame, tmp[:n]...)
		}
		t.Frame = append(t.Frame, '\n')
	}
	return t.Frame
}

func (t *Terminal) WriteText(screen [][]rune, x, y int, text string) error {

	for i, r := range text {
		screen[y+i][x] = r
	}

	t.Screen = screen
	bytes := t.Render()
	os.Stdout.Write(bytes)

	return nil
}
