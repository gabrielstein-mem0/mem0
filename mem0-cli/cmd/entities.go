package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewEntitiesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "entities",
		Short: "Manage entities",
	}

	cmd.AddCommand(newEntitiesListCmd())
	cmd.AddCommand(newEntitiesDeleteCmd())
	return cmd
}

func newEntitiesListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all entities",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newAPIClient(cmd)
			if err != nil {
				return err
			}

			entities, err := client.ListEntities()
			if err != nil {
				return err
			}

			newPrinter(cmd).PrintEntities(entities)
			return nil
		},
	}
}

func newEntitiesDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <entity-id>",
		Short: "Delete an entity",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newAPIClient(cmd)
			if err != nil {
				return err
			}

			entityType, _ := cmd.Flags().GetString("type")
			if err := client.DeleteEntity(entityType, args[0]); err != nil {
				return err
			}

			newPrinter(cmd).PrintMessage(fmt.Sprintf("Entity %s deleted.", args[0]))
			return nil
		},
	}

	cmd.Flags().String("type", "user", "Entity type (user, agent, app, run)")
	return cmd
}
