package cmd

import (
	"github.com/spf13/cobra"
)

func NewSearchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search <query>",
		Short: "Search memories",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newAPIClient(cmd)
			if err != nil {
				return err
			}

			limit, _ := cmd.Flags().GetInt("limit")
			memories, err := client.SearchMemories(args[0], resolveUserID(cmd), limit)
			if err != nil {
				return err
			}

			newPrinter(cmd).PrintMemories(memories)
			return nil
		},
	}

	cmd.Flags().Int("limit", 10, "Maximum number of results")
	return cmd
}
