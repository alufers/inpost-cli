package main

import (
	"encoding/json"
	"fmt"
	"path/filepath"
)

type Config struct {
	ConfigVersion int              `json:"configVersion"`
	Accounts      []*ConfigAccount `json:"accounts"`
}

type ConfigAccount struct {
	PhoneNumber  string `json:"phoneNumber"`
	AuthToken    string `json:"authToken"`
	RefreshToken string `json:"refreshToken"`
}

func LoadConfig() (*Config, error) {
	folder := configDirs.QueryFolderContainsFile("config.json")
	if folder != nil {
		configData, _ := folder.ReadFile("config.json")
		config := &Config{}
		err := json.Unmarshal(configData, config)
		if err != nil {
			return nil, fmt.Errorf("malformed %v: %w", filepath.Join(folder.Path, "config.json"), err)
		}
		if config.ConfigVersion < 2 {
			if err := MigrateConfig(); err != nil {
				return nil, fmt.Errorf("failed to migrate config: %w", err)
			}
			return LoadConfig()
		}
		return config, nil
	}

	return nil, nil
}

func MigrateConfig() error {
	return nil
}
