package trees

import (
	"github.com/RakutenReady/terraform-impact/utils"
)

type Node struct {
	Path         string
	Dependencies []*Node
}

func (node *Node) CountDependencies() int {
	total := len(node.Dependencies)
	for _, n := range node.Dependencies {
		total += n.CountDependencies()
	}

	return total
}

func (node *Node) ContainsRootDependency(dependency string) bool {
	for _, n := range node.Dependencies {
		if utils.SamePath(n.Path, dependency) {
			return true
		}
	}

	return false
}

func (node *Node) AnyDependency(dependencies []string) bool {
	for _, dependency := range dependencies {
		if node.ContainsDependency(dependency) {
			return true
		}
	}

	return false
}

func (node *Node) ContainsDependency(dependency string) bool {
	for _, n := range node.Dependencies {
		if utils.SamePath(n.Path, dependency) {
			return true
		}

		if n.ContainsDependency(dependency) {
			return true
		}
	}

	return false
}
