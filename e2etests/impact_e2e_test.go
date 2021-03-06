package e2etests

import (
	"fmt"
	"os"
	"testing"

	tu "github.com/RakutenReady/terraform-impact/testutils"
	_ "github.com/RakutenReady/terraform-impact/testutils/setup"
)

func TestImpactSucceedsWithExpectedOut(t *testing.T) {
	runTest(t, func() {
		testCases := []struct {
			Args []string
			Want []string
		}{
			// normal use cases
			{
				[]string{
					"test_resources/terraform/gcp/modules/db/pg/outputs.tf",
					"test_resources/terraform/gcp/modules/datadog/standard_monitor/main.tf",
					"test_resources/terraform/gcp/modules/google/runtime_config/variables.tf",
					"-r", tu.GcpRootDir,
				},
				[]string{
					tu.GcpCompanyDatadogOnlyServiceStateDir,
					tu.GcpDatadogPgGoogleServiceStateDir,
					tu.GcpPgOnlyServiceStateDir,
				},
			},
			{
				[]string{
					"test_resources/terraform/gcp/modules/db/pg/outputs.tf",
					"test_resources/terraform/gcp/modules/datadog/standard_monitor/main.tf",
					"test_resources/terraform/gcp/modules/google/runtime_config/variables.tf",
					fmt.Sprintf("--rootdir=%v", tu.GcpRootDir),
					fmt.Sprintf("--pattern=%v", tu.GcpCompanyStateDir),
				},
				[]string{tu.GcpCompanyDatadogOnlyServiceStateDir},
			},
			{
				[]string{
					"test_resources/terraform/gcp/modules/db/pg/outputs.tf",
					"-r", tu.GcpRootDir,
				},
				[]string{
					tu.GcpDatadogPgGoogleServiceStateDir,
					tu.GcpPgOnlyServiceStateDir,
				},
			},
			{
				[]string{
					"-r", tu.GcpCompanyStateDir,
					"-l",
				},
				tu.GetGcpCompanyStates(),
			},
			// symlink
			{
				[]string{
					"test_resources/terraform/gcp/global/terraform.tf",
					"-r", tu.GcpRootDir,
				},
				[]string{tu.GcpDatadogPgGoogleServiceStateDir},
			},
			// no result
			{
				[]string{
					"test_resources/terraform/gcp/modules/datadog/standard_monitor/main.tf",
					"-r", tu.GcpRootDir,
					"-p", "states/path/that/does/not/exist",
				},
				[]string{},
			},
			{
				[]string{
					"not_existing",
					"other/ansible/ardita.json",
					"-r", tu.GcpRootDir,
				},
				[]string{},
			},
			// no result using github PR impacter
			{
				[]string{
					getPullRequestUrl(),
					"-r", tu.GcpRootDir,
					"-u", fmt.Sprintf("%v:%v", os.Getenv("GITHUB_USERNAME"), os.Getenv("GITHUB_PASSWORD")),
				},
				[]string{},
			},
			// unused module
			{
				[]string{
					"test_resources/terraform/gcp/modules/unused_module/output.tf",
					"-r", tu.GcpRootDir,
				},
				[]string{},
			},
		}

		for _, testCase := range testCases {
			cmd := execMain(testCase.Args)

			out, err := cmd.CombinedOutput()

			assertNoErrors(t, err, testCase.Args)
			assertOutIsImpactedStates(t, out, testCase.Want, testCase.Args)
		}
	})
}

func TestImpactOutputersSucceeds(t *testing.T) {
	runTest(t, func() {
		testCases := []struct {
			Args []string
			Want []string
		}{
			// normal use cases
			{
				[]string{
					"test_resources/terraform/gcp/modules/db/pg/outputs.tf",
					"test_resources/terraform/gcp/modules/datadog/standard_monitor/main.tf",
					"test_resources/terraform/gcp/modules/google/runtime_config/variables.tf",
					"-r", tu.GcpRootDir,
					fmt.Sprintf("--output=%v", getJsonFile()),
				},
				[]string{
					tu.GcpCompanyDatadogOnlyServiceStateDir,
					tu.GcpDatadogPgGoogleServiceStateDir,
					tu.GcpPgOnlyServiceStateDir,
				},
			},
			{
				[]string{
					"-r", tu.GcpRootDir,
					"-l",
					fmt.Sprintf("--output=%v", getJsonFile()),
				},
				tu.GetGcpStates(),
			},
			// unused module for empty list
			{
				[]string{
					"test_resources/terraform/gcp/modules/unused_module/output.tf",
					"-r", tu.GcpRootDir,
					fmt.Sprintf("--output=%v", getJsonFile()),
				},
				[]string{},
			},
		}

		for _, testCase := range testCases {
			cmd := execMain(testCase.Args)

			out, err := cmd.CombinedOutput()

			assertNoErrors(t, err, testCase.Args)
			assertOutIsEmpty(t, out, testCase.Args)
			assertJsonIsImpactedStates(t, testCase.Want, testCase.Args)
		}
	})
}

func TestImpactFailsContainsErrMsg(t *testing.T) {
	runTest(t, func() {
		testCases := []struct {
			Args               []string
			WantContainsErrMsg string
		}{
			// failing because of unparseable modules
			{
				[]string{"some_file"},
				"test_resources/terraform/aws/states/poorly-written-state/modules/bob",
			},
			{
				[]string{
					"whatever-file",
					"-r", tu.TerraformRootDir,
					"-p", "aws/states/poorly-written-state",
				},
				"test_resources/terraform/aws/states/poorly-written-state/modules/bob",
			},
			// failing because of wrong creds to access PR
			{
				[]string{
					getPullRequestUrl(),
					"-r", tu.GcpRootDir,
					"-u", "user-65e17355-7fcc-4a83-8d25-8ce5d6064c2b:pwd123",
				},
				fmt.Sprintf("PR with link [%v] returned status [404]", getPullRequestUrl()),
			},
			// failing because passing files and -l option
			{
				[]string{
					"should/fail/because/files/are/note/allowed/with/l/opt",
					"-l",
					"-r", tu.GcpRootDir,
				},
				fmt.Sprintf(" "),
			},
		}

		for _, testCase := range testCases {
			cmd := execMain(testCase.Args)

			out, err := cmd.CombinedOutput()

			assertExitErrorCodeIs1(t, err, testCase.Args)
			assertOutContainsErrorMsg(t, out, testCase.WantContainsErrMsg, testCase.Args)
		}
	})
}
