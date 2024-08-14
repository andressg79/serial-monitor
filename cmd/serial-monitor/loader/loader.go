package loader

import (
	"github.com/andressg79/serial-monitor/internal/pkg/serial"
	"github.com/andressg79/serial-monitor/internal/service"
	"github.com/andressg79/serial-monitor/internal/service/console"
)

type Loader struct {
}

func NewLoader() Loader {
	return Loader{}
}

func (l Loader) PortLister() service.PortLister {
	return service.NewPortLister(serial.NewSerial(), console.NewPortShower())
}
