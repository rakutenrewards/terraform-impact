package testutils

import (
	"testing"

	_ "github.com/RakutenReady/terraform-impact/testutils/setup"
)

func TestStateFilesExist(t *testing.T) {
	assertExistAll(t, GetAwsStates())
	assertExistAll(t, GetGcpStates())
	assertExistAll(t, GetGcpCompanyStates())
	assertExistAll(t, GetStates())
}
