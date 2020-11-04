package main

import (
	"github.com/hachi-n/fairy_uncle"
	"github.com/hachi-n/fairy_uncle/cmd/fairy_uncle/internal/options"
	"github.com/urfave/cli/v2"
)

func showCommand() *cli.Command {
	var key string
	return &cli.Command{
		Name:        "show-record",
		Usage:       "fairy_uncle show-record --key ${NOTIFICATION_NAME}",
		Description: "fairy uncle show-record settings.",
		Flags: []cli.Flag{
			options.KeyFlag(&key, true),
		},
		Action: func(c *cli.Context) error {
			return fairy_uncle.ShowRecord(key)
		},
	}
}
