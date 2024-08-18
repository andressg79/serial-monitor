package console

import "fmt"

type HeaderStatusShower struct {
	X, Y      int
	CleanFunc func(x, y int)
}

func NewHeaderStatusShower(x, y int, cleanFunc func(x, y int)) *HeaderStatusShower {
	return &HeaderStatusShower{
		X:         x,
		Y:         y,
		CleanFunc: cleanFunc,
	}
}

func (h *HeaderStatusShower) Show(port string, buad int) {
	h.CleanFunc(h.X, h.Y)
	fmt.Println(UnderlineText("Serial Monitor, "))
	fmt.Println("\t - port: ", BloodText(port))
	fmt.Println("\t - buad: ", YellowText(fmt.Sprintf("%d", buad)))
	fmt.Println("------------------------------------------")
}
