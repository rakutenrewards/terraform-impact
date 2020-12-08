package cli

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultValues(t *testing.T) {
	emptyOpts := ImpactOptions{}

	assert := assert.New(t)
	assert.Equal(".", emptyOpts.GetRootDir(), "RootDir should default to current directory")
	assert.Equal("", emptyOpts.GetPattern(), "Pattern should default to empty string")
	assert.Equal("", emptyOpts.GetOutput(), "Output should default to empty string")
}

func TestValues(t *testing.T) {
	opts := validImpactOptions()

	assert := assert.New(t)
	assert.Equal(opts.RootDir, opts.GetRootDir(), "RootDir should not default")
	assert.Equal(opts.Pattern, opts.GetPattern(), "Pattern should not default")
	assert.Equal(opts.Output, opts.GetOutput(), "Output should not default")
}

func TestGetCredentials(t *testing.T) {
	oldUsername := os.Getenv("GITHUB_USERNAME")
	oldPassword := os.Getenv("GITHUB_PASSWORD")
	defer os.Setenv("GITHUB_PASSWORD", oldPassword)
	defer os.Setenv("GITHUB_USERNAME", oldUsername)

	testCases := []struct {
		Credentials     string
		EnvUsername     string
		EnvPassword     string
		WantCredentials string
	}{
		{"username:password123", "not-used", "not-used", "username:password123"},
		{"username:password123", "", "not-used", "username:password123"},
		{"username:password123", "not-used", "", "username:password123"},
		{"username:password123", "", "", "username:password123"},
		{"", "env-username", "env-password123", "env-username:env-password123"},
		{"", "env-username", "", ""},
		{"", "", "env-password123", ""},
		{"", "", "", ""},
	}

	assert := assert.New(t)
	for _, testCase := range testCases {
		opts := ImpactOptions{
			Credentials: testCase.Credentials,
		}
		os.Setenv("GITHUB_USERNAME", testCase.EnvUsername)
		os.Setenv("GITHUB_PASSWORD", testCase.EnvPassword)

		result := opts.GetCredentials()

		assert.Equalf(testCase.WantCredentials, result, "With opts [%v], credentials are not matching", opts)
	}
}
