package main

import (
	"github.com/hachi-n/fairy_uncle"
	"github.com/urfave/cli/v2"
)

func notifyCommand() *cli.Command {
	return &cli.Command{
		Name:        "notify",
		Usage:       "fairy_uncle notify",
		Description: "fairy uncle notify.",
		Flags: []cli.Flag{
		},
		Action: func(c *cli.Context) error {
			return fairy_uncle.Notify()
		},
	}
}
