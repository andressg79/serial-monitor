package console

import (
	"fmt"
)

type InfiniteShower struct {
	baseShower
}

func NewInfiniteShower(x, y int, writer Writer, cleanFunc func(x, y int)) *InfiniteShower {
	return &InfiniteShower{
		baseShower: baseShower{
			X:         x,
			Y:         y,
			Buffer:    make(chan string),
			CleanFunc: cleanFunc,
			Writer:    writer,
		},
	}
}

func (i *InfiniteShower) Show() {
	i.CleanFunc(i.X, i.Y)
	defer i.Writer.Stop()

	for {
		s := <-i.Buffer
		fmt.Fprintf(i.Writer, s)
	}
}
