package main

import (
	"github.com/andressg79/serial-monitor/cmd/serial-monitor/cmd"
	"github.com/andressg79/serial-monitor/cmd/serial-monitor/loader"
	"github.com/spf13/cobra"
)

func main() {
	l := loader.NewLoader()

	rootCmd := &cobra.Command{Use: "serial-monitor"}
	rootCmd.AddCommand(
		cmd.List(l.PortLister()),
		cmd.Monitor(),
	)

	if err := rootCmd.Execute(); err != nil {
		panic(err.Error())
	}
}
