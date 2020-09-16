package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/alufers/inpost-cli/swagger"
	"github.com/antihax/optional"
	"github.com/shibukawa/configdir"
	"github.com/urfave/cli/v2"
)

var LoginCmd = &cli.Command{
	Name:        "login",
	Usage:       "log in to inpost via SMS code",
	Description: "Without any arguments it logs you in interactively. When only --phone-number is passed it sends a text with the code. When both --phone-number and --sms-code are passed it confirms the login.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "phone-number",
			Usage: "phone number with country code to send the SMS to",
		},
		&cli.StringFlag{
			Name:  "sms-code",
			Usage: "SMS code recieved from a previously sent message, use this options to log-in non-interactively",
		},
	},
	Action: func(c *cli.Context) error {
		var phoneNumber = c.String("phone-number")
		var smsCode = c.String("sms-code")

		phoneNumberPassed := true
		if phoneNumber == "" && smsCode == "" {
			phoneNumberPassed = false
			fmt.Print("Enter your phone number: ")
			reader := bufio.NewReader(os.Stdin)
			phoneNumber, _ = reader.ReadString('\n')
			phoneNumber = strings.TrimSuffix(phoneNumber, "\n")

		}
		apiClient := swagger.NewAPIClient(swagger.NewConfiguration())
		if phoneNumber != "" && smsCode == "" {
			_, resp, err := apiClient.DefaultApi.V1SendSMSCodePhoneNumberGet(c.Context, phoneNumber)
			if err != nil {
				log.Fatalf("failed to send SMS code: %v", err)
			}
			data, err := ioutil.ReadAll(resp.Body)
			log.Printf("SendSMSCode response: %v", string(data))
		}

		if smsCode == "" && !phoneNumberPassed {
			fmt.Print("Enter the sms code you recieved: ")
			reader := bufio.NewReader(os.Stdin)
			smsCode, _ = reader.ReadString('\n')
			smsCode = strings.TrimSuffix(smsCode, "\n")

		}
		if smsCode != "" {
			data, _, err := apiClient.DefaultApi.V1ConfirmSMSCodePhoneNumberSmsCodePost(c.Context, phoneNumber, smsCode, &swagger.DefaultApiV1ConfirmSMSCodePhoneNumberSmsCodePostOpts{
				Body: optional.NewInterface(swagger.PhoneOsRequest{
					PhoneOS: "Android",
				}),
			})
			if err != nil {
				log.Fatalf("failed to confirm sms code: %v", err)
			}
			log.Printf("ConfirmSMSCode resp data: %#v", data)
			folders := configDirs.QueryFolders(configdir.Global)
			log.Printf("Saving AuthToken and RefreshToken in %v", filepath.Join(folders[0].Path, "config.json"))
			marshaledJSON, err := json.Marshal(data)
			if err != nil {
				return err
			}
			folders[0].WriteFile("config.json", marshaledJSON)
		}
		return nil
	},
}
