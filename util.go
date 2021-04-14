package main

import (
	"context"
	"fmt"
	"math"
	"strings"

	"time"

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
		return "", fmt.Errorf("no package with matching shipment number of '%v' found", passedNumber)
	}
	if len(matchingPackages) > 1 {
		msg := "The passed package number '%v' is ambiguous, it matches: "
		for _, p := range matchingPackages {
			msg += p.ShipmentNumber + ", "
		}
		msg += "please provide more digits of the number"
		return "", fmt.Errorf(msg, passedNumber)
	}
	return matchingPackages[0].ShipmentNumber, nil
}

// formatDuration formats a duration without the seconds
func formatDuration(d time.Duration) string {
	var format string
	seconds := math.Floor(d.Seconds())

	days := math.Floor(seconds / (60 * 60 * 24))

	if days > 0 {
		format += fmt.Sprintf("%vd ", days)
		seconds -= days * 60 * 60 * 24
	}

	hours := math.Floor(seconds / (60 * 60))

	if hours > 0 {
		format += fmt.Sprintf("%vh ", hours)
		seconds -= hours * 60 * 60
	}

	minutes := math.Floor(seconds / 60)

	if minutes > 0 {
		format += fmt.Sprintf("%vm ", minutes)
		seconds -= minutes * 60
	}

	if format == "" {
		format += fmt.Sprintf("%vs ", seconds)
	}

	return strings.TrimSpace(format)
}
