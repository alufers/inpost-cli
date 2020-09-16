package main

import (
	"fmt"
	"log"

	"github.com/alufers/inpost-cli/swagger"
	"github.com/urfave/cli/v2"
)

var TerminateOpenCompartment = &cli.Command{
	Name:        "terminate-open-compartment",
	Description: "Terminates a compartment opening session by uuid of the session",
	Category:    "debugging commands",
	Flags:       []cli.Flag{},
	Action: func(c *cli.Context) error {
		if err := refreshTokenIfNeeded(c.Context); err != nil {
			return fmt.Errorf("failed to refresh token: %v", err)
		}
		cfg := swagger.NewConfiguration()
		cfg.DefaultHeader["Authorization"] = "Bearer " + config.AuthToken
		apiClient := swagger.NewAPIClient(cfg)
		sessionUuid := c.Args().Get(0)
		if sessionUuid == "" {
			return fmt.Errorf("Please pass a sesion uuid")
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
