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
	List() ([]string, error)
}

type ImpacterImpl struct {
	Inner Impacter
}

func (impacter ImpacterImpl) List() ([]string, error) {
	innerList, innerErr := impacter.Inner.List()
	if innerErr != nil {
		return nil, innerErr
	}

	var result []string
	for _, file := range innerList {
		if file != "" {
			result = append(result, filepath.Clean(file))
			// ensures deleted files are taken into account
			result = append(result, filepath.Dir(file))
		}
	}

	return result, nil
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

func (impacter CommandLineImpacter) List() ([]string, error) {
	return impacter.Files, nil
}
