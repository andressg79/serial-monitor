package main

import (
	"github.com/andressg79/serial-monitor/cmd/serial-monitor/cmd"
	"github.com/andressg79/serial-monitor/cmd/serial-monitor/loader"
	"github.com/spf13/cobra"
)

// main is the entry point of the serial-monitor application.
//
// It initializes a new loader, sets up the available commands, and executes the root command.
// No parameters.
// No return values.
func main() {
	l := loader.NewLoader()

	showers, monitors := l.Monitors()

	rootCmd := &cobra.Command{Use: "serial-monitor"}
	rootCmd.AddCommand(
		cmd.List(l.PortLister()),
		cmd.Monitor(showers, monitors),
	)

	if err := rootCmd.Execute(); err != nil {
		panic(err.Error())
	}
}
