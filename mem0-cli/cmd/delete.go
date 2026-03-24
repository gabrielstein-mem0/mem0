package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [memory-id]",
		Short: "Delete a memory or all memories",
		RunE: func(cmd *cobra.Command, args []string) error {
			all, _ := cmd.Flags().GetBool("all")

			client, err := newAPIClient(cmd)
			if err != nil {
				return err
			}

			p := newPrinter(cmd)

			if all {
				userID := resolveUserID(cmd)
				if userID == "" {
					return fmt.Errorf("--user-id is required when using --all")
				}
				if err := client.DeleteAllMemories(userID); err != nil {
					return err
				}
				p.PrintMessage("All memories deleted.")
				return nil
			}

			if len(args) != 1 {
				return fmt.Errorf("provide a memory ID or use --all")
			}

			if err := client.DeleteMemory(args[0]); err != nil {
				return err
			}
			p.PrintMessage(fmt.Sprintf("Memory %s deleted.", args[0]))
			return nil
		},
	}

	cmd.Flags().Bool("all", false, "Delete all memories (requires --user-id)")
	return cmd
}
