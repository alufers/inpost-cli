package main

import (
	"log"
	"os"

	"github.com/shibukawa/configdir"
	"github.com/urfave/cli/v2"
)

var configDirs = configdir.New("alufers", "inpost-cli")

func main() {
	app := &cli.App{
		Usage:       "A CLI utility replacing the InPost app",
		Description: "This utility allows you to list your packages, generate QR codes in the terminal to collect them and remotely open compartments. \n   Warning: There is no limit to how far away you can open the compartments from. Use with caution as you can get your package stolen this way.",
		Version:     InpostCliVersion,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				TakesFile:   true,
				Usage:       "Config file location",
				DefaultText: configDirs.QueryFolders(configdir.Global)[0].Path + "/config.json",
			},
			&cli.StringSliceFlag{
				Name:    "phone-number",
				Aliases: []string{"p"},
				Usage:   "Phone number to use, can be used multiple times. Uses all phone numbers by default.",
			},
		},
		Commands: []*cli.Command{
			LoginCmd,
			ListParcelsCmd,
			ParcelInfoCmd,
			OpenCompartmentCmd,
			TerminateOpenCompartment,
			RefreshTokenCmd,
			TokenInfoCmd,
			TrackCmd,
			ListReturnsCmd,
		},
		Before: LoadConfig,
	}
	app.EnableBashCompletion = true
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
