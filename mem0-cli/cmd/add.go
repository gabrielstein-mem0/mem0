package cmd

import (
	"fmt"

	"github.com/fatih/color"
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

			result, err := client.AddMemory(args[0], resolveUserID(cmd))
			if err != nil {
				return err
			}

			p := newPrinter(cmd)
			if p.Format == "json" {
				p.PrintJSON(result)
			} else {
				color.Green("Memory queued successfully.")
				fmt.Printf("Event ID: %s\n", result.EventID)
				fmt.Printf("Status:   %s\n", result.Status)
			}
			return nil
		},
	}
}
