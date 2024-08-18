package serial

import (
	"fmt"

	"go.bug.st/serial"
)

type Serial struct {
	portName  string
	buad      int
	connected bool
	err       error
	port      serial.Port
}

// NewSerial returns a new instance of the Serial struct.
//
// No parameters.
// Returns a Serial struct.
func NewSerial() *Serial {
	return &Serial{}
}

// Err returns the error associated with the Serial instance.
//
// No parameters.
// Returns an error.
func (s *Serial) Err() error {
	return s.err
}

// GetPortsList returns a list of available serial ports.
//
// No parameters.
// Returns a slice of strings and an error.
func (s Serial) GetPortsList() ([]string, error) {
	return serial.GetPortsList()
}

// ValidPort checks if the given port is a valid serial port.
//
// The port parameter is a string representing the name of the port to check.
// It returns a boolean indicating whether the port is valid and an error if the port listing fails.
func (s *Serial) ValidPort(port string) (bool, error) {
	ports, err := serial.GetPortsList()
	if err != nil {
		s.err = err
		return false, err
	}

	for _, p := range ports {
		if p == port {
			return true, nil
		}
	}
	return false, nil
}

// Connect establishes a connection to a serial port.
//
// The portName parameter is a string representing the name of the port to connect to.
// The buad parameter is an integer representing the baud rate of the connection.
// It returns a boolean indicating whether the connection was successful and an error if the connection fails.
func (s *Serial) Connect(portName string, buad int) (bool, error) {
	if found, err := s.ValidPort(portName); !found {
		s.err = err
		return false, err
	}

	s.portName = portName
	s.buad = buad

	mode := &serial.Mode{
		BaudRate: buad,
	}
	port, err := serial.Open(portName, mode)
	if err != nil {
		s.err = err
		return false, err
	}

	s.connected = true
	s.port = port
	return true, nil
}

// Close closes the serial port connection.
//
// It takes no parameters.
// It returns an error if the port closure fails.
func (s *Serial) Close() error {
	return s.port.Close()
}

// ReadToFun reads data from the serial port and calls the provided function for each read.
//
// bufferSize: the size of the buffer to read from the serial port.
// into: a function that takes the number of bytes read and the buffer as arguments.
// error: returns an error if there was a problem reading from the serial port or if the provided function returns an error.
func (s *Serial) ReadToFun(bufferSize int, into func(buff []byte) error) error {
	if !s.connected {
		return fmt.Errorf("not connected")
	}

	buff := make([]byte, bufferSize)
	for {
		n, err := s.port.Read(buff)
		if err != nil {
			return err
		}
		if n == 0 {
			break
		}
		if err := into(buff); err != nil {
			return err
		}
	}
	return nil
}
