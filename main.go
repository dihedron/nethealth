package main

import (
	"os"

	"github.com/dihedron/nethealth/cmd/nethealth/command"
	"github.com/jessevdk/go-flags"
)

func main() {
	options := command.Commands{}
	if _, err := flags.NewParser(&options, flags.Default).Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}

}
