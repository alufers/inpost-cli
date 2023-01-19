package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/alufers/inpost-cli/swagger"
	"github.com/antihax/optional"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
)

func PrintParcelListTable(parcels []swagger.Parcel, showDates bool) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Shipment number", "Sender", "Status", "Pickup point", "City", "Open code", "Stored for"})
	for _, v := range parcels {
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
			if showDates {
				expiryTime = v.ExpiryDate.Format("2006-01-02 15:04:05")
			} else {
				expiryTime = formatDuration(time.Until(v.ExpiryDate))
			}
		}
		timeSinceStatus := ""
		for _, s := range v.StatusHistory {
			if s.Status == v.Status {
				if showDates {
					timeSinceStatus = s.Date.Format("2006-01-02 15:04:05")
				} else {
					timeSinceStatus = formatDuration(time.Since(s.Date))
				}
				break
			}
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
			if !v.ExpiryDate.IsZero() && time.Until(v.ExpiryDate) < 12*time.Hour {
				colors = append(colors, tablewriter.Colors{tablewriter.Bold, tablewriter.BgWhiteColor, tablewriter.BgRedColor})
			} else {
				colors = append(colors, tablewriter.Colors{tablewriter.FgHiYellowColor})
			}
		}
		table.Rich([]string{v.ShipmentNumber, v.SenderName, v.Status + prependSpaceIfNotEmpty(timeSinceStatus), pickupPoint, pickupCity, v.OpenCode, expiryTime}, colors)
	}
	table.Render()
}

var ListParcelsCmd = &cli.Command{
	Name:      "list-parcels",
	Aliases:   []string{"p", "parcels", "ls"},
	Usage:     "[--format table|json] [--date|-d] [--status STATUS]...  -- List parcels assignet to your accounts",
	ArgsUsage: "fffff",
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
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
		&cli.BoolFlag{
			Name:    "date",
			Aliases: []string{"d"},
			Usage:   "Show dates instead of time ago",
		},
	},
	Action: func(c *cli.Context) error {
		if err := RefreshAllTokens(c); err != nil {
			return fmt.Errorf("failed to refresh token: %v", err)
		}
		limitedAccounts := GetLimitedAccounts(c)
		// uab, _ := time.Unix(0, 0).AddDate(0, -1, 0).MarshalJSON()
		ua := "1970-01-01T00:00:00.001Z"

		parcelsByPhoneNumber := map[string][]swagger.Parcel{}
		parcelsForJson := []struct {
			swagger.Parcel
			RecipientPhoneNumber string `json:"recipient_phone_number"`
		}{}
		for _, acc := range limitedAccounts {
			apiClient := acc.ToClient()
			data, _, err := apiClient.DefaultApi.V1ParcelGet(c.Context, &swagger.DefaultApiV1ParcelGetOpts{
				UpdatedAfter: optional.NewString(ua),
			})
			if err != nil {
				return fmt.Errorf("failed to request parcels: %w", err)
			}

			filteredParcels := []swagger.Parcel{}
			for _, parcel := range data {
				if len(c.StringSlice("status")) == 0 {
					filteredParcels = append(filteredParcels, parcel)
				}
				for _, status := range c.StringSlice("status") {
					if parcel.Status == status {
						filteredParcels = append(filteredParcels, parcel)
						break
					}
				}
			}
			parcelsByPhoneNumber[acc.PhoneNumber] = filteredParcels
			for _, parcel := range filteredParcels {
				parcelsForJson = append(parcelsForJson, struct {
					swagger.Parcel
					RecipientPhoneNumber string `json:"recipient_phone_number"`
				}{
					Parcel:               parcel,
					RecipientPhoneNumber: acc.PhoneNumber,
				})
			}
		}

		switch c.String("format") {
		case "table":

			for phoneNumber, parcels := range parcelsByPhoneNumber {
				if len(parcelsByPhoneNumber) > 1 {
					fmt.Printf("%v:\n", phoneNumber)
				}
				PrintParcelListTable(parcels, c.Bool("date"))
			}

		case "json":
			s, _ := json.MarshalIndent(parcelsForJson, "", "    ")
			fmt.Print(string(s))
		}

		return nil
	},
}
