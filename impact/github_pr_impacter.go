package impact

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"
)

type GitHubPullRequestImpacter struct {
	Url      string
	Username string
	Password string
}

func NewGitHubPullRequestImpacter(url string, username string, password string) GitHubPullRequestImpacter {
	return GitHubPullRequestImpacter{url, username, password}
}

type gitHubFileResponse struct {
	Filename string `json:"filename"`
}

// Lists pull requests files (Limit of 3000 from GitHub API)
func (impacter GitHubPullRequestImpacter) List() ([]string, error) {
	apiEndpoint, err := buildGitHubEndpoint(impacter.Url)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if len(impacter.Username) > 0 && len(impacter.Password) > 0 {
		request.SetBasicAuth(impacter.Username, impacter.Password)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("PR with link [%v] returned status [%v]", impacter.Url, response.StatusCode)
	}

	defer response.Body.Close()
	var filesResponse []gitHubFileResponse
	err = json.NewDecoder(response.Body).Decode(&filesResponse)
	if err != nil {
		return nil, err
	}

	filenames := []string{}
	for _, fileResponse := range filesResponse {
		filenames = append(filenames, fileResponse.Filename)
	}

	return filenames, nil
}

func buildGitHubEndpoint(url string) (string, error) {
	re := regexp.MustCompile(`^https:\/\/github\.com\/(?P<owner>.+)\/(?P<repository>.+)\/pull\/(?P<prnumber>\d+)$`)

	matches := re.FindStringSubmatch(url)
	if len(matches) != 4 {
		return "", fmt.Errorf("Url [%v] does not match github PR url pattern", url)
	}

	owner := matches[1]
	repository := matches[2]
	prnumber := matches[3]

	return fmt.Sprintf("https://api.github.com/repos/%v/%v/pulls/%v/files", owner, repository, prnumber), nil
}
