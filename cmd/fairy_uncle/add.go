package main

import (
	"github.com/hachi-n/fairy_uncle"
	"github.com/hachi-n/fairy_uncle/cmd/fairy_uncle/internal/options"
	"github.com/urfave/cli/v2"
)

func addCommand() *cli.Command {
	var jsonStr, key string
	return &cli.Command{
		Name:        "add-record",
		Usage:       "fairy_uncle add-record --json ${YOUR_JSON_PATH} --key ${NOTIFICATION_NAME}",
		Description: "fairy uncle is adding settings.",
		Flags: []cli.Flag{
			options.JsonFlag(&jsonStr, true),
			options.KeyFlag(&key, true),
		},
		Action: func(c *cli.Context) error {
			return fairy_uncle.AddRecord(key, jsonStr)
		},
	}
}
