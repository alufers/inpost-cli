package main

import (
	"fmt"
	"math"
	"strings"

	"time"

	"github.com/alufers/inpost-cli/swagger"
	"github.com/antihax/optional"
	"github.com/urfave/cli/v2"
)

func prependSpaceIfNotEmpty(d string) string {
	if strings.TrimSpace(d) == "" {
		return ""
	}
	return " " + d
}

func resolveShipmentNumber(c *cli.Context, passedNumber string) (string, *swagger.APIClient, error) {
	passedNumber = strings.TrimSpace(passedNumber)
	if passedNumber == "" {
		return "", nil, fmt.Errorf("shipment number cannot be empty (you can pass only the beginning of the number if the package is yours)")
	}

	matchingPackages := []swagger.Parcel{}
	ua := "1970-01-01T00:00:00.001Z"
	limitedAccounts := GetLimitedAccounts(c)
	var matchingClient *swagger.APIClient

	if len(passedNumber) == 24 && len(limitedAccounts) == 1 {
		return passedNumber, limitedAccounts[0].ToClient(), nil
	}

	if len(limitedAccounts) == 0 {
		return "", nil, fmt.Errorf("no inpost accounts configured, please use `inpost-cli login` to configure one")
	}

	for _, acc := range limitedAccounts {
		apiClient := acc.ToClient()
		data, _, err := apiClient.DefaultApi.V1ParcelGet(c.Context, &swagger.DefaultApiV1ParcelGetOpts{
			UpdatedAfter: optional.NewString(ua),
		})
		if err != nil {
			return "", nil, fmt.Errorf("failed to request parcels (phone number: %v) to resolve shipment number: %w", acc.PhoneNumber, err)
		}
		// match prefixes
		for _, p := range data {
			if strings.HasPrefix(p.ShipmentNumber, passedNumber) {
				matchingPackages = append(matchingPackages, p)
				matchingClient = apiClient
			}
		}
	}

	// remove duplicates
	seen := make(map[string]bool)
	for _, p := range matchingPackages {
		if _, ok := seen[p.ShipmentNumber]; !ok {
			seen[p.ShipmentNumber] = true
		}
	}

	if len(matchingPackages) == 0 {
		return "", nil, fmt.Errorf("no package with matching shipment number of '%v' found in %d accounts", passedNumber, len(limitedAccounts))
	}
	if len(matchingPackages) > 1 {
		msg := "The passed package number '%v' is ambiguous, it matches: "
		for _, p := range matchingPackages {
			msg += p.ShipmentNumber + ", "
		}
		msg += "please provide more digits of the number"
		return "", nil, fmt.Errorf(msg, passedNumber)
	}
	return matchingPackages[0].ShipmentNumber, matchingClient, nil
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
