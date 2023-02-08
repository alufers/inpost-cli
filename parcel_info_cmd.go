package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/alufers/inpost-cli/swagger"
	"github.com/antihax/optional"
	"github.com/mdp/qrterminal/v3"
	"github.com/urfave/cli/v2"
	"rsc.io/qr"

	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

var ParcelInfoCmd = &cli.Command{
	Name:        "parcel-info",
	Aliases:     []string{"qr", "info"},
	Description: "shows detailed information about a parcel, including a QR code",
	Usage:       "[--format=text|json] [--qr H|M|L] <SHIPMENT_NUMBER> -- show info about a package and a QR code",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "format",
			Aliases:     []string{"f"},
			DefaultText: "text",
			Value:       "text",
			Usage:       "The output format: json or text (with a QR code)",
		},
		&cli.StringFlag{
			Name:        "qr",
			Aliases:     []string{"q"},
			DefaultText: "H",
			Value:       "H",
			Usage:       "QR code quality settings: H, M, L, none (no QR code)",
		},
	},
	Action: func(c *cli.Context) error {
		level, err := stringToQRLevel(c.String("qr"))
		if err != nil {
			return err
		}

		if err := RefreshAllTokens(c); err != nil {
			return fmt.Errorf("failed to refresh token: %v", err)
		}

		shipmentNumber, apiClient, err := resolveShipmentNumber(c, c.Args().Get(0))
		if err != nil {
			return fmt.Errorf("failed to resolve shipment number: %w", err)
		}
		parcels, _, err := apiClient.DefaultApi.V1ParcelsGet(c.Context, &swagger.DefaultApiV1ParcelsGetOpts{
			ShipmentNumbers: optional.NewString(shipmentNumber),
		})
		if err != nil {
			return fmt.Errorf("failed to get parcels: %v", err)
		}

		if len(parcels) <= 0 {
			return fmt.Errorf("no such parcel found")
		}
		p := parcels[0]
		infoBuf := &bytes.Buffer{}
		switch c.String("format") {
		case "text":
			fmt.Fprintf(infoBuf, "Shipment number: %v\n", p.ShipmentNumber)
			fmt.Fprintf(infoBuf, "Sender name: %v\n", p.SenderName)
			fmt.Fprintf(infoBuf, "Status: %v\n", p.Status)
			fmt.Fprintf(infoBuf, "Recipient phone number: %v\n", p.PhoneNumber)
			if p.OpenCode != "" {
				fmt.Fprintf(infoBuf, "Open code: %v\n", p.OpenCode)
			}
			if !p.ExpiryDate.IsZero() {
				fmt.Fprintf(infoBuf, "Expiry time: %v\n", formatDuration(time.Until(p.ExpiryDate)))
			}
			if p.PickupPoint != nil {
				fmt.Fprintf(infoBuf, "Pickup point:\n")
				if p.PickupPoint.AddressDetails != nil {
					fmt.Fprintf(infoBuf, "    Address: %v%v%v,\n",
						p.PickupPoint.AddressDetails.Street,
						prependSpaceIfNotEmpty(p.PickupPoint.AddressDetails.BuildingNumber),
						prependSpaceIfNotEmpty(p.PickupPoint.AddressDetails.FlatNumber),
					)
					fmt.Fprintf(infoBuf, "    City: %v, %v\n",
						p.PickupPoint.AddressDetails.City,
						p.PickupPoint.AddressDetails.Province,
					)
					fmt.Fprintf(infoBuf, "    Location description: %v\n", p.PickupPoint.LocationDescription)
				}
			}
			if p.OpenCode != "" {
				fmt.Fprintf(infoBuf, "QR code URL: https://inpost.pl/%v/code/%v\n", p.PhoneNumber, p.OpenCode)
			}
			qrBuf := &bytes.Buffer{}
			var qrWidth int
			if p.QrCode != "" && level != nil {
				qrterminal.GenerateWithConfig(p.QrCode, qrterminal.Config{
					Level:     *level,
					Writer:    qrBuf,
					BlackChar: qrterminal.BLACK,
					WhiteChar: qrterminal.WHITE,
					QuietZone: 2,
				}) // inpost uses HIGH error correction
				code, err := qr.Encode(p.QrCode, *level)

				if err == nil {
					qrWidth = code.Size
				}
			}

			qrLines := strings.Split(qrBuf.String(), "\n")
			infoLines := strings.Split(infoBuf.String(), "\n")

			var infoWidth int

			for _, l := range infoLines {
				if len(l) > infoWidth {
					infoWidth = len(l)
				}
			}

			termWidth, _ := terminal.Width()

			if infoWidth+qrWidth+2 > int(termWidth) {
				// print vertically
				for _, line := range qrLines {
					fmt.Println(line)
				}
				for _, line := range infoLines {
					fmt.Println(line)
				}
			} else {
				for i := 0; i < len(qrLines) || i < len(infoLines); i++ {
					if i < len(qrLines) {
						fmt.Print(qrLines[i])
					}
					if i < len(qrLines) && i < len(infoLines) {
						fmt.Print("  ")
					}
					if i < len(infoLines) {
						fmt.Print(infoLines[i])
					}
					fmt.Print("\n")

				}
			}

		case "json":
			s, _ := json.MarshalIndent(p, "", "    ")
			fmt.Print(string(s))
		}

		return nil
	},
}
