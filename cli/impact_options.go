package cli

import (
	"fmt"
	"os"
)

type ImpactOptions struct {
	Files       []string `docopt:"<files>"`
	RootDir     string   `docopt:"--rootdir"`
	Pattern     string   `docopt:"--pattern"`
	Credentials string   `docopt:"--user"`
	Output      string   `docopt:"--output"`
	ListStates  bool     `docopt:"--list-states"`
}

func (opt *ImpactOptions) GetRootDir() string {
	return getOrDefault(opt.RootDir, ".")
}

func (opt *ImpactOptions) GetPattern() string {
	return getOrDefault(opt.Pattern, "")
}

func (opt *ImpactOptions) GetOutput() string {
	return getOrDefault(opt.Output, "")
}

func (opt *ImpactOptions) GetCredentials() string {
	if len(opt.Credentials) > 0 {
		return opt.Credentials
	}

	envUsername := os.Getenv("GITHUB_USERNAME")
	envPassword := os.Getenv("GITHUB_PASSWORD")
	if len(envUsername) > 0 && len(envPassword) > 0 {
		return fmt.Sprintf("%v:%v", envUsername, envPassword)
	}

	return ""
}

func (opt *ImpactOptions) GetFiles() []string {
	return opt.Files
}

func getOrDefault(value string, def string) string {
	if len(value) > 0 {
		return value
	}

	return def
}
