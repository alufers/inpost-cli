package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/alufers/inpost-cli/swagger"
	"github.com/antihax/optional"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
)

var ListParcelsCmd = &cli.Command{
	Name:    "list-parcels",
	Aliases: []string{"p", "parcels", "ls"},
	Usage:   "lists parcels tied to the currently logged-in user",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "status",
			Aliases: []string{"s"},
			Usage:   "delivered, ready_to_pickup, confirmed etc.",
		},
		&cli.StringFlag{
			Name:        "format",
			Aliases:     []string{"f"},
			DefaultText: "table",
			Value:       "table",
			Usage:       "The output format: json or table",
		},
	},
	Action: func(c *cli.Context) error {
		if err := refreshTokenIfNeeded(c.Context); err != nil {
			return fmt.Errorf("failed to refresh token: %v", err)
		}
		cfg := swagger.NewConfiguration()
		cfg.DefaultHeader["Authorization"] = "Bearer " + config.AuthToken
		apiClient := swagger.NewAPIClient(cfg)
		// uab, _ := time.Unix(0, 0).AddDate(0, -1, 0).MarshalJSON()
		ua := "1970-01-01T00:00:00.001Z"
		data, _, err := apiClient.DefaultApi.V1ParcelGet(c.Context, &swagger.DefaultApiV1ParcelGetOpts{
			UpdatedAfter: optional.NewString(ua),
		})
		if err != nil {
			return fmt.Errorf("failed to request parcels: %w", err)
		}
		filteredParcels := []swagger.Parcel{}
		for _, parcel := range data {
			if c.String("status") == "" || parcel.Status == c.String("status") {
				filteredParcels = append(filteredParcels, parcel)
			}
		}
		switch c.String("format") {
		case "table":
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Shipment number", "Sender", "Status", "Pickup point", "City", "Open code", "Stored for"})
			for _, v := range filteredParcels {
				var pickupPoint string
				var pickupCity string
				if v.PickupPoint != nil {
					pickupPoint = fmt.Sprintf("%v %v %v",
						v.PickupPoint.Name,
						v.PickupPoint.AddressDetails.Street, v.PickupPoint.AddressDetails.BuildingNumber,
					)
					pickupCity = v.PickupPoint.AddressDetails.City
				}
				expiryTime := ""
				if !v.ExpiryDate.IsZero() {
					expiryTime = v.ExpiryDate.Sub(time.Now()).String()
					expiryTime = regexp.MustCompile("[0-9\\.]*s$").ReplaceAllString(expiryTime, "")
				}
				var colors []tablewriter.Colors
				if v.Status == "delivered" {
					colors = []tablewriter.Colors{
						{tablewriter.FgHiBlackColor},
						{tablewriter.FgHiBlackColor},
						{tablewriter.FgHiBlackColor},
						{tablewriter.FgHiBlackColor},
						{tablewriter.FgHiBlackColor},
						{tablewriter.FgHiBlackColor},
						{tablewriter.FgHiBlackColor},
					}
				}
				if v.Status == "ready_to_pickup" {
					colors = []tablewriter.Colors{
						{tablewriter.Bold, tablewriter.FgHiYellowColor},
						{tablewriter.FgHiYellowColor},
						{tablewriter.FgHiYellowColor},
						{tablewriter.FgHiYellowColor},
						{tablewriter.FgHiYellowColor},
						{tablewriter.Bold, tablewriter.FgHiYellowColor},
					}
					if !v.ExpiryDate.IsZero() && v.ExpiryDate.Sub(time.Now()) < 12*time.Hour {
						colors = append(colors, tablewriter.Colors{tablewriter.Bold, tablewriter.BgWhiteColor, tablewriter.BgRedColor})
					} else {
						colors = append(colors, tablewriter.Colors{tablewriter.FgHiYellowColor})
					}
				}
				table.Rich([]string{v.ShipmentNumber, v.SenderName, v.Status, pickupPoint, pickupCity, v.OpenCode, expiryTime}, colors)
			}
			table.Render()
		case "json":
			s, _ := json.MarshalIndent(filteredParcels, "", "    ")
			fmt.Print(string(s))
		}

		return nil
	},
}
