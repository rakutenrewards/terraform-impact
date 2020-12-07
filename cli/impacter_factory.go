package cli

import (
	"fmt"
	"strings"

	"github.com/RakutenReady/terraform-impact/impact"
)

func createImpacter(opts ImpactOptions) impact.Impacter {
	return impact.NewImpacter(createInnerImpacter(opts))
}

func createInnerImpacter(opts ImpactOptions) impact.Impacter {
	if matches, gitHubImpacter := extractGitHubImpacter(opts); matches {
		return *gitHubImpacter
	}

	return impact.NewCommandLineImpacter(opts.Files)
}

func extractGitHubImpacter(opts ImpactOptions) (bool, *impact.GitHubPullRequestImpacter) {
	if !strings.HasPrefix(opts.Files[0], "https://github.com/") {
		return false, nil
	}

	if len(opts.Credentials) == 0 {
		impacter := impact.NewGitHubPullRequestImpacter(opts.Files[0], "", "")

		return true, &impacter
	}

	credentials := strings.SplitN(opts.Credentials, ":", 2)
	if len(credentials) != 2 {
		panic(fmt.Errorf("Invalid credentials format, use <username:password>"))
	}

	username := credentials[0]
	password := credentials[1]
	impacter := impact.NewGitHubPullRequestImpacter(opts.Files[0], username, password)

	return true, &impacter
}
