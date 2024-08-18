package cmd

import (
	"fmt"
	"slices"
	"strings"

	"github.com/andressg79/serial-monitor/internal/service"
	"github.com/spf13/cobra"
)

// Monitor creates a new Cobra command for monitoring serial ports.
//
// The function takes in two parameters:
// - showers: a slice of strings representing the available shower modes
// - monitors: a map of shower modes to functions that return a service.Monitor
//
// The function returns a pointer to a Cobra command.
func Monitor(showers []string, monitors map[string]func() service.Monitor) *cobra.Command {
	var port string
	var buad int
	var shower string

	var cmdMonitor = &cobra.Command{
		Use:     "monitor",
		Example: "monitor -p /dev/ttyUSB0 -b 9600",
		Short:   "Monitor serial ports",
		Long:    `Given a port and a baudrate, this command will show the messages received from the serial port.`,
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Monitor Handler")
			if port == "" {
				fmt.Println("Port is required")
				return
			}
			if buad == 0 {
				fmt.Println("Baudrate is required")
				return
			}

			if !slices.Contains(showers, shower) {
				fmt.Println("Shower mode is required")
				fmt.Println("Available shower modes: ", strings.Join(showers, ", "))
				return
			}

			s, ok := monitors[shower]
			if !ok {
				fmt.Println("Shower mode is not supported")
				return
			}

			if err := s().Run(port, buad); err != nil {
				fmt.Println(err.Error())
			}
		},
	}

	cmdMonitor.Flags().StringVarP(&port, "port", "p", "", "Path or name of serial port")
	cmdMonitor.Flags().IntVarP(&buad, "buad", "b", 9600, "Connection buadrate")
	cmdMonitor.MarkFlagsRequiredTogether("port", "buad")
	cmdMonitor.MarkFlagRequired("port")
	cmdMonitor.Flags().StringVarP(&shower, "shower", "s", "classic", "Shower mode")

	return cmdMonitor
}
