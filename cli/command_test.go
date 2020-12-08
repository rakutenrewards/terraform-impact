package cli

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/RakutenReady/terraform-impact/impact"
)

type impacterMock struct {
	mock.Mock
}

func (m *impacterMock) List() ([]string, error) {
	args := m.Called()

	return args.Get(0).([]string), args.Error(1)
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

func (m *impactOutputerMock) Output(xs []string) error {
	args := m.Called(xs)

	return args.Error(0)
}

type impactFactoryMock struct {
	mock.Mock
}

func (m *impactFactoryMock) Create(opts ImpactOptions) (impact.Impacter, impact.ImpactService, impact.Outputer) {
	args := m.Called(opts)

	return args.Get(0).(impact.Impacter), args.Get(1).(impact.ImpactService), args.Get(2).(impact.Outputer)
}

func TestValidCommandRun(t *testing.T) {
	opts := validImpactOptions()
	impacter, service, outputer, factory := makeMocks()
	factory.On("Create", opts).Return(impacter, service, outputer)

	impacterList := []string{"a", "b", "c"}
	impacter.On("List").Return(impacterList, nil)

	serviceResult := []string{"d", "e", "f"}
	service.On("Impact", impacterList).Return(serviceResult, nil)
	outputer.On("Output", serviceResult).Return(nil)

	cmd := ImpactCommand{factory}

	err := cmd.Run(opts)

	assert.Nil(t, err, "Error should be nil")
	factory.AssertExpectations(t)
	impacter.AssertExpectations(t)
	service.AssertExpectations(t)
	outputer.AssertExpectations(t)
}

func TestCommandRunWhenImpacterListReturnsError(t *testing.T) {
	wantErr := fmt.Errorf("impacter failed")

	opts := validImpactOptions()
	impacter, service, outputer, factory := makeMocks()
	factory.On("Create", opts).Return(impacter, service, outputer)
	impacter.On("List").Return([]string{}, wantErr)

	cmd := ImpactCommand{factory}

	err := cmd.Run(opts)

	assert.EqualValues(t, err, wantErr, "Should return Impacter.List error")
	factory.AssertExpectations(t)
	impacter.AssertExpectations(t)
	service.AssertExpectations(t)
	outputer.AssertExpectations(t)
}

func TestCommandRunWhenServiceReturnsError(t *testing.T) {
	wantErr := fmt.Errorf("service failed")

	opts := validImpactOptions()
	impacter, service, outputer, factory := makeMocks()
	factory.On("Create", opts).Return(impacter, service, outputer)

	impacterList := []string{"a", "b", "c"}
	impacter.On("List").Return(impacterList, nil)
	service.On("Impact", impacterList).Return([]string{}, wantErr)

	cmd := ImpactCommand{factory}

	err := cmd.Run(opts)

	assert.EqualValues(t, err, wantErr, "Should return Service.Impact error")
	factory.AssertExpectations(t)
	impacter.AssertExpectations(t)
	service.AssertExpectations(t)
	outputer.AssertExpectations(t)
}

func TestCommandRunWhenOutputerReturnsError(t *testing.T) {
	wantErr := fmt.Errorf("outputer failed")

	opts := validImpactOptions()
	impacter, service, outputer, factory := makeMocks()
	factory.On("Create", opts).Return(impacter, service, outputer)

	impacterList := []string{"a", "b", "c"}
	impacter.On("List").Return(impacterList, nil)

	serviceResult := []string{"d", "e", "f"}
	service.On("Impact", impacterList).Return(serviceResult, nil)
	outputer.On("Output", serviceResult).Return(wantErr)

	cmd := ImpactCommand{factory}

	err := cmd.Run(opts)

	assert.EqualValues(t, err, wantErr, "Should return Outputer.Output error")
	factory.AssertExpectations(t)
	impacter.AssertExpectations(t)
	service.AssertExpectations(t)
	outputer.AssertExpectations(t)
}

func makeMocks() (*impacterMock, *impactServiceMock, *impactOutputerMock, *impactFactoryMock) {
	return &impacterMock{}, &impactServiceMock{}, &impactOutputerMock{}, &impactFactoryMock{}
}
