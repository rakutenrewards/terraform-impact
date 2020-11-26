package impact

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	tu "github.com/RakutenReady/terraform-impact/testutils"
	_ "github.com/RakutenReady/terraform-impact/testutils/setup"
	"github.com/RakutenReady/terraform-impact/trees"
)

type depsTestCase struct {
	StateDir             string
	WantRootDependencies []string
	WantAllDependencies  []string
}

func TestValidBuildStateDependenciesTree(t *testing.T) {
	testCases := []*depsTestCase{
		&depsTestCase{tu.AwsGatewayStateDir, tu.GetAwsGatewayStateRootDependencies(), tu.GetAwsGatewayStateDependencies()},
		&depsTestCase{tu.GcpCompanyStateDir, tu.GetGcpCompanyStateRootDependencies(), tu.GetGcpCompanyStateDependencies()},
		&depsTestCase{tu.GcpCompanyDatadogOnlyServiceStateDir, tu.GetGcpCompanyDatadogOnlyServiceStateRootDependencies(), tu.GetGcpCompanyDatadogOnlyServiceStateDependencies()},
		&depsTestCase{tu.GcpDatadogPgGoogleServiceStateDir, tu.GetGcpDatadogPgGoogleServiceStateRootDependencies(), tu.GetGcpDatadogPgGoogleServiceStateDependencies()},
		&depsTestCase{tu.GcpPgOnlyServiceStateDir, tu.GetGcpPgOnlyServiceStateRootDependencies(), tu.GetGcpPgOnlyServiceStateDependencies()},
	}

	assert := assert.New(t)
	for _, testCase := range testCases {
		visitedNodes := make(map[string]*trees.Node)

		// running each test case twice to test visitedNodes map
		for range []int{1, 2} {
			runDepsTestCase(assert, testCase, visitedNodes)
		}
	}
}

func TestFailingBuildStateDependenciesTree(t *testing.T) {
	testCases := []struct {
		Paths   []string
		WantErr string
	}{
		{tu.GetNeitherModulesNorStates(), "is not a Terraform state"},
		{[]string{tu.AwsPoorlyWrittenStateDir}, "is not a Terraform module"},
	}
	assert := assert.New(t)
	builder := StateDependenciesTreeBuilder{}
	visitedNodes := make(map[string]*trees.Node)
	for _, testCase := range testCases {
		for _, path := range testCase.Paths {
			_, err := builder.Build(path, visitedNodes)

			msg := fmt.Sprintf(`On: BuildStateDependenciesTree("%v", visitedNodes)`, path)
			assert.Errorf(err, msg)
			assert.Contains(err.Error(), testCase.WantErr, msg)
		}
	}
}

func runDepsTestCase(assert *assert.Assertions, testCase *depsTestCase, visitedNodes map[string]*trees.Node) {
	builder := StateDependenciesTreeBuilder{}
	node, err := builder.Build(testCase.StateDir, visitedNodes)

	baseMsg := fmt.Sprintf(`On: BuildStateDependenciesTree("%v", visitedNodes)`, testCase.StateDir)
	assert.Nilf(err, "%v [err] should be nil", baseMsg)

	assertRootNode(assert, node, testCase, baseMsg)
	assertAllDependencies(assert, node, testCase, baseMsg)
}

func assertRootNode(assert *assert.Assertions, root *trees.Node, testCase *depsTestCase, baseMsg string) {
	assert.Equalf(testCase.StateDir, root.Path, "%v Node.Path", baseMsg)

	assert.Lenf(root.Dependencies, len(testCase.WantRootDependencies), `%v len(root.Dependencies)`, baseMsg)
	for _, expectedRootDep := range testCase.WantRootDependencies {
		result := root.ContainsRootDependency(expectedRootDep)

		assert.Truef(result, `%v root.ContainsRootDependency("%v") should return true`, baseMsg, expectedRootDep)
	}
}

func assertAllDependencies(assert *assert.Assertions, root *trees.Node, testCase *depsTestCase, baseMsg string) {
	nbDeps := root.CountDependencies()
	assert.Equalf(len(testCase.WantAllDependencies), nbDeps, `%v root.CountDependencies()`, baseMsg)

	for _, expectedDep := range testCase.WantAllDependencies {
		result := root.ContainsDependency(expectedDep)

		assert.Truef(result, `%v root.ContainsDependency("%v") should return true`, baseMsg, expectedDep)
	}
}
