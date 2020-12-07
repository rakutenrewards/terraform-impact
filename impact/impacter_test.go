package impact

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type innerImpacterMock struct {
	mock.Mock
}

func (i *innerImpacterMock) List() ([]string, error) {
	args := i.Called()

	return args.Get(0).([]string), args.Error(1)
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
		inner.On("List").Return(testCase.InnerPaths, nil)
		impacter := NewImpacter(inner)

		result, err := impacter.List()

		msg := fmt.Sprintf(`On ImpacterImpl.List() with inner List() returning %v, nil`, testCase.InnerPaths)
		assert.Nil(err, msg)
		assert.ElementsMatch(testCase.Want, result, msg)
		inner.AssertExpectations(t)
	}
}

func TestImpacterWithInnerListError(t *testing.T) {
	inner := &innerImpacterMock{}
	wantErr := fmt.Errorf("Some error msg")
	inner.On("List").Return([]string{"something"}, wantErr)
	impacter := NewImpacter(inner)

	result, err := impacter.List()

	assert := assert.New(t)
	msg := `On ImpacterImpl.List() with inner List() returning ["something"], nil\n%v`
	assert.Nil(result, msg, "Expects resulting list to be nil")
	assert.EqualValues(wantErr, err, msg, "Expects returned error to be equal to inner error")
}
