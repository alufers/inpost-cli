package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/alufers/inpost-cli/swagger"
	"github.com/mdp/qrterminal/v3"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
	"rsc.io/qr"
)

/*
Possible status values:
 	NEW,
    ACCEPTED,
    USED,
    REJECTED,
    EXPIRED,
    DELIVERED,
    CANCELED,
    OTHER;

	TODO: add to swagger


*/

func PrintReturnsListTable(parcels []swagger.ReturnTicketNetwork, showDates bool) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Org name", "Created date", "Expry date", "Status", "Return code"})
	for _, v := range parcels {

		expiryTime := ""
		if !v.ExpiryDate.IsZero() {
			if showDates {
				expiryTime = v.ExpiryDate.Format("2006-01-02 15:04:05")
			} else {
				expiryTime = formatDuration(time.Until(v.ExpiryDate))
			}
		}

		var colors []tablewriter.Colors

		if v.Status == "ACCEPTED" {
			var expiryColor = tablewriter.Colors{tablewriter.FgHiMagentaColor}
			// if expiry is near, highlight it in red
			if !v.ExpiryDate.IsZero() && time.Until(v.ExpiryDate) < 12*time.Hour {
				expiryColor = tablewriter.Colors{tablewriter.Bold, tablewriter.BgWhiteColor, tablewriter.BgRedColor}
			}
			colors = []tablewriter.Colors{
				{tablewriter.Bold, tablewriter.FgHiMagentaColor},
				{tablewriter.FgHiMagentaColor},
				expiryColor,
				{tablewriter.FgHiMagentaColor},
				{tablewriter.Bold, tablewriter.FgHiMagentaColor},
			}

			colors = append(colors)

		}
		table.Rich([]string{
			v.OrganizationName,
			v.CreatedDate.Format("2006-01-02 15:04:05"),
			expiryTime,
			v.Status,
			v.ReturnCode,
		}, colors)
	}
	table.Render()
}

func PrintReturnsList(parcels []swagger.ReturnTicketNetwork, showDates bool, qrLevel *qr.Level) {
	for _, v := range parcels {
		fmt.Printf("-----------\n")
		expiryTime := ""
		if !v.ExpiryDate.IsZero() {
			if showDates {
				expiryTime = v.ExpiryDate.Format("2006-01-02 15:04:05")
			} else {
				expiryTime = formatDuration(time.Until(v.ExpiryDate))
			}
		}
		fmt.Printf("Org name: %s\n", v.OrganizationName)
		fmt.Printf("Created date: %s\n", v.CreatedDate.Format("2006-01-02 15:04:05"))
		fmt.Printf("Expiry date: %s\n", expiryTime)
		fmt.Printf("Status: %s\n", v.Status)
		fmt.Printf("Return code: %s\n", v.ReturnCode)
		if strings.TrimSpace(v.QrCode) != "" {
			qrBuf := &bytes.Buffer{}

			qrterminal.GenerateWithConfig(v.QrCode, qrterminal.Config{
				Level:     *qrLevel,
				Writer:    qrBuf,
				BlackChar: qrterminal.BLACK,
				WhiteChar: qrterminal.WHITE,
				QuietZone: 2,
			}) // inpost uses HIGH error correction
			_, err := qr.Encode(v.QrCode, *qrLevel)
			if err != nil {
				fmt.Printf("Error generating QR code: %s", err)
			}
			fmt.Printf("QR code: \n%s\n", qrBuf.String())
		}
	}

}

var ListReturnsCmd = &cli.Command{
	Name:      "list-returns",
	Aliases:   []string{"r", "returns", "r-ls"},
	Usage:     "[--format list|table|json] [--date|-d] [--status STATUS]... -- List pending returns",
	ArgsUsage: "fffff",
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "status",
			Aliases: []string{"s"},
			Usage:   "ACCEPTED etc.",
		},
		&cli.StringFlag{
			Name:        "format",
			Aliases:     []string{"f"},
			DefaultText: "list",
			Value:       "list",
			Usage:       "The output format: list (with QR codes), json or table",
		},
		&cli.StringFlag{
			Name:        "qr",
			Aliases:     []string{"q"},
			DefaultText: "H",
			Value:       "H",
			Usage:       "QR code quality settings: H, M, L, none (no QR code)",
		},
		&cli.BoolFlag{
			Name:    "date",
			Aliases: []string{"d"},
			Usage:   "Show dates instead of time ago",
		},
	},
	Action: func(c *cli.Context) error {
		qrLevel, err := stringToQRLevel(c.String("qr"))
		if err != nil {
			return err
		}
		if err := RefreshAllTokens(c); err != nil {
			return fmt.Errorf("failed to refresh token: %v", err)
		}
		limitedAccounts := GetLimitedAccounts(c)
		// uab, _ := time.Unix(0, 0).AddDate(0, -1, 0).MarshalJSON()
		// ua := "1970-01-01T00:00:00.001Z"

		returnsByPhoneNumber := map[string][]swagger.ReturnTicketNetwork{}
		returnsForJson := []struct {
			swagger.ReturnTicketNetwork
			RecipientPhoneNumber string `json:"recipient_phone_number"`
		}{}
		for _, acc := range limitedAccounts {
			apiClient := acc.ToClient()
			data, _, err := apiClient.DefaultApi.V1ReturnsTicketsGet(c.Context)
			if err != nil {
				return fmt.Errorf("failed to request returns: %w", err)
			}

			filteredParcels := []swagger.ReturnTicketNetwork{}
			for _, parcel := range data.Tickets {
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
			returnsByPhoneNumber[acc.PhoneNumber] = filteredParcels
			for _, parcel := range filteredParcels {
				returnsForJson = append(returnsForJson, struct {
					swagger.ReturnTicketNetwork
					RecipientPhoneNumber string `json:"recipient_phone_number"`
				}{
					ReturnTicketNetwork:  parcel,
					RecipientPhoneNumber: acc.PhoneNumber,
				})
			}
		}

		switch c.String("format") {
		case "list":
			for phoneNumber, parcels := range returnsByPhoneNumber {
				if len(returnsByPhoneNumber) > 1 {
					fmt.Printf("%v:\n", phoneNumber)
				}
				PrintReturnsList(parcels, c.Bool("date"), qrLevel)
			}
		case "table":

			for phoneNumber, parcels := range returnsByPhoneNumber {
				if len(returnsByPhoneNumber) > 1 {
					fmt.Printf("%v:\n", phoneNumber)
				}
				PrintReturnsListTable(parcels, c.Bool("date"))
			}

		case "json":
			s, _ := json.MarshalIndent(returnsForJson, "", "    ")
			fmt.Print(string(s))
		}

		return nil
	},
}
