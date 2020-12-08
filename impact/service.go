package impact

import (
	"github.com/RakutenReady/terraform-impact/trees"
)

// The goal of `ImpactService` interface is to ease tests.
// There should not be any other implementation than `impactServiceImpl`.
type ImpactService interface {
	Impact(paths []string) (impactedStates []string, err error)
}

type impactServiceImpl struct {
	lister      StateLister
	depsBuilder DependenciesTreeBuilder
}

func NewImpactService(lister StateLister) ImpactService {
	return impactServiceImpl{
		lister, NewStateDependenciesTreeBuilder(),
	}
}

// Impact lists the state directories and maps each state directory to a
// Node with its dependencies. Then, filters in state nodes
// which contains any of the [paths] provided as function parameter.
// Finally, returns the list of impacted state directories as (states, nil).
// If any error arises during the process, return (nil, error).
func (service impactServiceImpl) Impact(paths []string) ([]string, error) {
	stateDirs := service.lister.List()

	var stateNodes []*trees.Node
	visitedNodes := make(map[string]*trees.Node)
	for _, stateDir := range stateDirs {
		stateNode, err := service.depsBuilder.Build(stateDir, visitedNodes)
		if err != nil {
			return nil, err
		}

		stateNodes = append(stateNodes, stateNode)
	}

	impactedStates := []string{}
	for _, stateNode := range stateNodes {
		if stateNode.AnyDependency(paths) {
			impactedStates = append(impactedStates, stateNode.Path)
		}
	}

	return impactedStates, nil
}
