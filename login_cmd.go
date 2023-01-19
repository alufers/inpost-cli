package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/alufers/inpost-cli/swagger"
	"github.com/antihax/optional"
	"github.com/urfave/cli/v2"
)

var LoginCmd = &cli.Command{
	Name:        "login",
	Usage:       "[--phone-number PHONE_NUMBER] [--sms-code PASSWORD] -- Log in to InPost, use without arguments for intaractive mode",
	Description: "Without any arguments it logs you in interactively. When only --phone-number is passed it sends a text with the code. When both --phone-number and --sms-code are passed it confirms the login.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "phone-number",
			Usage: "phone number with country code to send the SMS to",
		},
		&cli.StringFlag{
			Name:  "sms-code",
			Usage: "SMS code received from a previously sent message, use this options to log-in non-interactively",
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
			phoneNumber = strings.Replace(phoneNumber, "\r", "", -1)
			phoneNumber = strings.Replace(phoneNumber, "\n", "", -1)
		}
		apiClient := swagger.NewAPIClient(swagger.NewConfiguration())
		if phoneNumber != "" && smsCode == "" {
			_, resp, err := apiClient.DefaultApi.V1SendSMSCodePhoneNumberGet(c.Context, phoneNumber)
			if err != nil {
				log.Fatalf("failed to send SMS code: %v", err)
			}
			data, _ := ioutil.ReadAll(resp.Body)
			log.Printf("SendSMSCode response: %v", string(data))
		}

		if smsCode == "" && !phoneNumberPassed {
			fmt.Print("Enter the sms code you received: ")
			reader := bufio.NewReader(os.Stdin)
			smsCode, _ = reader.ReadString('\n')

			smsCode = strings.Replace(smsCode, "\r", "", -1)
			smsCode = strings.Replace(smsCode, "\n", "", -1)

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

			cfg := GetConfig(c)
			// remove entries with the same PhoneNumber from cfg.Accounts
			for i, acc := range cfg.Accounts {
				if acc.PhoneNumber == phoneNumber {
					cfg.Accounts = append(cfg.Accounts[:i], cfg.Accounts[i+1:]...)
					log.Printf("Removed existing account with phone number %s from config", phoneNumber)
				}
			}
			cfg.Accounts = append(cfg.Accounts, &ConfigAccount{
				PhoneNumber:  phoneNumber,
				AuthToken:    data.AuthToken,
				RefreshToken: data.RefreshToken,
			})
			return SaveConfig(c)
		}
		return nil
	},
}
