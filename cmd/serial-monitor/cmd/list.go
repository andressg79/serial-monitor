package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"
)

type PortLister interface {
	List() error
}

// List returns a cobra command for listing available serial ports.
//
// It takes a PortLister interface as a parameter.
// It returns a pointer to a cobra Command.
func List(lister PortLister) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List the available serial ports",
		Long:  `List the available serial ports.`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if err := lister.List(); err != nil {
				slog.Error("error listing ports", slog.Any("error", err))
			}
		},
	}
}
