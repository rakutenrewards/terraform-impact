package setup

import (
	"os"
	"path/filepath"
	"runtime"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)

	dir := filepath.Join(filepath.Dir(filename), "../..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}
