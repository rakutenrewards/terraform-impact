package testutils

const (
	// aws
	AwsGatewayDbModuleDir          = "test_resources/terraform/aws/states/gateway/modules/db"
	AwsGatewayDbPgModuleDir        = "test_resources/terraform/aws/states/gateway/modules/db/pg"
	AwsGatewayDbPgMonitorModuleDir = "test_resources/terraform/aws/states/gateway/modules/db/pg/monitor"
	// gcp
	GcpDatadogStandardMonitorModuleDir         = "test_resources/terraform/gcp/modules/datadog/standard_monitor"
	GcpDatadogInstanceGroupMonitorSetModuleDir = "test_resources/terraform/gcp/modules/datadog/instance_group_monitor_set"
	GcpDbPgModuleDir                           = "test_resources/terraform/gcp/modules/db/pg"
	GcpGoogleRuntimeConfiModuleDir             = "test_resources/terraform/gcp/modules/google/runtime_config"
	GcpUnusedModuleDir                         = "test_resources/terraform/gcp/modules/unused_module"
)

func GetModules() []string {
	return []string{
		AwsGatewayDbModuleDir,
		AwsGatewayDbPgModuleDir,
		AwsGatewayDbPgMonitorModuleDir,
		GcpDatadogStandardMonitorModuleDir,
		GcpDatadogInstanceGroupMonitorSetModuleDir,
		GcpDbPgModuleDir,
		GcpGoogleRuntimeConfiModuleDir,
		GcpUnusedModuleDir,
	}
}

// aws
func GetAwsDbModuleDependencies() []string {
	return []string{
		AwsGatewayDbModuleDir,
		"test_resources/terraform/aws/states/gateway/modules/db/main.tf",
		"test_resources/terraform/aws/states/gateway/modules/db/outputs.tf",
		"test_resources/terraform/aws/states/gateway/modules/db/variables.tf",
	}
}

func GetAwsDbPgModuleDependencies() []string {
	deps := []string{
		AwsGatewayDbPgModuleDir,
		"test_resources/terraform/aws/states/gateway/modules/db/pg/main.tf",
		"test_resources/terraform/aws/states/gateway/modules/db/pg/variables.tf",
		AwsGatewayDbModuleDir,
	}

	return append(deps, GetAwsDbModuleDependencies()...)
}

func GetAwsDbPgMonitorModuleDependencies() []string {
	return []string{
		AwsGatewayDbPgMonitorModuleDir,
		"test_resources/terraform/aws/states/gateway/modules/db/pg/monitor/main.tf",
	}
}

// gcp
func GetGcpDatadogStandardMonitorModuleDependencies() []string {
	return []string{
		GcpDatadogStandardMonitorModuleDir,
		"test_resources/terraform/gcp/modules/datadog/standard_monitor/main.tf",
		"test_resources/terraform/gcp/modules/datadog/standard_monitor/outputs.tf",
		"test_resources/terraform/gcp/modules/datadog/standard_monitor/variables.tf",
		"test_resources/terraform/gcp/modules/datadog/standard_monitor/versions.tf",
	}
}

func GetGcpDatadogInstanceGroupMonitorSetModuleDependencies() []string {
	deps := []string{
		GcpDatadogInstanceGroupMonitorSetModuleDir,
		"test_resources/terraform/gcp/modules/datadog/instance_group_monitor_set/main.tf",
		"test_resources/terraform/gcp/modules/datadog/instance_group_monitor_set/variables.tf",
		GcpDatadogStandardMonitorModuleDir,
	}

	return append(deps, GetGcpDatadogStandardMonitorModuleDependencies()...)
}

func GetGcpDbPgModuleDependencies() []string {
	deps := []string{
		GcpDbPgModuleDir,
		"test_resources/terraform/gcp/modules/db/pg/main.tf",
		"test_resources/terraform/gcp/modules/db/pg/outputs.tf",
		"test_resources/terraform/gcp/modules/db/pg/variables.tf",
		GcpDatadogStandardMonitorModuleDir,
	}

	return append(deps, GetGcpDatadogStandardMonitorModuleDependencies()...)
}

func GetGcpGoogleRuntimeConfigModuleDependencies() []string {
	return []string{
		GcpGoogleRuntimeConfiModuleDir,
		"test_resources/terraform/gcp/modules/google/runtime_config/main.tf",
		"test_resources/terraform/gcp/modules/google/runtime_config/variables.tf",
	}
}
