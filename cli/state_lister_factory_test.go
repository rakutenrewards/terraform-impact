package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/RakutenReady/terraform-impact/impact"
)

func TestCreateDiscoveryStateLister(t *testing.T) {
	result := createStateLister(validImpactOptions())

	assert := assert.New(t)
	assert.IsType(impact.DiscoveryStateLister{}, result, "Result should be of DiscoveryStateLister type")

	lister := result.(impact.DiscoveryStateLister)
	assert.Equal("RootDir", lister.RootDir)
	assert.Equal("Pattern", lister.Substring)
}
