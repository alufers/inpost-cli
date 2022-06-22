package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/alufers/inpost-cli/swagger"
	"github.com/antihax/optional"
	"github.com/dgrijalva/jwt-go"
	"github.com/urfave/cli/v2"
)

func doRefreshToken(c *cli.Context, cfgAccount *ConfigAccount) error {

	apiClient := cfgAccount.ToClient()
	resp, _, err := apiClient.DefaultApi.V1AuthenticatePost(c.Context, &swagger.DefaultApiV1AuthenticatePostOpts{
		Body: optional.NewInterface(&swagger.AuthenticateRequest{
			PhoneOS:      "Android",
			RefreshToken: cfgAccount.RefreshToken,
		}),
	})
	if err != nil {
		return fmt.Errorf("failed to refresh the token: %w", err)
	}
	cfgAccount.AuthToken = resp.AuthToken

	return SaveConfig(c)
}

func refreshTokenIfNeeded(c *cli.Context, cfgAccount *ConfigAccount) error {
	if cfgAccount != nil && cfgAccount.RefreshToken != "" {
		tok, _ := jwt.ParseWithClaims(strings.TrimPrefix(cfgAccount.AuthToken, "Bearer "), &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("lolxd"), nil
		})
		claims := tok.Claims.(*jwt.StandardClaims)
		expirationDate := time.Unix(claims.ExpiresAt, 0)
		if expirationDate.Before(time.Now()) {
			log.Printf("Token expired %v ago, refreshing...", time.Now().Sub(expirationDate))
			return doRefreshToken(c, cfgAccount)
		}
	}
	return nil // we do not return an error, the methods will fail with 401 anyway
}

func RefreshAllTokens(c *cli.Context) error {

	for _, cfgAccount := range GetLimitedAccounts(c) {
		err := refreshTokenIfNeeded(c, cfgAccount)
		if err != nil {
			return err
		}
	}
	return nil
}
