package console

import (
	"fmt"
	"strings"
)

type Writer interface {
	Write(p []byte) (n int, err error)
	Stop()
	Flush() error
}

type OneLineShower struct {
	baseShower
}

func NewOneLineShower(x, y int, writer Writer, cleanFunc func(x, y int)) *OneLineShower {
	return &OneLineShower{
		baseShower: baseShower{
			X:         x,
			Y:         y,
			Buffer:    make(chan string),
			CleanFunc: cleanFunc,
			Writer:    writer,	
		},
	}
}

func (o *OneLineShower) Show() {
	o.CleanFunc(o.X, o.Y)
	defer o.Writer.Stop()

	for {
		s := <-o.Buffer
		s = strings.ReplaceAll(s, "\n", "")
		s = strings.ReplaceAll(s, "\r", "")
		fmt.Fprintf(o.Writer, "%s\n", s)
		o.Writer.Flush()
	}
}
