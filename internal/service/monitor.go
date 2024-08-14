package service

import (
	"fmt"
	"log"
	"strings"

	"github.com/buger/goterm"
	"github.com/gosuri/uilive"
	"go.bug.st/serial"
)

type Monitor struct {
}

func NewMonitor() Monitor {
	return Monitor{}
}

func (m Monitor) Run(portName string, buad int) error {
	if !m.validPort(portName) {
		return fmt.Errorf("port %s not found", portName)
	}

	mode := &serial.Mode{
		BaudRate: buad,
	}
	port, err := serial.Open(portName, mode)
	if err != nil {
		return err
	}

	buffer := make(chan string)
	go func() {
		var localBuffer string = ""
		buff := make([]byte, 100)
		for {
			n, err := port.Read(buff)
			if err != nil {
				log.Fatal(err)
				break
			}
			if n == 0 {
				break
			}
			localBuffer += string(buff[:n])
			if strings.Contains(localBuffer, "\n") {
				toChannel, _ := strings.CutSuffix(localBuffer, "\n")
				buffer <- toChannel
				localBuffer, _ = strings.CutPrefix(localBuffer, "\n")
			}
		}
	}()

	m.ClearScreen()
	writer := uilive.New()
	defer func() {
		writer.Stop()
	}()

	for {
		s := <-buffer
		s = strings.ReplaceAll(s, "\n", "")
		s = strings.ReplaceAll(s, "\r", "")
		fmt.Fprintf(writer, "%s\n", s)
		writer.Flush()
	}
}

func (m Monitor) validPort(port string) bool {
	ports, err := serial.GetPortsList()
	if err != nil {
		return false
	}

	for _, p := range ports {
		if p == port {
			return true
		}
	}
	return false
}

func (m Monitor) ClearScreen() {
	goterm.Clear()
	goterm.MoveCursor(1, 1)
	goterm.Flush()
}
