package cmd

import (
	"fmt"

	"github.com/andressg79/serial-monitor/internal/service"
	"github.com/spf13/cobra"
)

func Monitor() *cobra.Command {
	var port string
	var buad int
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
			m := service.NewMonitor()
			if err := m.Run(port, buad); err != nil {
				fmt.Println(err.Error())
			}
		},
	}

	cmdMonitor.Flags().StringVarP(&port, "port", "p", "", "Path or name of serial port")
	cmdMonitor.Flags().IntVarP(&buad, "buad", "b", 9600, "Connection buadrate")

	return cmdMonitor
}
