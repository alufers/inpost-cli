package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/alufers/inpost-cli/swagger"
	"github.com/antihax/optional"
	"github.com/dgrijalva/jwt-go"
	"github.com/shibukawa/configdir"
	"github.com/urfave/cli/v2"
)

var configDirs = configdir.New("alufers", "inpost-cli")
var config *swagger.ConfirmSmsResponse = nil

func loadConfig() {
	folder := configDirs.QueryFolderContainsFile("config.json")
	if folder != nil {
		configData, _ := folder.ReadFile("config.json")
		config = &swagger.ConfirmSmsResponse{}
		err := json.Unmarshal(configData, config)
		if err != nil {
			log.Fatalf("failed to parse config.json: %v", err)
		}

	}
}

func doRefreshToken(ctx context.Context) error {
	cfg := swagger.NewConfiguration()
	cfg.DefaultHeader["Authorization"] = "Bearer " + config.AuthToken
	apiClient := swagger.NewAPIClient(cfg)
	resp, _, err := apiClient.DefaultApi.V1AuthenticatePost(ctx, &swagger.DefaultApiV1AuthenticatePostOpts{
		Body: optional.NewInterface(&swagger.AuthenticateRequest{
			PhoneOS:      "Android",
			RefreshToken: config.RefreshToken,
		}),
	})
	if err != nil {
		return fmt.Errorf("failed to refresh the token: %w", err)
	}
	config.AuthToken = resp.AuthToken
	folders := configDirs.QueryFolders(configdir.Global)
	marshaledJSON, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal config.json: %w", err)
	}
	err = folders[0].WriteFile("config.json", marshaledJSON)
	if err != nil {
		return fmt.Errorf("failed to write config.json: %w", err)
	}
	return nil
}

func refreshTokenIfNeeded(ctx context.Context) error {
	if config != nil && config.RefreshToken != "" {
		tok, _ := jwt.ParseWithClaims(strings.TrimPrefix(config.AuthToken, "Bearer "), &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("lolxd"), nil
		})
		claims := tok.Claims.(*jwt.StandardClaims)
		expirationDate := time.Unix(claims.ExpiresAt, 0)
		if expirationDate.Before(time.Now()) {
			log.Printf("Token expired %v ago, refreshing...", time.Now().Sub(expirationDate))
			return doRefreshToken(ctx)
		}
	}
	return nil // we do not return an error, the methods will fail with 401 anyway
}

func main() {

	loadConfig()

	app := &cli.App{
		Usage:       "A CLI utility replacing the InPost app",
		Description: "This utility allows you to list your packages, generate QR codes in the terminal to collect them and remotely open compartments. \n   Warning: There is no limit to how far away you can open the compartments from. Use with caution as you can get your package stolen this way.",
		Version:     InpostCliVersion,
		Commands: []*cli.Command{
			LoginCmd,
			ListParcelsCmd,
			ParcelInfoCmd,
			OpenCompartmentCmd,
			TerminateOpenCompartment,
			RefreshTokenCmd,
			TokenInfoCmd,
		},
	}
	app.EnableBashCompletion = true
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
