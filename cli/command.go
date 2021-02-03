package cli

type ImpactCommand struct {
	Factory impactFactory
}

func NewImpactCommand() ImpactCommand {
	return ImpactCommand{newImpactFactory()}
}

func (cmd ImpactCommand) Usage() string {
	return `Terraform Impact.

This tool takes a list of files as input and outputs a list of all the Terraform states
impacted by any of those files. An impact is described as a file creation, modification
or deletion.

Usage:
  impact <files>... [--rootdir <dir>] [--pattern <regexp>] [--user <credentials>]
	[--output <file>]
  impact -l | --list-states [--rootdir <dir>] [--pattern <regexp>] [--output <file>]
  impact -h | --help
  impact -v | --version

Arguments:
  <files>                  List of files that could impact any of the Terraform states.
                           When <files> is the url to a GitHub pull request, uses files
                           from the pull request.

Options:
  -r --rootdir <dir>       The directory from where the state discovery begins.
  -p --pattern <regexp>    A regex pattern to filter states. Only states whose path matches
                           the pattern will be taken into account.
  -u --user <credentials>  Credentials to access GitHub pull requests. Follows the curl
                           format 'username:password'. You can also pass credentials as
                           environment variables through: GITHUB_USERNAME and
                           GITHUB_PASSWORD. Note that the option always takes precendence
                           over environment variables.
  -o --output <file>       Outputs the result in the given file instead of using stdout.
                           Available formats: Json.
  -l --list-states         Lists the states discovered by the tool.
  -h --help                Show this screen.
  -v --version             Show version.`
}

func (cmd ImpactCommand) Run(opts ImpactOptions) error {
	impacter, service, outputer := cmd.Factory.Create(opts)

	impacterFiles, err := impacter.List()
	if err != nil {
		return err
	}

	result, err := service.Impact(impacterFiles)
	if err != nil {
		return err
	}

	return outputer.Output(result)
}
