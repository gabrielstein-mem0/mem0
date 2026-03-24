package cmd

import (
	"github.com/spf13/cobra"
)

func NewListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List memories",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newAPIClient(cmd)
			if err != nil {
				return err
			}

			limit, _ := cmd.Flags().GetInt("limit")
			page, _ := cmd.Flags().GetInt("page")
			memories, err := client.ListMemories(resolveUserID(cmd), limit, page)
			if err != nil {
				return err
			}

			newPrinter(cmd).PrintMemories(memories)
			return nil
		},
	}

	cmd.Flags().Int("limit", 20, "Number of memories per page")
	cmd.Flags().Int("page", 1, "Page number")
	return cmd
}
