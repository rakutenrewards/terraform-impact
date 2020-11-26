package e2etests

import (
	"testing"

	tu "github.com/RakutenReady/terraform-impact/testutils"
	_ "github.com/RakutenReady/terraform-impact/testutils/setup"
)

func TestProgramSucceedsWithExpectedOut(t *testing.T) {
	runTest(t, func() {
		testCases := []struct {
			Args []string
			Want []string
		}{
			// normal use cases
			{
				[]string{
					"-r", tu.GcpRootDir,
					"test_resources/terraform/gcp/modules/db/pg/outputs.tf",
					"test_resources/terraform/gcp/modules/datadog/standard_monitor/main.tf",
					"test_resources/terraform/gcp/modules/google/runtime_config/variables.tf",
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
					"-p", tu.GcpCompanyStateDir,
					"test_resources/terraform/gcp/modules/db/pg/outputs.tf",
					"test_resources/terraform/gcp/modules/datadog/standard_monitor/main.tf",
					"test_resources/terraform/gcp/modules/google/runtime_config/variables.tf",
				},
				[]string{tu.GcpCompanyDatadogOnlyServiceStateDir},
			},
			{
				[]string{
					"-r", tu.GcpRootDir,
					"test_resources/terraform/gcp/modules/db/pg/outputs.tf",
				},
				[]string{
					tu.GcpDatadogPgGoogleServiceStateDir,
					tu.GcpPgOnlyServiceStateDir,
				},
			},
			// symlink
			{
				[]string{
					"-r", tu.GcpRootDir,
					"test_resources/terraform/gcp/global/terraform.tf",
				},
				[]string{tu.GcpDatadogPgGoogleServiceStateDir},
			},
			// no result
			{
				[]string{
					"-r", tu.GcpRootDir,
					"-p", "states/path/that/does/not/exist",
					"test_resources/terraform/gcp/modules/datadog/standard_monitor/main.tf",
				},
				[]string{},
			},
			{
				[]string{
					"-r", tu.GcpRootDir,
					"not_existing",
					"other/ansible/ardita.json",
				},
				[]string{},
			},
			// unused module
			{
				[]string{
					"-r", tu.GcpRootDir,
					"test_resources/terraform/gcp/modules/unused_module/output.tf",
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

func TestProgramFailsBecauseOfWronglyWrittenTerraform(t *testing.T) {
	runTest(t, func() {
		testCases := []struct {
			Args       []string
			WantErrMsg string
		}{
			// failing because of inexistent modules
			{
				[]string{"some_file"},
				"test_resources/terraform/aws/states/poorly-written-state/modules/bob",
			},
			{
				[]string{
					"-r", tu.TerraformRootDir,
					"-p", "aws/states/poorly-written-state",
					"whatever",
				},
				"test_resources/terraform/aws/states/poorly-written-state/modules/bob",
			},
		}

		for _, testCase := range testCases {
			cmd := execMain(testCase.Args)

			out, err := cmd.CombinedOutput()

			assertExitErrorCodeIs1(t, err, testCase.Args)
			assertOutContainsErrorMsg(t, out, testCase.WantErrMsg, testCase.Args)
		}
	})
}

func TestProgramSucceedsWithHelpMessage(t *testing.T) {
	runTest(t, func() {
		testCases := []e2eArgsCase{
			{[]string{"-h"}},
			{[]string{"--help"}},
			{[]string{"bob", "ok", "--help", "bye"}},
			{[]string{"bob", "ok", "-h", "bye"}},
		}

		for _, testCase := range testCases {
			cmd := execMain(testCase.Args)

			out, err := cmd.CombinedOutput()

			assertOutIsHelp(t, out, testCase.Args)
			assertNoErrors(t, err, testCase.Args)
		}
	})
}

func TestProgramFailsWithHelpMessage(t *testing.T) {
	runTest(t, func() {
		testCases := []e2eArgsCase{
			{[]string{"-j", "-e"}},
			{[]string{}},
		}

		for _, testCase := range testCases {
			cmd := execMain(testCase.Args)

			out, err := cmd.CombinedOutput()

			assertOutIsHelp(t, out, testCase.Args)
			assertExitErrorCodeIs1(t, err, testCase.Args)
		}
	})
}
