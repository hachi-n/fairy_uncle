package main

import (
	"github.com/hachi-n/fairy_uncle"
	"github.com/urfave/cli/v2"
)

func listCommand() *cli.Command {
	return &cli.Command{
		Name:        "list",
		Usage:       "fairy_uncle list",
		Description: "fairy uncle list settings.",
		Flags: []cli.Flag{
		},
		Action: func(c *cli.Context) error {
			return fairy_uncle.ListConfig()
		},
	}
}
