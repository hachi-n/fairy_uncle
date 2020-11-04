package options

import "github.com/urfave/cli/v2"

func JsonFlag(destination *string, required bool) *cli.StringFlag {
	return &cli.StringFlag{
		Name:        "json",
		Usage:       "",
		Required:    required,
		Destination: destination,
	}
}

func KeyFlag(destination *string, required bool) *cli.StringFlag {
	return &cli.StringFlag{
		Name:        "key",
		Usage:       "",
		Required:    required,
		Destination: destination,
	}
}

