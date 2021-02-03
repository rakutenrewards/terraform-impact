package impact

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Reusing methods and mocks from service_test.go
func TestImpactReturnsListFromStateLister(t *testing.T) {
	stateLister := &stateListerMock{}
	wantedResult := []string{"root", "root-a", "root-b"}

	stateLister.On("List").Return(wantedResult)

	service := ListOnlyImpactService{stateLister}

	testCases := []struct {
		Paths []string
	}{
		{
			[]string{"nope", "nothing", "root-a-0"},
		},
		{
			[]string{"nope", "ardita", "none"},
		},
		{
			[]string{},
		},
	}

	assert := assert.New(t)
	for _, testCase := range testCases {
		result, err := service.Impact(testCase.Paths)

		msg := onImpactMsg(testCase.Paths)
		assert.Nil(err, msg)
		assert.NotNil(result, msg)
		assert.ElementsMatch(wantedResult, result, msg)

		stateLister.AssertExpectations(t)
	}
}
