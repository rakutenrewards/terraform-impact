package main

import (
	"fmt"
	"os"

	"github.com/RakutenReady/terraform-impact/cli"

	"github.com/docopt/docopt.go"
)

func main() {
	args := os.Args[1:]
	cmd := cli.NewImpactCommand()
	if len(args) == 0 {
		fmt.Println(cmd.Usage())
		os.Exit(1)
	}

	opts, parseErr := docopt.ParseArgs(cmd.Usage(), args, "0.0.0")
	if parseErr != nil {
		os.Exit(1)
	}

	runErr := cmd.Run(opts)
	if runErr != nil {
		fmt.Println(runErr)
		os.Exit(1)
	}
}
