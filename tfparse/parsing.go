package tfparse

import (
	"path/filepath"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
)

// IsStateDir looks for a `main.tf` file in the provided directory.
// If it finds one, checks if that file contains a `backend` block
// within a `terraform` block.
// Returns true if it does, false otherwise.
func IsStateDir(dirPath string) bool {
	candidatePath := filepath.Join(dirPath, "main.tf")
	parser := hclparse.NewParser()
	parsedCandidate, err := parser.ParseHCLFile(candidatePath)
	if err != nil {
		return false
	}

	return isTerraformStateFile(parsedCandidate)
}

func isTerraformStateFile(file *hcl.File) bool {
	content, _, diags := file.Body.PartialContent(rootSchema)
	if diags != nil {
		return false
	}

	for _, block := range content.Blocks {
		if block.Type == "terraform" {
			tfContent, _, tfDiags := block.Body.PartialContent(stateSchema)
			if tfDiags != nil {
				continue
			}

			for _, tfBlock := range tfContent.Blocks {
				if tfBlock.Type == "backend" {
					return true
				}
			}
		}
	}

	return false
}
