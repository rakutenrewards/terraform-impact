package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/RakutenReady/terraform-impact/impact"
)

func TestCreateStdOutOutputer(t *testing.T) {
	opts := validImpactOptions()
	opts.Output = ""

	result := createOutputer(opts)

	assert.IsType(t, impact.StdOutOutputer{}, result, "Result should be of StdOutOutputer type")
}

func TestJsonOutputer(t *testing.T) {
	opts := validImpactOptions()
	opts.Output = "some-file.json"

	result := createOutputer(opts)

	assert := assert.New(t)
	assert.IsType(impact.JsonOutputer{}, result, "Result should be of JsonOutputer type")

	jsonOutputer := result.(impact.JsonOutputer)
	assert.Equal(opts.Output, jsonOutputer.FilePath, "Output file path should match")
}

func TestUnsupportedOutputer(t *testing.T) {
	wantErrMsg := "Unknown output format [bob.txt]"
	opts := validImpactOptions()
	opts.Output = "bob.txt"

	shouldPanicFn := func() {
		createOutputer(opts)
	}

	assert.PanicsWithError(t, wantErrMsg, shouldPanicFn, "Should panic when called with [bob.txt]")
}
