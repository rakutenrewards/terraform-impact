package impact

import (
	"fmt"
	"strings"
)

// The `Outputer` interface allows flexibilty for other type
// of outputers to be added into the tool.
type Outputer interface {
	Output([]string) error
}

type StdOutOutputer struct{}

func NewStdOutOutputer() StdOutOutputer {
	return StdOutOutputer{}
}

func (out StdOutOutputer) Output(results []string) error {
	fmt.Print(strings.Join(results, "\n"))

	return nil
}
