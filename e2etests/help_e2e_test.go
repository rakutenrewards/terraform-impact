package e2etests

import (
	"testing"

	_ "github.com/RakutenReady/terraform-impact/testutils/setup"
)

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
