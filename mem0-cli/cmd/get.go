package cmd

import (
	"github.com/spf13/cobra"
)

func NewGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get <memory-id>",
		Short: "Get a memory by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newAPIClient(cmd)
			if err != nil {
				return err
			}

			mem, err := client.GetMemory(args[0])
			if err != nil {
				return err
			}

			newPrinter(cmd).PrintMemory(mem)
			return nil
		},
	}
}
