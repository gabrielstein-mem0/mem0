package cmd

import (
	"fmt"

	"github.com/mem0ai/mem0/mem0-cli/internal/api"
	"github.com/mem0ai/mem0/mem0-cli/internal/config"
	"github.com/mem0ai/mem0/mem0-cli/internal/output"
	"github.com/spf13/cobra"
)

func newAPIClient(cmd *cobra.Command) (*api.Client, error) {
	flagKey, _ := cmd.Flags().GetString("api-key")
	apiKey := config.ResolveAPIKey(flagKey)
	if apiKey == "" {
		return nil, fmt.Errorf("no API key configured — run \"mem0 auth login\" or set MEM0_API_KEY")
	}

	baseURL, _ := cmd.Flags().GetString("base-url")
	if baseURL == "" {
		cfg, _ := config.Load()
		if cfg != nil && cfg.BaseURL != "" {
			baseURL = cfg.BaseURL
		} else {
			baseURL = config.DefaultBaseURL
		}
	}

	return api.NewClient(apiKey, baseURL), nil
}

func newPrinter(cmd *cobra.Command) *output.Printer {
	format, _ := cmd.Flags().GetString("output")
	return output.NewPrinter(format)
}

func resolveUserID(cmd *cobra.Command) string {
	userID, _ := cmd.Flags().GetString("user-id")
	if userID != "" {
		return userID
	}
	cfg, _ := config.Load()
	if cfg != nil {
		return cfg.UserID
	}
	return ""
}
