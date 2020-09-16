package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/alufers/inpost-cli/swagger"
	"github.com/antihax/optional"
)

func prependSpaceIfNotEmpty(d string) string {
	if strings.TrimSpace(d) == "" {
		return ""
	}
	return " " + d
}

func resolveShipmentNumber(ctx context.Context, apiClient *swagger.APIClient, passedNumber string) (string, error) {
	passedNumber = strings.TrimSpace(passedNumber)
	if passedNumber == "" {
		return "", fmt.Errorf("shipment number cannot be empty (you can pass only the beginning of the number if the package is yours)")
	}
	if len(passedNumber) == 24 {
		return passedNumber, nil
	}

	ua := "1970-01-01T00:00:00.001Z"
	data, _, err := apiClient.DefaultApi.V1ParcelGet(ctx, &swagger.DefaultApiV1ParcelGetOpts{
		UpdatedAfter: optional.NewString(ua),
	})
	if err != nil {
		return "", fmt.Errorf("failed to request parcels to resolve shipment number: %w", err)
	}

	matchingPackages := []swagger.Parcel{}

	for _, p := range data {
		if strings.HasPrefix(p.ShipmentNumber, passedNumber) {
			matchingPackages = append(matchingPackages, p)
		}
	}
	if len(matchingPackages) == 0 {
		return "", fmt.Errorf("no package with matching shipmend number of '%v' found", passedNumber)
	}
	if len(matchingPackages) > 1 {
		msg := "The passed package number '%v' is ambigous, it matches: "
		for _, p := range matchingPackages {
			msg += p.ShipmentNumber + ", "
		}
		msg += "please provide more digits of the number"
		return "", fmt.Errorf(msg, passedNumber)
	}
	return matchingPackages[0].ShipmentNumber, nil
}
