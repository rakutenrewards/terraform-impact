package testutils

const (
	// aws
	AwsGatewayStateDir       = "test_resources/terraform/aws/states/gateway"
	AwsPoorlyWrittenStateDir = "test_resources/terraform/aws/states/poorly-written-state"
	// gcp
	GcpCompanyStateDir                   = "test_resources/terraform/gcp/states/company"
	GcpCompanyDatadogOnlyServiceStateDir = "test_resources/terraform/gcp/states/company/datadog-only-service"
	GcpDatadogPgGoogleServiceStateDir    = "test_resources/terraform/gcp/states/datadog-pg-google-service"
	GcpPgOnlyServiceStateDir             = "test_resources/terraform/gcp/states/pg-only-service"
)

func GetAwsStates() []string {
	return []string{
		AwsGatewayStateDir,
		AwsPoorlyWrittenStateDir,
	}
}

func GetGcpStates() []string {
	return []string{
		GcpCompanyStateDir,
		GcpCompanyDatadogOnlyServiceStateDir,
		GcpDatadogPgGoogleServiceStateDir,
		GcpPgOnlyServiceStateDir,
	}
}

func GetGcpCompanyStates() []string {
	return []string{
		GcpCompanyStateDir,
		GcpCompanyDatadogOnlyServiceStateDir,
	}
}

func GetStates() []string {
	s := GetGcpStates()

	return append(s, GetAwsStates()...)
}
