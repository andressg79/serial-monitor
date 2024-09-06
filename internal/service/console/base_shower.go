package console

import "strings"

type baseShower struct {
	X, Y        int
	CuterString string
	Buffer      chan string
	LocalBuffer string
	CleanFunc   func(x, y int)
	Writer      Writer
}

func (b *baseShower) Listener(buff []byte) error {
	b.LocalBuffer += string(buff)
	if strings.Contains(b.LocalBuffer, b.CuterString) {
		toChannel, _ := strings.CutSuffix(b.LocalBuffer, b.CuterString)
		b.Buffer <- toChannel
		b.LocalBuffer, _ = strings.CutPrefix(b.LocalBuffer, b.CuterString)
	}
	return nil
}
