package cli

import (
	"github.com/RakutenReady/terraform-impact/impact"
)

func createOutputer(options ImpactOptions) impact.Outputer {
	return impact.NewStdOutOutputer()
}
