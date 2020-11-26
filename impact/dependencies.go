package impact

import (
	"fmt"
	"path/filepath"

	"github.com/RakutenReady/terraform-impact/tfparse"
	"github.com/RakutenReady/terraform-impact/trees"
	"github.com/RakutenReady/terraform-impact/utils"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

type DependenciesTreeBuilder interface {
	Build(stateDir string, visitedNodes map[string]*trees.Node) (*trees.Node, error)
}

type StateDependenciesTreeBuilder struct {
}

func NewStateDependenciesTreeBuilder() StateDependenciesTreeBuilder {
	return StateDependenciesTreeBuilder{}
}

// Build recursively looks for `module` block and its `source` attribute to
// add them to the provided state directory dependencies.
// Also adds files in the directory as dependencies and fills the visitedNodes
// map to accelerate further state node building.
// Returns the state directory (node, nil) or (nil, error) when any error arises.
// Note that symlink files are followed and all added as dependencies.
// Note that states are also viewed as modules by the `tfconfig` library.
func (builder StateDependenciesTreeBuilder) Build(stateDir string, visitedNodes map[string]*trees.Node) (*trees.Node, error) {
	if !tfparse.IsStateDir(stateDir) {
		return nil, fmt.Errorf("state dir [%v] is not a Terraform state", stateDir)
	}

	return buildModuleDependenciesTree(stateDir, visitedNodes)
}

func buildModuleDependenciesTree(moduleDir string, visitedNodes map[string]*trees.Node) (*trees.Node, error) {
	if nodeFromVisited, in := visitedNodes[moduleDir]; in {
		return nodeFromVisited, nil
	}

	subModulesDependencies, err := listModuleDependencies(moduleDir)
	if err != nil {
		return nil, err
	}

	selfNode := &trees.Node{
		Path:         moduleDir,
		Dependencies: []*trees.Node{},
	}
	dependencies := append(listFileDependencies(moduleDir), selfNode)
	for _, subModuleDependency := range subModulesDependencies {
		moduleDependencyNode, err := buildModuleDependenciesTree(subModuleDependency, visitedNodes)
		if err != nil {
			return nil, err
		}

		visitedNodes[moduleDependencyNode.Path] = moduleDependencyNode
		dependencies = append(dependencies, moduleDependencyNode)
	}

	return &trees.Node{
		Path:         moduleDir,
		Dependencies: dependencies,
	}, nil
}

func listModuleDependencies(moduleDir string) ([]string, error) {
	if !tfconfig.IsModuleDir(moduleDir) {
		return nil, fmt.Errorf("module dir [%v] is not a Terraform module", moduleDir)
	}

	module, diags := tfconfig.LoadModule(moduleDir)
	if diags.HasErrors() {
		return nil, fmt.Errorf("module dir [%v] failed to load.\nDiagnostics:\n%v", moduleDir, diags.Error())
	}

	depSources := make(map[string]bool)
	var deps []string
	for _, moduleCall := range module.ModuleCalls {
		moduleSourcePath := filepath.Join(moduleDir, moduleCall.Source)

		_, alreadyExists := depSources[moduleSourcePath]
		if !alreadyExists {
			depSources[moduleSourcePath] = true
			deps = append(deps, moduleSourcePath)
		}
	}

	return deps, nil
}

func listFileDependencies(moduleDir string) []*trees.Node {
	var nodes []*trees.Node
	for _, file := range utils.ListFilesIn(moduleDir) {
		filePath := filepath.Join(moduleDir, file)
		for _, symlinkPath := range utils.TraceSymlinkFile(filePath) {
			node := &trees.Node{
				Path:         symlinkPath,
				Dependencies: []*trees.Node{},
			}

			nodes = append(nodes, node)
		}
	}

	return nodes
}
