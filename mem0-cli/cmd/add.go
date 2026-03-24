package cmd

import (
	"github.com/spf13/cobra"
)

func NewAddCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add <text>",
		Short: "Add a new memory",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newAPIClient(cmd)
			if err != nil {
				return err
			}

			mem, err := client.AddMemory(args[0], resolveUserID(cmd))
			if err != nil {
				return err
			}

			newPrinter(cmd).PrintMemory(mem)
			return nil
		},
	}
}
