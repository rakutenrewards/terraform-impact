package e2etests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertOutIsHelp(t *testing.T, output []byte, args []string) {
	assert.Containsf(t, string(output), "Usage:", "Run with args %v should output help message", args)
}

func assertOutIsEmpty(t *testing.T, output []byte, args []string) {
	assert.Emptyf(t, string(output), "Run with args %v", args)
}

func assertOutContainsErrorMsg(t *testing.T, out []byte, wantErrMsg string, args []string) {
	assert.Containsf(t, string(out), wantErrMsg, "Run with args %v should contains error message", args)
}

func assertOutIsImpactedStates(t *testing.T, out []byte, want []string, args []string) {
	wantStr := strings.Join(want, "\n")

	assert.Equalf(t, wantStr, string(out), "Run with args %v", args)
}

func assertJsonIsImpactedStates(t *testing.T, want []string, args []string) {
	var data struct {
		States []string `json:"states"`
	}

	bytes, _ := ioutil.ReadFile(getJsonFile())
	json.Unmarshal(bytes, &data)

	assert.ElementsMatchf(t, data.States, want, "Json content should match when running with args %v", args)
}

func assertNoErrors(t *testing.T, err error, args []string) {
	assert.Nilf(t, err, "Run with args %v should not return an error", args)
}

func assertExitErrorCodeIs1(t *testing.T, err error, args []string) {
	assert.Error(t, err)

	msg := fmt.Sprintf("Run with args %v should be an ExitError 1", args)
	assert.IsType(t, &exec.ExitError{}, err, msg)

	exitErr := err.(*exec.ExitError)
	assert.Equal(t, 1, exitErr.ExitCode(), msg)
}
