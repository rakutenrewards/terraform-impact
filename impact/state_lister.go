package impact

import (
	"path/filepath"
	"regexp"

	"github.com/RakutenReady/terraform-impact/tfparse"
	"github.com/RakutenReady/terraform-impact/utils"
)

// The `StateLister` interface allows flexibilty for other type
// of state lister to be added into the tool. The default one is
// `DiscoveryStateLister`.
type StateLister interface {
	List() []string
}

type DiscoveryStateLister struct {
	RootDir string
	Regexp  string
}

func NewDiscoveryStateLister(rootDir string, regexp string) DiscoveryStateLister {
	return DiscoveryStateLister{
		rootDir,
		regexp,
	}
}

// List recursively looks into the file system tree to find directories.
// Upon finding a directory that is a state, checks if the directory
// path matches the provided regexp.
// Returns list of states directory path.
func (lister DiscoveryStateLister) List() []string {
	pathRegex := regexp.MustCompile(lister.Regexp)

	return lister.discoverStates(lister.RootDir, pathRegex)
}

func (lister *DiscoveryStateLister) discoverStates(candidateDir string, pathRegex *regexp.Regexp) []string {
	var stateDirs []string
	if tfparse.IsStateDir(candidateDir) {
		if pathRegex.MatchString(candidateDir) {
			stateDirs = append(stateDirs, candidateDir)
		}
	}

	for _, dirPath := range utils.ListDirsIn(candidateDir) {
		nextCandidatePath := filepath.Join(candidateDir, dirPath)

		stateDirs = append(stateDirs, lister.discoverStates(nextCandidatePath, pathRegex)...)
	}

	return stateDirs
}
