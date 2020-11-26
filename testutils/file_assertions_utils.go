package testutils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type assertable func(*testing.T, string)

func assertExists(t *testing.T, path string) {
	_, err := os.Stat(path)

	assert.Nilf(t, err, "Path [%v] should exist and does not", path)
}

func assertDoesNotExist(t *testing.T, path string) {
	_, err := os.Stat(path)

	assert.NotNilf(t, err, "Path [%v] should not exist and does", path)
}

func assertAll(t *testing.T, paths []string, assertion assertable) {
	for _, path := range paths {
		assertion(t, path)
	}
}

func assertExistAll(t *testing.T, paths []string) {
	assertAll(t, paths, assertExists)
}

func assertDoesNotExistAll(t *testing.T, paths []string) {
	assertAll(t, paths, assertDoesNotExist)
}
