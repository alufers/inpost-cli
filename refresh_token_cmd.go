package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var RefreshTokenCmd = &cli.Command{
	Name:     "refresh-token",
	Category: "debugging commands",
	Flags:    []cli.Flag{},
	Action: func(c *cli.Context) error {
		err := RefreshAllTokens(c)
		if err != nil {
			return fmt.Errorf("failed to refresh token: %v", err)
		}

		return nil
	},
}
