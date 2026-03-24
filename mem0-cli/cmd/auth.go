package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/mem0ai/mem0/mem0-cli/internal/config"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Manage authentication",
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate with an API key",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("Enter API key: ")
		reader := bufio.NewReader(os.Stdin)
		key, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read input: %w", err)
		}
		key = strings.TrimSpace(key)

		if key == "" {
			return fmt.Errorf("API key cannot be empty")
		}
		if !strings.HasPrefix(key, "m0-") {
			return fmt.Errorf("invalid API key: must start with \"m0-\"")
		}

		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}
		cfg.APIKey = key

		if err := config.Save(cfg); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}

		color.Green("Authenticated successfully.")
		return nil
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current authentication status",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		if cfg.APIKey == "" {
			fmt.Fprintln(os.Stderr, "Not logged in. Run \"mem0 auth login\" to authenticate.")
			os.Exit(1)
		}

		baseURL := cfg.BaseURL
		if baseURL == "" {
			baseURL = config.DefaultBaseURL
		}

		fmt.Printf("Email:    %s\n", valueOrDash(cfg.Email))
		fmt.Printf("API Key:  %s\n", maskKey(cfg.APIKey))
		fmt.Printf("Base URL: %s\n", baseURL)
		return nil
	},
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Remove stored credentials",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := config.Delete(); err != nil {
			return fmt.Errorf("failed to delete config: %w", err)
		}
		fmt.Println("Logged out successfully.")
		return nil
	},
}

func init() {
	authCmd.AddCommand(loginCmd)
	authCmd.AddCommand(statusCmd)
	authCmd.AddCommand(logoutCmd)
}

func maskKey(key string) string {
	if len(key) <= 8 {
		return strings.Repeat("*", len(key))
	}
	return key[:4] + strings.Repeat("*", len(key)-8) + key[len(key)-4:]
}

func valueOrDash(s string) string {
	if s == "" {
		return "-"
	}
	return s
}
