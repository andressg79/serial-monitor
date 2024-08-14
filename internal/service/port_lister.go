package service

type SerialLister interface {
	GetPortsList() ([]string, error)
}

type PortShower interface {
	Show([]string)
}

type PortLister struct {
	serial SerialLister
	shower PortShower
}

// NewPortLister returns a new instance of the PortLister struct.
//
// It takes a SerialLister and a PortShower as parameters.
// It returns a PortLister.
func NewPortLister(serial SerialLister, shower PortShower) PortLister {
	return PortLister{
		serial: serial,
		shower: shower,
	}
}

// List lists the available serial ports.
//
// It takes no parameters.
// It returns an error if the port listing fails.
func (p PortLister) List() error {
	ports, err := p.serial.GetPortsList()
	if err != nil {
		return err
	}
	p.shower.Show(ports)
	return nil
}
