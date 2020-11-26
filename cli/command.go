package cli

import (
	"github.com/docopt/docopt.go"
)

type impactOptions struct {
	RootDir string   `docopt:"--rootdir"`
	Pattern string   `docopt:"--pattern"`
	Files   []string `docopt:"<files>"`
}

type ImpactCommand struct {
	Opts    impactOptions
	Factory impactFactory
}

func NewImpactCommand() ImpactCommand {
	return ImpactCommand{
		impactOptions{},
		newImpactFactory(),
	}
}

func (cmd ImpactCommand) Usage() string {
	return `Terraform Impact.

This tool takes a list of files as input and outputs a list of all the Terraform states
impacted by any of those files. An impact is described as a file creation, modification
or deletion.

Usage:
  impact [--pattern <string>] [--rootdir <dir>] <files>...
  impact -h | --help
  impact -v | --version

Options:
  -r --rootdir <dir>     The directory from where the state discovery begins [default: .]
  -p --pattern <string>  A string to filter states. Only states whose path contains the
                         string will be taken into account. [default: ]
  -h --help              Show this screen.
  -v --version           Show version.`
}

func (cmd ImpactCommand) Run(docOpts docopt.Opts) error {
	err := docOpts.Bind(&cmd.Opts)
	if err != nil {
		return err
	}

	impacter, service, outputer := cmd.Factory.Create(cmd.Opts)

	result, err := service.Impact(impacter.List())
	if err != nil {
		return err
	}

	outputer.Output(result)

	return nil
}
