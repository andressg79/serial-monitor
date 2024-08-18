package console

import "github.com/buger/goterm"

func ClearScreen(x int, y int) {
	goterm.Clear()
	goterm.MoveCursor(x, y)
	goterm.Flush()
}
