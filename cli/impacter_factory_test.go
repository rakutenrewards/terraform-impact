package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/RakutenReady/terraform-impact/impact"
)

func TestCreateCommandLineImpacter(t *testing.T) {
	result := createImpacter(validImpactOptions())

	assert := assert.New(t)
	assert.IsType(impact.ImpacterImpl{}, result, "Result should be of ImpacterImpl type")

	impacter := result.(impact.ImpacterImpl)
	assert.IsType(impact.CommandLineImpacter{}, impacter.Inner, "Impacter.Inner should be of CommandLineImpacter type")

	inner := impacter.Inner.(impact.CommandLineImpacter)
	assert.ElementsMatch([]string{"File_1", "File_2", "File_3"}, inner.Files)
}
