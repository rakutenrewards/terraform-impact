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

	docOpts, parseErr := docopt.ParseArgs(cmd.Usage(), args, "1.0.0")
	if parseErr != nil {
		os.Exit(1)
	}

	impactOpts := cli.ImpactOptions{}
	bindErr := docOpts.Bind(&impactOpts)
	if bindErr != nil {
		fmt.Println(bindErr)
		os.Exit(1)
	}

	runErr := cmd.Run(impactOpts)
	if runErr != nil {
		fmt.Println(runErr)
		os.Exit(1)
	}
}
