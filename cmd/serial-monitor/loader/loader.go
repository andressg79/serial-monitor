package loader

import (
	utilConsole "github.com/andressg79/serial-monitor/internal/pkg/console"
	"github.com/andressg79/serial-monitor/internal/pkg/serial"
	"github.com/andressg79/serial-monitor/internal/service"
	"github.com/andressg79/serial-monitor/internal/service/console"
	"github.com/gosuri/uilive"
)

type Loader struct {
}

// NewLoader returns a new instance of Loader.
//
// No parameters.
// Returns a Loader instance.
func NewLoader() Loader {
	return Loader{}
}

// PortLister returns a new instance of the PortLister struct.
//
// It takes no parameters.
// It returns a PortLister.
func (l Loader) PortLister() service.PortLister {
	return service.NewPortLister(serial.NewSerial(), console.NewPortShower())
}

// InfiniteMonitor returns a service.Monitor that continuously displays the serial data.
//
// It takes no parameters.
// Returns a service.Monitor.
func (l Loader) infiniteMonitor() service.Monitor {
	return service.NewMonitor(
		serial.NewSerial(),
		console.NewHeaderStatusShower(1, 1, utilConsole.ClearScreen),
		console.NewInfiniteShower(1, 1, uilive.New(), utilConsole.ClearScreen),
	)
}

// OneLineMonitor returns a service.Monitor that displays the serial data one line at a time and refres the screen.
//
// It takes no parameters.
// Returns a service.Monitor.
func (l Loader) oneLineMonitor() service.Monitor {
	return service.NewMonitor(
		serial.NewSerial(),
		console.NewHeaderStatusShower(1, 1, utilConsole.ClearScreen),
		console.NewOneLineShower(4, 1, uilive.New(), utilConsole.ClearScreen),
	)
}

// Monitors returns a list of available monitor types and their corresponding monitor functions.
//
// It takes no parameters.
// Returns a list of strings representing the monitor types and a map of monitor type to its corresponding monitor function.
func (l Loader) Monitors() ([]string, map[string]func() service.Monitor) {
	return []string{
			"classic",
			"oneline",
		},
		map[string]func() service.Monitor{
			"classic": l.infiniteMonitor,
			"oneline": l.oneLineMonitor,
		}
}
