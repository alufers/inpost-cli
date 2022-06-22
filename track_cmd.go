package main

import (
	"encoding/json"
	"fmt"

	"github.com/alufers/inpost-cli/trackingapi"
	"github.com/urfave/cli/v2"
)

var TrackCmd = &cli.Command{
	Name:        "track",
	Usage:       "[--format=text|json] <SHIPMENT_NUMBER> -- show tracking",
	Description: "Shows detailed tracking information about the parcel",

	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "format",
			Aliases:     []string{"f"},
			DefaultText: "text",
			Value:       "text",
			Usage:       "The output format: json or text",
		},
	},
	Action: func(c *cli.Context) error {
		if err := RefreshAllTokens(c); err != nil {
			return fmt.Errorf("failed to refresh token: %v", err)
		}

		shipmentNumber, _, err := resolveShipmentNumber(c, c.Args().Get(0))
		if err != nil {
			return fmt.Errorf("failed to resolve shipment number: %w", err)
		}
		data, err := trackingapi.GetTrackingData(shipmentNumber)
		if err != nil {
			return fmt.Errorf("failed to get tracking data: %w", err)
		}

		if c.String("format") == "json" {
			s, _ := json.MarshalIndent(data, "", "    ")
			fmt.Print(string(s))
		} else {
			for _, p := range data.TrackingDetails {
				fmt.Printf("%v %v (%v)\n", p.Datetime.String(), p.Status, p.OriginStatus)
			}
		}

		return nil
	},
}
