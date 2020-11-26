package testutils

import (
	"testing"

	_ "github.com/RakutenReady/terraform-impact/testutils/setup"
)

func TestConstFilesExist(t *testing.T) {
	all := []string{
		TestResourcesRootDir,
		TerraformRootDir,
		TerraformDocsDir,
		AwsRootDir,
		GcpRootDir,
		GcpStatesDir,
		OtherRootDir,
	}

	assertExistAll(t, all)
}
