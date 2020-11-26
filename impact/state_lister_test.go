package impact

import (
	"testing"

	"github.com/stretchr/testify/assert"

	tu "github.com/RakutenReady/terraform-impact/testutils"
	_ "github.com/RakutenReady/terraform-impact/testutils/setup"
)

func TestMatchAllDiscoveryListStatesFrom(t *testing.T) {
	testCases := []struct {
		RootDirs []string
		Want     []string
	}{
		{
			[]string{tu.TestResourcesRootDir, tu.TerraformRootDir},
			tu.GetStates(),
		},
		{
			[]string{tu.AwsRootDir},
			tu.GetAwsStates(),
		},
		{
			[]string{tu.GcpRootDir, tu.GcpStatesDir},
			tu.GetGcpStates(),
		},
		{
			[]string{tu.GcpCompanyStateDir},
			tu.GetGcpCompanyStates(),
		},
		{
			[]string{tu.OtherRootDir, tu.TerraformDocsDir},
			[]string{},
		},
	}

	assert := assert.New(t)
	for _, testCase := range testCases {
		for _, rootDir := range testCase.RootDirs {
			lister := NewDiscoveryStateLister(rootDir, "")
			result := lister.List()

			assert.ElementsMatchf(testCase.Want, result, `DiscoveryStateLister("%v", "").List()`, rootDir)
		}
	}
}

func TestDiscoveryListStatesWithMatcherFrom(t *testing.T) {
	rootDir := tu.TestResourcesRootDir
	testCases := []struct {
		Substrings []string
		Want       []string
	}{
		{
			[]string{"test_resources/", "test_resources/terraform", "/terraform/", ""},
			tu.GetStates(),
		},
		{
			[]string{"/aws/", "/aws/states/"},
			tu.GetAwsStates(),
		},
		{
			[]string{"/gcp/", "/gcp/states/", "/terraform/gcp/"},
			tu.GetGcpStates(),
		},
		{
			[]string{"/company", "/gcp/states/company", "/company"},
			tu.GetGcpCompanyStates(),
		},
		{
			[]string{"company/datadog", "/datadog-only-service"},
			[]string{tu.GcpCompanyDatadogOnlyServiceStateDir},
		},
		{
			[]string{"test_resources/terraform/gcp/states/pg-only-service", "states/pg-only-service"},
			[]string{tu.GcpPgOnlyServiceStateDir},
		},
		{
			[]string{"other", "/docs/", "/does_not/exist/"},
			[]string{},
		},
	}

	assert := assert.New(t)
	for _, testCase := range testCases {
		for _, substring := range testCase.Substrings {
			lister := NewDiscoveryStateLister(rootDir, substring)
			result := lister.List()

			assert.ElementsMatchf(testCase.Want, result, `DiscoveryStateLister("%v", "%v").List()`, rootDir, substring)
		}
	}
}
