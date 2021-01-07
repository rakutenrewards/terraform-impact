package e2etests

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/RakutenReady/terraform-impact/impact"
)

func TestGitHubPullRequestImpacterWithAuth(t *testing.T) {
	runTest(t, func() {
		want := []string{
			".gitignore",
			"added/1.vault",
			"added/2.xml",
			"modified/3.tf",
			"deleted/2.java",
			"deleted/5.json",
		}

		impacter := getValidGitHubPrImpacter()

		result, err := impacter.List()

		assert := assert.New(t)
		assert.Nil(err, "GitHub PR impacter should return nil error")
		assert.ElementsMatch(want, result)
	})
}

func TestGitHubPullRequestImpacterWithInvalidAuth(t *testing.T) {
	runTest(t, func() {
		url := getPullRequestUrl()
		impacter := impact.NewGitHubPullRequestImpacter(url, "user-65e17355-7fcc-4a83-8d25-8ce5d6064c2b", "pwd123")
		wantErrMsg := fmt.Sprintf("PR with link [%v] returned status [404]", url)

		result, err := impacter.List()

		assert := assert.New(t)
		assert.Nil(result, "On invalid auth, list should be nil")
		assert.EqualError(err, wantErrMsg, "On invalid auth, should return error")
	})
}

func TestGitHubPullRequestImpacterWithInvalidUrl(t *testing.T) {
	runTest(t, func() {
		invalidUrl := "https://invalid-github.com/RakutenReady/terraform-impact/pull/3"
		wantErrMsg := fmt.Sprintf("Url [%v] does not match github PR url pattern", invalidUrl)
		impacter := impact.NewGitHubPullRequestImpacter(invalidUrl, "user", "pwd")

		result, err := impacter.List()

		assert := assert.New(t)
		assert.Nil(result, "On invalid url, list should be nil")
		assert.NotNil(err, "On invalid url, should return error")
		assert.EqualError(err, wantErrMsg, "On invalid url, should return error with message")
	})
}

func getValidGitHubPrImpacter() impact.Impacter {
	username := os.Getenv("GITHUB_USERNAME")
	password := os.Getenv("GITHUB_PASSWORD")

	return impact.GitHubPullRequestImpacter{
		Url:      getPullRequestUrl(),
		PerPage:  5,
		Username: username,
		Password: password,
	}
}
