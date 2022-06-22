package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

// TODO: finish
var LogoutCmd = &cli.Command{
	Name:        "logout",
	Usage:       "[--all] [PHONE NUMBER] -- logout from InPost",
	Description: "Without any arguments it asks you what account you want to remove.",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "all",
			Usage: "removes all accounts from the configuration file",
		},
	},
	Action: func(c *cli.Context) error {
		cfg := GetConfig(c)
		if c.Bool("all") {
			log.Printf("Logging out %d accounts", len(cfg.Accounts))
			cfg.Accounts = []*ConfigAccount{}

			return SaveConfig(c)
		}
		phoneNumber := c.Args().First()
		if phoneNumber == "" {
			fmt.Print("Enter the phone number you want to remove: ")
			reader := bufio.NewReader(os.Stdin)
			phoneNumber, _ = reader.ReadString('\n')
			phoneNumber = strings.Replace(phoneNumber, "\r", "", -1)
			phoneNumber = strings.Replace(phoneNumber, "\n", "", -1)
		}
		for i, account := range cfg.Accounts {
			if account.PhoneNumber == phoneNumber {
				cfg.Accounts = append(cfg.Accounts[:i], cfg.Accounts[i+1:]...)
				break
			}
		}

		return SaveConfig(c)
	},
}
