package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name: "fairy_uncle",
		Usage: "fairy_uncle [sub commands] [flags]",
		Description: "",
		Commands: []*cli.Command{
			addCommand(),
			destroyCommand(),
			listCommand(),
			notifyCommand(),
			showCommand(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
