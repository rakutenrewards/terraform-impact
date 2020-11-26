package impact

import (
	"path/filepath"
)

// The `Impacter` interface allows flexibilty for other type
// of impacters to be added into the tool. An impacter lists
// files that can impact states. `ImpacterImpl` should always
// be the one used with the `Inner Impacter` being the one
// switched around.
type Impacter interface {
	List() []string
}

type ImpacterImpl struct {
	Inner Impacter
}

func (impacter ImpacterImpl) List() []string {
	var result []string
	for _, file := range impacter.Inner.List() {
		if file != "" {
			result = append(result, filepath.Clean(file))
			// ensures deleted files are taken into account
			result = append(result, filepath.Dir(file))
		}
	}

	return result
}

func NewImpacter(inner Impacter) ImpacterImpl {
	return ImpacterImpl{inner}
}

type CommandLineImpacter struct {
	Files []string
}

func NewCommandLineImpacter(files []string) CommandLineImpacter {
	return CommandLineImpacter{files}
}

func (impacter CommandLineImpacter) List() []string {
	return impacter.Files
}
