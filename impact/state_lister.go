package impact

import (
	"path/filepath"
	"strings"

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
	RootDir   string
	Substring string
}

func NewDiscoveryStateLister(rootDir string, substring string) DiscoveryStateLister {
	return DiscoveryStateLister{
		rootDir,
		substring,
	}
}

// List recursively looks into the file system tree to find directories.
// Upon finding a directory that is a state, checks if the directory
// path contains the substring provided to the DiscoveryStateLister.
// Returns list of states directory path.
func (lister DiscoveryStateLister) List() []string {
	return lister.discoverStates(lister.RootDir)
}

func (lister *DiscoveryStateLister) discoverStates(candidateDir string) []string {
	var stateDirs []string
	if tfparse.IsStateDir(candidateDir) {
		if strings.Contains(candidateDir, lister.Substring) {
			stateDirs = append(stateDirs, candidateDir)
		}
	}

	for _, dirPath := range utils.ListDirsIn(candidateDir) {
		nextCandidatePath := filepath.Join(candidateDir, dirPath)

		stateDirs = append(stateDirs, lister.discoverStates(nextCandidatePath)...)
	}

	return stateDirs
}
