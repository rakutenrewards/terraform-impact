package cli

import (
	"github.com/RakutenReady/terraform-impact/impact"
)

func createStateLister(opts ImpactOptions) impact.StateLister {
	return impact.NewDiscoveryStateLister(opts.GetRootDir(), opts.GetPattern())
}
