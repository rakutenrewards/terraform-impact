package testutils

// aws
// root dependencies are the dependency nodes directly under the state dir root node
func GetAwsGatewayStateRootDependencies() []string {
	return []string{
		AwsGatewayStateDir,
		"test_resources/terraform/aws/states/gateway/gw.vault",
		"test_resources/terraform/aws/states/gateway/main.tf",
		"test_resources/terraform/aws/states/gateway/outputs.tf",
		"test_resources/terraform/aws/states/gateway/terraform.tf",
		"test_resources/terraform/aws/states/gateway/terraform.tfvars",
		AwsGatewayDbPgModuleDir,
		AwsGatewayDbPgMonitorModuleDir,
	}
}

func GetAwsGatewayStateDependencies() []string {
	deps := GetAwsGatewayStateRootDependencies()

	deps = append(deps, GetAwsDbPgModuleDependencies()...)
	deps = append(deps, GetAwsDbPgMonitorModuleDependencies()...)

	return deps
}

// gcp
func GetGcpCompanyStateRootDependencies() []string {
	return []string{
		GcpCompanyStateDir,
		"test_resources/terraform/gcp/states/company/main.tf",
		"test_resources/terraform/gcp/states/company/versions.tf",
	}
}

func GetGcpCompanyStateDependencies() []string {
	return GetGcpCompanyStateRootDependencies()
}

func GetGcpCompanyDatadogOnlyServiceStateRootDependencies() []string {
	return []string{
		GcpCompanyDatadogOnlyServiceStateDir,
		"test_resources/terraform/gcp/states/company/datadog-only-service/main.tf",
		"test_resources/terraform/gcp/states/company/datadog-only-service/versions.tf",
		GcpDatadogStandardMonitorModuleDir,
	}
}

func GetGcpCompanyDatadogOnlyServiceStateDependencies() []string {
	deps := GetGcpCompanyDatadogOnlyServiceStateRootDependencies()

	return append(deps, GetGcpDatadogStandardMonitorModuleDependencies()...)
}

func GetGcpDatadogPgGoogleServiceStateRootDependencies() []string {
	return []string{
		GcpDatadogPgGoogleServiceStateDir,
		"test_resources/terraform/gcp/states/datadog-pg-google-service/main.tf",
		"test_resources/terraform/gcp/states/datadog-pg-google-service/others.tf",
		"test_resources/terraform/gcp/states/datadog-pg-google-service/secrets.vault",
		"test_resources/terraform/gcp/states/datadog-pg-google-service/services.tf",
		"test_resources/terraform/gcp/states/datadog-pg-google-service/terraform.tf",
		"test_resources/terraform/gcp/states/datadog-pg-google-service/terraform.tfvars",
		"test_resources/terraform/gcp/states/global_vars/terraform.tf",     // symlink
		"test_resources/terraform/gcp/states/global_vars/terraform.tfvars", // symlink
		"test_resources/terraform/gcp/global/terraform.tf",                 // symlink
		"test_resources/terraform/gcp/global/terraform.tfvars",             // symlink
		"test_resources/terraform/gcp/states/datadog-pg-google-service/versions.tf",
		GcpGoogleRuntimeConfiModuleDir,
		GcpDatadogInstanceGroupMonitorSetModuleDir,
		GcpDbPgModuleDir,
	}
}

func GetGcpDatadogPgGoogleServiceStateDependencies() []string {
	deps := GetGcpDatadogPgGoogleServiceStateRootDependencies()
	deps = append(deps, GetGcpGoogleRuntimeConfigModuleDependencies()...)
	deps = append(deps, GetGcpDatadogInstanceGroupMonitorSetModuleDependencies()...)
	deps = append(deps, GetGcpDbPgModuleDependencies()...)

	return deps
}

func GetGcpPgOnlyServiceStateRootDependencies() []string {
	return []string{
		GcpPgOnlyServiceStateDir,
		"test_resources/terraform/gcp/states/pg-only-service/main.tf",
		"test_resources/terraform/gcp/states/pg-only-service/random.txt",
		"test_resources/terraform/gcp/states/pg-only-service/services.tf",
		"test_resources/terraform/gcp/states/pg-only-service/terraform.tf",
		"test_resources/terraform/gcp/states/pg-only-service/terraform.tfvars",
		"test_resources/terraform/gcp/states/pg-only-service/versions.tf",
		GcpDbPgModuleDir,
	}
}

func GetGcpPgOnlyServiceStateDependencies() []string {
	deps := GetGcpPgOnlyServiceStateRootDependencies()

	return append(deps, GetGcpDbPgModuleDependencies()...)
}
