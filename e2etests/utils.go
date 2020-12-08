package e2etests

import (
	"os/exec"
	"testing"
)

type e2eArgsCase struct {
	Args []string
}

func execMain(extraArgs []string) *exec.Cmd {
	baseArgs := []string{"run", "main.go"}
	args := append(baseArgs, extraArgs...)

	return exec.Command("go", args...)
}

func runTest(t *testing.T, runnable func()) {
	if testing.Short() {
		t.Skip("Skipping e2e test")
	}

	runnable()
}

func getJsonFile() string {
	return "test_resources/json-test-output.json"
}

func getPullRequestUrl() string {
	return "https://github.com/RakutenReady/terraform-impact/pull/3"
}
