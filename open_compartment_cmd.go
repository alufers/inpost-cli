package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/alufers/inpost-cli/swagger"
	"github.com/antihax/optional"
	"github.com/urfave/cli/v2"
)

var OpenCompartmentCmd = &cli.Command{
	Name:        "open-compartment",
	Aliases:     []string{"open"},
	Usage:       "[--open-code OPEN_CODE] [--no-confirm] <SHIPMENT_NUMBER> -- remotely open compartment",
	Description: "Remotely opens a compartment in the pick-up machine.\n   Warning: There is no distance limit. Use with caution as you can get your package stolen this way.",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "no-confirm",
			Usage: "Disable confirmation of opening",
		},
		&cli.StringFlag{
			Name:  "open-code",
			Usage: "Manually set the open code. Must be passed when opening packages from somebody's else's phone number. You can omit this for your own packages.",
		},
	},
	Action: func(c *cli.Context) error {
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
			return fmt.Errorf("No such parcel found")
		}
		openCode := c.String("open-code")
		if openCode == "" {
			openCode = parcels[0].OpenCode
		}

		latitude := parcels[0].PickupPoint.Location.Latitude
		longitude := parcels[0].PickupPoint.Location.Longitude
		resp, _, err := apiClient.DefaultApi.V1CollectValidatePost(c.Context, &swagger.DefaultApiV1CollectValidatePostOpts{
			Body: optional.NewInterface(swagger.ValidationRequest{
				GeoPoint: &swagger.GeoPoint{
					Accuracy:  10.0,
					Latitude:  latitude,
					Longitude: longitude,
				},
				Parcel: &swagger.ParcelCompartment{
					OpenCode:       openCode,
					ShipmentNumber: shipmentNumber,
				},
			}),
		})
		if err != nil {
			return fmt.Errorf("failed to validate compartment: %v", err)
		}
		ctrlCSignal := make(chan os.Signal)
		signal.Notify(ctrlCSignal, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-ctrlCSignal
			log.Printf("\nCtrl-C pressed. Terminating open session...")
			_, _, err := apiClient.DefaultApi.V1CollectTerminateSessionUuidPost(
				c.Context,
				resp.SessionUuid,
			)
			if err != nil {
				log.Printf("failed to terminate compartment opening: %v", err)
			}
			os.Exit(0)
		}()
		if !c.Bool("no-confirm") {
			var pickupPointInfo = "<unknown>"
			if parcels[0].PickupPoint != nil {
				pickupPointInfo = fmt.Sprintf("%v%v, %v (%v)",
					parcels[0].PickupPoint.AddressDetails.Street,
					prependSpaceIfNotEmpty(parcels[0].PickupPoint.AddressDetails.BuildingNumber),
					parcels[0].PickupPoint.AddressDetails.City,
					parcels[0].PickupPoint.LocationDescription)
			}
			fmt.Println("Warning: You are about to remotely open a compartment in the pick-up machine (Paczkomat)")
			fmt.Println("Warning: Please double check that you are nearby the correct machine. Somebody can steal your package if you make a mistake!")
			fmt.Printf(
				"Press 'y' and enter to open compartment at %v for shipment number %v from %v: \n (n): ",
				pickupPointInfo,
				parcels[0].ShipmentNumber,
				parcels[0].SenderName,
			)

			reader := bufio.NewReader(os.Stdin)
			ra, _, _ := reader.ReadRune()
			if ra != 'y' {
				_, _, err := apiClient.DefaultApi.V1CollectTerminateSessionUuidPost(
					c.Context,
					resp.SessionUuid,
				)
				if err != nil {
					return fmt.Errorf("failed to terminate compartment opening: %w", err)
				}
				return fmt.Errorf("aborted, session terminated")
			}
		}

		openResp, _, err := apiClient.DefaultApi.V1CollectCompartmentOpenSessionUuidPost(c.Context, resp.SessionUuid)
		if err != nil {
			return fmt.Errorf("failed to open compartnment: %v", err)
		}
		log.Printf("CompartmentOpen resp: %#v", openResp)

		return nil
	},
}
