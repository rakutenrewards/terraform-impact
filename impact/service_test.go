package impact

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/RakutenReady/terraform-impact/trees"
)

type stateListerMock struct {
	mock.Mock
}

func (s *stateListerMock) List() []string {
	args := s.Called()

	return args.Get(0).([]string)
}

type depsBuilderMock struct {
	mock.Mock
}

func (d *depsBuilderMock) Build(moduleDir string, visitedNodes map[string]*trees.Node) (*trees.Node, error) {
	args := d.Called(moduleDir, visitedNodes)

	return args.Get(0).(*trees.Node), args.Error(1)
}

func TestImpact(t *testing.T) {
	root, rootA, rootB := makeTestNodes()
	stateLister := &stateListerMock{}
	builder := &depsBuilderMock{}

	stateLister.On("List").Return([]string{"root", "root-a", "root-b"})
	builder.On("Build", "root", mock.Anything).Return(root, nil)
	builder.On("Build", "root-a", mock.Anything).Return(rootA, nil)
	builder.On("Build", "root-b", mock.Anything).Return(rootB, nil)

	service := impactServiceImpl{stateLister, builder}

	testCases := []struct {
		Paths []string
		Want  []string
	}{
		{
			[]string{"nope", "nothing", "root-a-0"},
			[]string{"root", "root-a", "root-b"},
		},
		{
			[]string{"nope", "nothing", "root-a-2"},
			[]string{"root", "root-a", "root-b"},
		},
		{
			[]string{"nope", "root-b-0", "ardita"},
			[]string{"root", "root-b"},
		},
		{
			[]string{"nope", "root-b-1", "ardita"},
			[]string{"root", "root-b"},
		},
		{
			[]string{"nope", "root-a", "ardita"},
			[]string{"root"},
		},
		{
			[]string{"nope", "root-b", "ardita"},
			[]string{"root"},
		},
		{
			[]string{"nope", "ardita", "none"},
			[]string{},
		},
		{
			[]string{},
			[]string{},
		},
	}

	assert := assert.New(t)
	for _, testCase := range testCases {
		result, err := service.Impact(testCase.Paths)

		msg := onImpactMsg(testCase.Paths)
		assert.Nil(err, msg)
		assert.ElementsMatch(testCase.Want, result, msg)

		stateLister.AssertExpectations(t)
		builder.AssertExpectations(t)
	}
}

func TestImpactWithEmptyStateDirsList(t *testing.T) {
	stateLister := &stateListerMock{}
	builder := &depsBuilderMock{}

	stateLister.On("List").Return([]string{})
	builder.On("Build", mock.Anything, mock.Anything).Return(nil, nil)

	service := impactServiceImpl{stateLister, builder}

	args := []string{"whatever", "it does not matter"}
	result, err := service.Impact(args)

	assert.Nil(t, err, onImpactMsg(args))
	assert.Emptyf(t, result, onImpactMsg(args))

	stateLister.AssertExpectations(t)
	builder.AssertNumberOfCalls(t, "Build", 0)
}

func TestImpactWithEmptyDependenciesTrees(t *testing.T) {
	emptyNode := &trees.Node{
		Path:         "empty",
		Dependencies: []*trees.Node{},
	}

	stateLister := &stateListerMock{}
	builder := &depsBuilderMock{}

	stateLister.On("List").Return([]string{"root", "root-a", "root-b"})
	builder.On("Build", mock.Anything, mock.Anything).Return(emptyNode, nil)

	service := impactServiceImpl{stateLister, builder}

	args := []string{"whatever", "it does not matter"}
	result, err := service.Impact(args)

	assert.Nil(t, err, onImpactMsg(args))
	assert.Emptyf(t, result, onImpactMsg(args))

	stateLister.AssertExpectations(t)
	builder.AssertNumberOfCalls(t, "Build", 3)
	builder.AssertCalled(t, "Build", "root", mock.Anything)
	builder.AssertCalled(t, "Build", "root-a", mock.Anything)
	builder.AssertCalled(t, "Build", "root-b", mock.Anything)
}

func TestImpactWithErrorWhenBuildingDependenciesTrees(t *testing.T) {
	wantErrMsg := "Some error msg"

	stateLister := &stateListerMock{}
	builder := &depsBuilderMock{}

	stateLister.On("List").Return([]string{"root", "root-b"})
	builder.On("Build", mock.Anything, mock.Anything).Return((*trees.Node)(nil), fmt.Errorf(wantErrMsg))

	service := impactServiceImpl{stateLister, builder}

	args := []string{"whatever", "it does not matter"}
	result, err := service.Impact(args)

	assert.Nil(t, result, onImpactMsg(args))
	assert.EqualError(t, err, wantErrMsg, onImpactMsg(args))

	stateLister.AssertExpectations(t)
	builder.AssertNumberOfCalls(t, "Build", 1)
	builder.AssertCalled(t, "Build", "root", mock.Anything)
}

func onImpactMsg(args []string) string {
	return fmt.Sprintf("On impact(%v):", args)
}

func makeTestNodes() (root *trees.Node, rootA *trees.Node, rootB *trees.Node) {
	rootAdeps := make([]*trees.Node, 3)
	for i := 0; i < 3; i++ {
		rootAdeps[i] = &trees.Node{
			Path:         fmt.Sprintf("root-a-%d", i),
			Dependencies: []*trees.Node{},
		}
	}
	rootA = &trees.Node{
		Path:         "root-a",
		Dependencies: rootAdeps,
	}

	rootBdeps := make([]*trees.Node, 2)
	for i := 0; i < 2; i++ {
		rootBdeps[i] = &trees.Node{
			Path:         fmt.Sprintf("root-b-%d", i),
			Dependencies: []*trees.Node{},
		}
	}
	rootBdeps = append(rootBdeps, rootAdeps...)
	rootB = &trees.Node{
		Path:         "root-b",
		Dependencies: rootBdeps,
	}

	root = &trees.Node{
		Path: "root",
		Dependencies: []*trees.Node{
			rootA,
			rootB,
		},
	}

	return root, rootA, rootB
}
