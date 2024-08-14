package console

import (
	"fmt"
)

type PortShower struct {
}

// NewPortShower creates a new instance of the PortShower struct.
//
// It does not take any parameters.
// It returns a PortShower struct.
func NewPortShower() PortShower {
	return PortShower{}
}

// Show displays a list of available ports.
//
// The ports parameter is a slice of strings representing the available ports.
// No return value.
func (p PortShower) Show(ports []string) {
	if len(ports) == 0 {
		fmt.Println("no ports found")
		return
	}

	fmt.Println("Ports:")
	for _, port := range ports {
		fmt.Printf("\t-\t%v\n", port)
	}
}
