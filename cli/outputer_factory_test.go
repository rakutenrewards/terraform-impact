package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/RakutenReady/terraform-impact/impact"
)

func TestCreateStdOutOutputer(t *testing.T) {
	result := createOutputer(validImpactOptions())

	assert.IsType(t, impact.StdOutOutputer{}, result, "Result should be of StdOutOutputer type")
}
