package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/shibukawa/configdir"
	"github.com/urfave/cli/v2"
)

type Config struct {
	ConfigVersion int              `json:"configVersion"`
	Accounts      []*ConfigAccount `json:"accounts"`

	// add the fields from previous versions here so that we can migrate them
	LegacyAuthToken    string `json:"authToken,omitempty"`
	LegacyRefreshToken string `json:"refreshToken,omitempty"`
}

var ConfigKey = struct{}{}

func LoadConfig(c *cli.Context) error {
	var configPath string
	if c.String("config") != "" {
		configPath = c.String("config")
	} else {
		folder := configDirs.QueryFolderContainsFile("config.json")
		if folder == nil {
			c.Context = context.WithValue(c.Context, ConfigKey, &Config{
				ConfigVersion: 2,
				Accounts:      []*ConfigAccount{},
			})
			return nil
		}
		configPath = filepath.Join(folder.Path, "config.json")
	}

	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read config file from '%v': %w", configPath, err)
	}
	config := &Config{}
	err = json.Unmarshal(configData, config)
	if err != nil {
		return fmt.Errorf("malformed %v: %w", configPath, err)
	}
	if config.ConfigVersion < 2 {
		config = MigrateConfig(config)
		c.Context = context.WithValue(c.Context, ConfigKey, config)
		if err := SaveConfig(c); err != nil {
			return fmt.Errorf("failed to save migrated config: %w", err)
		}
		return nil
	}
	c.Context = context.WithValue(c.Context, ConfigKey, config)
	if len(c.StringSlice("phone-number")) > 0 && len(GetLimitedAccounts(c)) == 0 {
		return fmt.Errorf("Invalid phone number filter: %v, no accounts matching found", c.StringSlice("phone-number"))
	}
	return nil
}

func GetConfig(c *cli.Context) *Config {
	return c.Context.Value(ConfigKey).(*Config)
}

func GetLimitedAccounts(c *cli.Context) []*ConfigAccount {
	cfg := GetConfig(c)
	accounts := make([]*ConfigAccount, 0)
	phoneNumbersStr := c.StringSlice("phone-number")
	for _, account := range cfg.Accounts {
		if len(phoneNumbersStr) == 0 {
			accounts = append(accounts, account)
			continue
		}
		for _, phoneNumber := range phoneNumbersStr {
			if account.PhoneNumber == phoneNumber {
				accounts = append(accounts, account)
				break
			}
		}
	}
	return accounts
}

func MigrateConfig(c *Config) *Config {
	c.ConfigVersion = 2
	if c.Accounts == nil {
		c.Accounts = []*ConfigAccount{}
	}
	if c.LegacyAuthToken != "" && c.LegacyRefreshToken != "" {
		tok, _ := jwt.ParseWithClaims(strings.TrimPrefix(c.LegacyAuthToken, "Bearer "), &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("lolxd"), nil
		})
		claims := tok.Claims.(*jwt.StandardClaims)
		c.Accounts = append(c.Accounts, &ConfigAccount{
			PhoneNumber:  claims.Subject,
			AuthToken:    c.LegacyAuthToken,
			RefreshToken: c.LegacyRefreshToken,
		})
	}
	c.LegacyAuthToken = ""
	c.LegacyRefreshToken = ""
	log.Printf("Migrated config to version 2")
	return c
}

func SaveConfig(c *cli.Context) error {
	cfg := c.Context.Value(ConfigKey).(*Config)
	configPath := c.String("config")
	if configPath == "" {
		folder := configDirs.QueryFolderContainsFile("config.json")
		if folder != nil {
			configPath = filepath.Join(folder.Path, "config.json")
		} else {
			configPath = filepath.Join(configDirs.QueryFolders(configdir.Global)[0].Path, "config.json")
		}
	} else {
		configPath = filepath.Join(configPath, "config.json")
	}
	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}
	configData, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}
	if err := ioutil.WriteFile(configPath, configData, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}
	log.Printf("Saved config to %v", configPath)
	return nil
}
