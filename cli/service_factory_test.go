package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/RakutenReady/terraform-impact/impact"
)

func TestCreateListOnlyImpactService(t *testing.T) {
	opts := validImpactOptions()
	opts.ListStates = true

	result := createService(opts)

	assert := assert.New(t)
	assert.IsType(impact.ListOnlyImpactService{}, result, "Result should be of ListOnlyImpactService type")
}

func TestCreateImpactServiceImpl(t *testing.T) {
	opts := validImpactOptions()
	opts.ListStates = false

	result := createService(opts)

	assert := assert.New(t)
	assert.IsType(impact.ImpactServiceImpl{}, result, "Result should be of ImpactServiceImpl type")
}
