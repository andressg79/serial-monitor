package service

import "fmt"

type SerialMonitor interface {
	Connect(portName string, buad int) (bool, error)
	ReadToFun(bufferSize int, into func(buff []byte) error) error
}

type HeaderShower interface {
	Show(port string, buad int)
}

type MonitorShower interface {
	Listener(buff []byte) error
	Show()
}

type Monitor struct {
	serial SerialMonitor
	header HeaderShower
	shower MonitorShower
}

func NewMonitor(serial SerialMonitor, header HeaderShower, shower MonitorShower) Monitor {
	return Monitor{
		serial: serial,
		header: header,
		shower: shower,
	}
}

func (m Monitor) Run(portName string, buad int) error {
	connected, err := m.serial.Connect(portName, buad)
	if err != nil {
		return err
	}

	if !connected {
		return fmt.Errorf("not connected")
	}

	m.header.Show(portName, buad)

	go m.serial.ReadToFun(1, m.shower.Listener)
	m.shower.Show()
	return nil
}
