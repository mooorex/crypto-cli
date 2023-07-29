package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mooorex/crypto-cli/cmd/crypto"
	"github.com/urfave/cli/v2"
)

func init() {
	err := initApp()
	if err != nil {
		log.Fatal(err)
	}
}

func initApp() error {
	app := cli.NewApp()
	app.Name = "crypto-cli"
	app.Usage = "Secruely transfer important data using ecies."
	app.Version = "0.0.1"
	app.Compiled = time.Now()
	app.Action = func(c *cli.Context) error {
		fmt.Println(app.Name, app.Version)
		fmt.Println(app.Usage)
		return nil
	}
	app.Commands = crypto.CryptoCommands

	return app.Run(os.Args)
}
