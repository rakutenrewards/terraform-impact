package cli

import (
	"github.com/RakutenReady/terraform-impact/impact"
)

func createImpacter(opts impactOptions) impact.Impacter {
	inner := impact.NewCommandLineImpacter(opts.Files)
	return impact.NewImpacter(inner)
}
