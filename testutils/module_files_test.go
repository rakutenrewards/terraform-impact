package testutils

import (
	"testing"

	_ "github.com/RakutenReady/terraform-impact/testutils/setup"
)

func TestModulesFilesExist(t *testing.T) {
	assertExistAll(t, GetModules())
}

func TestAwsModulesDependenciesExist(t *testing.T) {
	assertExistAll(t, GetAwsDbModuleDependencies())
	assertExistAll(t, GetAwsDbPgModuleDependencies())
	assertExistAll(t, GetAwsDbPgMonitorModuleDependencies())
}

func TestGcpModulesDependenciesExist(t *testing.T) {
	assertExistAll(t, GetGcpDatadogStandardMonitorModuleDependencies())
	assertExistAll(t, GetGcpDatadogInstanceGroupMonitorSetModuleDependencies())
	assertExistAll(t, GetGcpDbPgModuleDependencies())
	assertExistAll(t, GetGcpGoogleRuntimeConfigModuleDependencies())
}
