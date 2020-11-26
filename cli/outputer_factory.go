package cli

import (
	"github.com/RakutenReady/terraform-impact/impact"
)

func createOutputer(options impactOptions) impact.Outputer {
	return impact.NewStdOutOutputer()
}
