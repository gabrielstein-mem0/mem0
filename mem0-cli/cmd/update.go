package cmd

import (
	"github.com/spf13/cobra"
)

func NewUpdateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "update <memory-id> <text>",
		Short: "Update a memory",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newAPIClient(cmd)
			if err != nil {
				return err
			}

			mem, err := client.UpdateMemory(args[0], args[1])
			if err != nil {
				return err
			}

			newPrinter(cmd).PrintMemory(mem)
			return nil
		},
	}
}
