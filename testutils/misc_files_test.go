package testutils

import (
	"testing"

	_ "github.com/RakutenReady/terraform-impact/testutils/setup"
)

func TestNeitherModulesNorStateFilesExist(t *testing.T) {
	assertExistAll(t, GetNeitherModulesNorStates())
}

func TestExistentFiles(t *testing.T) {
	assertExistAll(t, GetExistentFiles())
}

func TestInexistence(t *testing.T) {
	assertDoesNotExistAll(t, GetInexistentPaths())
	assertDoesNotExistAll(t, GetInexistentDirs())
	assertDoesNotExistAll(t, GetInexistentFiles())
}
