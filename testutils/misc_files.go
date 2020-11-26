package testutils

func GetNeitherModulesNorStates() []string {
	return []string{
		// aws
		"test_resources/terraform/aws",
		"test_resources/terraform/aws/states",
		"test_resources/terraform/aws/states/gateway/modules",
		"test_resources/terraform/aws/states/gateway/main.tf",
		"test_resources/terraform/aws/states/gateway/terraform.tfvars",
		// gcp
		"test_resources/terraform/gcp",
		"test_resources/terraform/gcp/global",
		"test_resources/terraform/gcp/modules",
		"test_resources/terraform/gcp/modules/datadog",
		"test_resources/terraform/gcp/modules/db",
		"test_resources/terraform/gcp/modules/db/pg/main.tf",
		"test_resources/terraform/gcp/states",
		"test_resources/terraform/gcp/states/company/datadog-only-service/main.tf",
		"test_resources/terraform/gcp/states/company/datadog-only-service/versions.tf",
		"test_resources/terraform/gcp/states/datadog-pg-google-service/secrets.vault",
		"test_resources/terraform/gcp/states/datadog-pg-google-service/terraform.tfvars",
		// others
		"test_resources/other",
		"test_resources/other/ansible/nothing.yml",
		"test_resources/other/ardita.json",
		"test_resources/terraform/docs/aws.md",
		"test_resources/terraform/docs/secrets",
		"test_resources/terraform/docs/secrets/jenkins.md",
	}
}

func GetInexistentPaths() []string {
	return []string{
		"test_resources/nope",
		"test_resources/nope/bob.json",
	}
}

func GetExistentFiles() []string {
	return []string{
		"test_resources/other/ansible/nothing.yml",
		"test_resources/other/ardita.json",
		"test_resources/terraform/docs/aws.md",
		"test_resources/terraform/docs/secrets/jenkins.md",
	}
}

func GetInexistentFiles() []string {
	return []string{
		"test_resources/nope/ansible/nothing.yml",
		"test_resources/nope/ardita.json",
	}
}

func GetInexistentDirs() []string {
	return []string{
		"test_resources/ardita",
		"test_resources/charles",
	}
}
