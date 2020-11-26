package testutils

import (
	"testing"

	_ "github.com/RakutenReady/terraform-impact/testutils/setup"
)

func TestStateRootDependenciesExist(t *testing.T) {
	assertExistAll(t, GetAwsGatewayStateRootDependencies())
	assertExistAll(t, GetGcpCompanyStateRootDependencies())
	assertExistAll(t, GetGcpCompanyDatadogOnlyServiceStateRootDependencies())
	assertExistAll(t, GetGcpDatadogPgGoogleServiceStateRootDependencies())
	assertExistAll(t, GetGcpPgOnlyServiceStateRootDependencies())
}

func TestStateDependenciesExist(t *testing.T) {
	assertExistAll(t, GetAwsGatewayStateDependencies())
	assertExistAll(t, GetGcpCompanyStateDependencies())
	assertExistAll(t, GetGcpCompanyDatadogOnlyServiceStateDependencies())
	assertExistAll(t, GetGcpDatadogPgGoogleServiceStateDependencies())
	assertExistAll(t, GetGcpPgOnlyServiceStateDependencies())
}
