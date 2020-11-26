package tfparse

import (
	"github.com/stretchr/testify/assert"
	"testing"

	tu "github.com/RakutenReady/terraform-impact/testutils"
	_ "github.com/RakutenReady/terraform-impact/testutils/setup"
)

func TestIsStateDirTrue(t *testing.T) {
	assert := assert.New(t)
	for _, state := range tu.GetStates() {
		result := IsStateDir(state)

		assert.Truef(result, "\nOn: IsStateDir(\"%v\")\nActual: %v\nWant: true", state, result)
	}
}

func TestIsStateDirFalse(t *testing.T) {
	all := append(tu.GetModules(), tu.GetNeitherModulesNorStates()...)
	all = append(all, tu.GetInexistentPaths()...)

	assert := assert.New(t)
	for _, path := range all {
		result := IsStateDir(path)

		assert.Falsef(result, "\nOn: IsStateDir(\"%v\")\nActual: %v\nWant: false", path, result)
	}
}
