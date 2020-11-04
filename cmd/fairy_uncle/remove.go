package main

import (
	"github.com/hachi-n/fairy_uncle"
	"github.com/hachi-n/fairy_uncle/cmd/fairy_uncle/internal/options"
	"github.com/urfave/cli/v2"
)

func destroyCommand() *cli.Command {
	var key string
	return &cli.Command{
		Name:        "remove-record",
		Usage:       "fairy_uncle remove-record --key ${NOTIFICATION_NAME}",
		Description: "fairy uncle remove-record settings.",
		Flags: []cli.Flag{
			options.KeyFlag(&key, true),
		},
		Action: func(c *cli.Context) error {
			return fairy_uncle.RemoveRecord(key)
		},
	}
}
