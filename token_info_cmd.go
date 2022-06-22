package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/urfave/cli/v2"
)

var TokenInfoCmd = &cli.Command{
	Name:     "token-info",
	Category: "debugging commands",
	Flags:    []cli.Flag{},
	Action: func(c *cli.Context) error {
		didPrint := false
		for _, acc := range GetLimitedAccounts(c) {
			didPrint = true
			tok, _ := jwt.ParseWithClaims(strings.TrimPrefix(acc.AuthToken, "Bearer "), &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte("lolxd"), nil
			})
			claims := tok.Claims.(*jwt.StandardClaims)
			fmt.Printf("%v\n", acc.PhoneNumber)
			fmt.Printf("  claims: %+v\n", claims)
			fmt.Printf("  time until expiration: %v\n", time.Unix(claims.ExpiresAt, 0).Sub(time.Now()))
		}
		if !didPrint {
			return fmt.Errorf("no accounts found")
		}
		return nil
	},
}
