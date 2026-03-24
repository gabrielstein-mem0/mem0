package cmd

import (
	"fmt"
	"os"

	"github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "mem0",
	Short:         "CLI for the Mem0 Platform",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	defaultOutput := "json"
	if isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		defaultOutput = "table"
	}

	rootCmd.PersistentFlags().StringP("output", "o", defaultOutput, "Output format: json or table")
	rootCmd.PersistentFlags().String("api-key", "", "API key (overrides MEM0_API_KEY env and config)")
	rootCmd.PersistentFlags().String("base-url", "", "API base URL")
	rootCmd.PersistentFlags().String("user-id", "", "User ID for memory operations")

	rootCmd.AddCommand(authCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(NewAddCmd())
	rootCmd.AddCommand(NewSearchCmd())
	rootCmd.AddCommand(NewListCmd())
	rootCmd.AddCommand(NewGetCmd())
	rootCmd.AddCommand(NewUpdateCmd())
	rootCmd.AddCommand(NewDeleteCmd())
	rootCmd.AddCommand(NewEntitiesCmd())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
