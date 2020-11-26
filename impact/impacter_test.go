package impact

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type innerImpacterMock struct {
	mock.Mock
}

func (i *innerImpacterMock) List() []string {
	args := i.Called()

	return args.Get(0).([]string)
}

func TestImpacterImplCleans(t *testing.T) {
	testCases := []struct {
		InnerPaths []string
		Want       []string
	}{
		{
			[]string{"./a/b/../../a/b/c/code.json", "", "/a/b/c/d", ""},
			[]string{"a/b/c/code.json", "a/b/c", "/a/b/c/d", "/a/b/c"},
		},
		{
			[]string{},
			[]string{},
		},
	}

	assert := assert.New(t)
	for _, testCase := range testCases {
		inner := &innerImpacterMock{}
		inner.On("List").Return(testCase.InnerPaths)
		impacter := NewImpacter(inner)

		result := impacter.List()

		assert.ElementsMatchf(testCase.Want, result, `On ImpacterImpl.List() with inner List() = %v`, testCase.InnerPaths)
		inner.AssertExpectations(t)
	}
}
