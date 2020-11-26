package trees

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type containsTestCase struct {
	Node          *Node
	ContainsValue string
	Want          bool
}

func TestNodeContainsRootDependency(t *testing.T) {
	root, rootA, rootB := makeTestNodes()

	testCases := []containsTestCase{
		// root
		{root, "root", false},
		{root, "root/a", true},
		{root, "./root/a/../a/../a/../a", true},
		{root, "root/a/0", false},
		{root, "root/a/1", false},
		{root, "root/b", true},
		{root, "root/b/0", false},
		{root, "root/b/1", false},
		{root, "ardita", false},
		{root, "charlotte", false},
		// root/a
		{rootA, "root", false},
		{rootA, "root/a", false},
		{rootA, "root/a/0", true},
		{rootA, "root/a/1", true},
		{rootA, "root/a/2", true},
		{rootA, "root/b", false},
		{rootA, "root/b/0", false},
		{rootA, "root/b/1", false},
		{rootA, "ardita", false},
		{rootA, "charlotte", false},
		// root/b
		{rootB, "root", false},
		{rootB, "root/a", false},
		{rootB, "root/a/0", true},
		{rootB, "root/a/1", true},
		{rootB, "root/a/2", true},
		{rootB, "root/b", false},
		{rootB, "root/b/0", true},
		{rootB, "root/b/1", true},
		{rootB, "ardita", false},
		{rootB, "charlotte", false},
	}

	assert := assert.New(t)
	for _, testCase := range testCases {
		actual := testCase.Node.ContainsRootDependency(testCase.ContainsValue)

		assert.Equalf(testCase.Want, actual, `%v.ContainsRootDependency("%v") != %v`, testCase.Node.Path, testCase.ContainsValue, testCase.Want)
	}
}

func TestNodeContainsDependency(t *testing.T) {
	root, rootA, rootB := makeTestNodes()

	testCases := []containsTestCase{
		// root
		{root, "root", false},
		{root, "root/a", true},
		{root, "root/a/0", true},
		{root, "root/a/1", true},
		{root, "root/a/2", true},
		{root, "root/b", true},
		{root, "./root/a/1/../../b/0", true},
		{root, "root/b/0", true},
		{root, "root/b/1", true},
		{root, "ardita", false},
		{root, "charlotte", false},
		// root/a
		{rootA, "root", false},
		{rootA, "root/a", false},
		{rootA, "root/a/0", true},
		{rootA, "root/a/1", true},
		{rootA, "root/a/2", true},
		{rootA, "root/b", false},
		{rootA, "root/b/0", false},
		{rootA, "root/b/1", false},
		{rootA, "ardita", false},
		{rootA, "charlotte", false},
		// root/b
		{rootB, "root", false},
		{rootB, "root/a", false},
		{rootB, "root/a/0", true},
		{rootB, "root/a/1", true},
		{rootB, "root/a/2", true},
		{rootB, "root/b", false},
		{rootB, "root/b/0", true},
		{rootB, "root/b/1", true},
		{rootB, "ardita", false},
		{rootB, "charlotte", false},
	}

	assert := assert.New(t)
	for _, testCase := range testCases {
		actual := testCase.Node.ContainsDependency(testCase.ContainsValue)

		assert.Equalf(testCase.Want, actual, `%v.ContainsDependency("%v") != %v`, testCase.Node.Path, testCase.ContainsValue, testCase.Want)
	}
}

func TestNodeCountDependencies(t *testing.T) {
	root, rootA, rootB := makeTestNodes()

	testCases := []struct {
		Node *Node
		Want int
	}{
		{root, 10}, // rootA, rootB + rootAdeps + rootBdeps
		{rootA, 3},
		{rootB, 5},
	}

	assert := assert.New(t)
	for _, testCase := range testCases {
		result := testCase.Node.CountDependencies()

		assert.Equalf(testCase.Want, result, "%v.CountDependencies() != %v", testCase.Node.Path, testCase.Want)
	}
}

func TestNodeAnyDependency(t *testing.T) {
	root, rootA, rootB := makeTestNodes()

	testCases := []struct {
		node         *Node
		dependencies []string
		want         bool
	}{
		// root
		{root, []string{"./root/a/1/../../b/0"}, true},
		{root, []string{"root", "root/a", "root/b"}, true},
		{root, []string{"nope", "other/nope", "root/a/1"}, true},
		{root, []string{"nope", "other/nope", "root/b/0"}, true},
		{root, []string{"nope", "other/nope"}, false},
		{root, []string{}, false},
		// root/A
		{rootA, []string{"root/a", "root/b"}, false},
		{rootA, []string{"root/a/0"}, true},
		{rootA, []string{"nope", "other/nope", "root/a/1"}, true},
		{rootA, []string{"nope", "other/nope", "root/b/0"}, false},
		{rootA, []string{"nope", "other/nope"}, false},
		{rootA, []string{}, false},
		// root/B
		{rootB, []string{"root", "root/a", "root/b"}, false},
		{rootB, []string{"root/b/1"}, true},
		{rootB, []string{"nope", "other/nope", "root/a/1"}, true},
		{rootB, []string{"nope", "other/nope", "root/a/2"}, true},
		{rootB, []string{"nope", "other/nope"}, false},
		{rootB, []string{}, false},
	}

	assert := assert.New(t)
	for _, testCase := range testCases {
		result := testCase.node.AnyDependency(testCase.dependencies)

		assert.Equalf(testCase.want, result, "On: %v.AnyDependency(%v)", testCase.node.Path, testCase.dependencies)
	}
}

func makeTestNodes() (root *Node, rootA *Node, rootB *Node) {
	rootAdeps := make([]*Node, 3)
	for i := 0; i < 3; i++ {
		rootAdeps[i] = &Node{
			fmt.Sprintf("root/a/%d", i),
			[]*Node{},
		}
	}
	rootA = &Node{
		"root/a",
		rootAdeps,
	}

	rootBdeps := make([]*Node, 2)
	for i := 0; i < 2; i++ {
		rootBdeps[i] = &Node{
			fmt.Sprintf("root/b/%d", i),
			[]*Node{},
		}
	}
	rootBdeps = append(rootBdeps, rootAdeps...)
	rootB = &Node{
		"root/b",
		rootBdeps,
	}

	root = &Node{
		"root",
		[]*Node{
			rootA,
			rootB,
		},
	}

	return root, rootA, rootB
}
