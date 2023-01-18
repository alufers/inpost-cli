package main

import (
	"fmt"
	"log"

	"github.com/urfave/cli/v2"
)

var TerminateOpenCompartment = &cli.Command{
	Name:        "terminate-open-compartment",
	Description: "Terminates a compartment opening session by uuid of the session.  Only one account must be logged in or a phone number filter must be set.",
	Category:    "debugging commands",
	Flags:       []cli.Flag{},
	Action: func(c *cli.Context) error {
		if err := RefreshAllTokens(c); err != nil {
			return fmt.Errorf("failed to refresh token: %v", err)
		}
		limitedAccounts := GetLimitedAccounts(c)
		if len(limitedAccounts) != 1 {
			return fmt.Errorf("must be logged in with one account or limited with a phone number filter (-p)")
		}

		apiClient := limitedAccounts[0].ToClient()
		sessionUuid := c.Args().Get(0)
		if sessionUuid == "" {
			return fmt.Errorf("please pass a sesion uuid")
		}

		resp, _, err := apiClient.DefaultApi.V1CollectTerminateSessionUuidPost(
			c.Context,
			sessionUuid,
		)
		if err != nil {
			return fmt.Errorf("failed to terminate compartment opening: %v", err)
		}

		log.Printf("%#v", resp)

		return nil
	},
}
