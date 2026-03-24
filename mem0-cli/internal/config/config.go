package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const DefaultBaseURL = "https://api.mem0.ai"

type Config struct {
	APIKey  string `json:"api_key,omitempty"`
	UserID  string `json:"user_id,omitempty"`
	Email   string `json:"email,omitempty"`
	BaseURL string `json:"base_url,omitempty"`
}

func ConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".mem0", "config.json")
}

func Load() (*Config, error) {
	data, err := os.ReadFile(ConfigPath())
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{}, nil
		}
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func Save(cfg *Config) error {
	dir := filepath.Dir(ConfigPath())
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(ConfigPath(), data, 0600)
}

func Delete() error {
	err := os.Remove(ConfigPath())
	if err != nil && os.IsNotExist(err) {
		return nil
	}
	return err
}

// ResolveAPIKey returns the API key using precedence: flag > env > config file.
func ResolveAPIKey(flagValue string) string {
	if flagValue != "" {
		return flagValue
	}
	if v := os.Getenv("MEM0_API_KEY"); v != "" {
		return v
	}
	cfg, err := Load()
	if err != nil {
		return ""
	}
	return cfg.APIKey
}
