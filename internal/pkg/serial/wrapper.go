package serial

import "go.bug.st/serial"

type Serial struct {
}

func NewSerial() Serial {
	return Serial{}
}

func (s Serial) GetPortsList() ([]string, error) {
	return serial.GetPortsList()
}
