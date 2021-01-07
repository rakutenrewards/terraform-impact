package impact

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"
)

// Only valid for v3 api version
const githubPerPageMaximumValue = 100
const githubRecommandedAcceptHeader = "application/vnd.github.v3+json"

type GitHubPullRequestImpacter struct {
	Url      string
	PerPage  int
	Username string
	Password string
}

func NewGitHubPullRequestImpacter(url string, username string, password string) GitHubPullRequestImpacter {
	return GitHubPullRequestImpacter{url, githubPerPageMaximumValue, username, password}
}

type gitHubFileResponse struct {
	Filename string `json:"filename"`
}

// Lists pull requests files (Limit of 3000 from GitHub API)
func (impacter GitHubPullRequestImpacter) List() ([]string, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	apiBaseUrl, err := buildGitHubBaseUrl(impacter.Url)
	if err != nil {
		return nil, err
	}

	filenames := []string{}
	page := 1
	for {
		apiEndpoint := impacter.buildGitHubPaginatedEndpoint(apiBaseUrl, page)
		request, err := impacter.newGetRequest(apiEndpoint)
		if err != nil {
			return nil, err
		}

		responseFilenames, err := impacter.executeRequest(client, request)
		if err != nil {
			return nil, err
		}
		if len(responseFilenames) == 0 {
			break
		}

		filenames = append(filenames, responseFilenames...)
		page += 1
	}

	return filenames, nil
}

func (impacter *GitHubPullRequestImpacter) buildGitHubPaginatedEndpoint(baseUrl string, page int) string {
	return fmt.Sprintf(baseUrl+"?per_page=%v&page=%v", impacter.PerPage, page)
}

func (impacter *GitHubPullRequestImpacter) newGetRequest(endpoint string) (*http.Request, error) {
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Accept", githubRecommandedAcceptHeader)
	if len(impacter.Username) > 0 && len(impacter.Password) > 0 {
		request.SetBasicAuth(impacter.Username, impacter.Password)
	}

	return request, nil
}

func (impacter *GitHubPullRequestImpacter) executeRequest(client *http.Client, request *http.Request) ([]string, error) {
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

func buildGitHubBaseUrl(url string) (string, error) {
	re := regexp.MustCompile(`^https://github\.com/(?P<owner>.+)/(?P<repository>.+)/pull/(?P<prnumber>\d+)$`)

	matches := re.FindStringSubmatch(url)
	if len(matches) != 4 {
		return "", fmt.Errorf("Url [%v] does not match github PR url pattern", url)
	}

	owner := matches[1]
	repository := matches[2]
	prnumber := matches[3]

	return fmt.Sprintf("https://api.github.com/repos/%v/%v/pulls/%v/files", owner, repository, prnumber), nil
}
