package main

import (
	"log"
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
		tok, _ := jwt.ParseWithClaims(strings.TrimPrefix(config.AuthToken, "Bearer "), &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("lolxd"), nil
		})
		claims := tok.Claims.(*jwt.StandardClaims)
		log.Printf("claims: %+v", claims)
		log.Printf("time until expiration: %v", time.Unix(claims.ExpiresAt, 0).Sub(time.Now()))
		return nil
	},
}
