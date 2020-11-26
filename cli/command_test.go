package cli

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/docopt/docopt.go"

	"github.com/RakutenReady/terraform-impact/impact"
)

type impacterMock struct {
	mock.Mock
}

func (m *impacterMock) List() []string {
	args := m.Called()

	return args.Get(0).([]string)
}

type impactServiceMock struct {
	mock.Mock
}

func (m *impactServiceMock) Impact(xs []string) (impactedStates []string, err error) {
	args := m.Called(xs)

	return args.Get(0).([]string), args.Error(1)
}

type impactOutputerMock struct {
	mock.Mock
}

func (m *impactOutputerMock) Output(xs []string) {
	m.Called(xs)
}

type impactFactoryMock struct {
	mock.Mock
}

func (m *impactFactoryMock) Create(opts impactOptions) (impact.Impacter, impact.ImpactService, impact.Outputer) {
	args := m.Called(opts)

	return args.Get(0).(impact.Impacter), args.Get(1).(impact.ImpactService), args.Get(2).(impact.Outputer)
}

func TestValidCommandRun(t *testing.T) {
	impacter, service, outputer, factory := makeMocks()

	impacterList := []string{"a", "b", "c"}
	impacter.On("List").Return(impacterList)

	serviceResult := []string{"d", "e", "f"}
	service.On("Impact", impacterList).Return(serviceResult, nil)
	outputer.On("Output", serviceResult)

	cmd := ImpactCommand{
		impactOptions{},
		factory,
	}

	args := []string{"file_1", "file_2"}
	opts, _ := docopt.ParseArgs(cmd.Usage(), args, "0.0.0")
	factory.On("Create", mock.Anything).Return(impacter, service, outputer)

	err := cmd.Run(opts)

	assert.Nil(t, err, "Error should be nil")
	factory.AssertExpectations(t)
	impacter.AssertExpectations(t)
	service.AssertExpectations(t)
	outputer.AssertExpectations(t)
}

func TestCommandRunWhenServiceReturnsError(t *testing.T) {
	wantErrMsg := "service failed"

	impacter, service, outputer, factory := makeMocks()
	impacterList := []string{"a", "b", "c"}
	impacter.On("List").Return(impacterList)

	service.On("Impact", impacterList).Return([]string{}, fmt.Errorf(wantErrMsg))

	cmd := ImpactCommand{
		impactOptions{},
		factory,
	}

	args := []string{"file_1", "file_2"}
	opts, _ := docopt.ParseArgs(cmd.Usage(), args, "0.0.0")
	factory.On("Create", mock.Anything).Return(impacter, service, outputer)

	err := cmd.Run(opts)

	assert.EqualError(t, err, wantErrMsg, "Should throw error when Service.Impact fails")
	factory.AssertExpectations(t)
	impacter.AssertExpectations(t)
	service.AssertExpectations(t)
	outputer.AssertExpectations(t)
}

func makeMocks() (*impacterMock, *impactServiceMock, *impactOutputerMock, *impactFactoryMock) {
	return &impacterMock{}, &impactServiceMock{}, &impactOutputerMock{}, &impactFactoryMock{}
}
