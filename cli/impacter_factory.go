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
	if !strings.HasPrefix(opts.GetFiles()[0], "https://github.com/") {
		return false, nil
	}

	credentials := opts.GetCredentials()
	if len(credentials) == 0 {
		impacter := impact.NewGitHubPullRequestImpacter(opts.Files[0], "", "")

		return true, &impacter
	}

	usernamePassword := strings.SplitN(credentials, ":", 2)
	if len(usernamePassword) != 2 {
		panic(fmt.Errorf("Invalid credentials format, use <username:password>"))
	}

	username := usernamePassword[0]
	password := usernamePassword[1]
	impacter := impact.NewGitHubPullRequestImpacter(opts.Files[0], username, password)

	return true, &impacter
}
