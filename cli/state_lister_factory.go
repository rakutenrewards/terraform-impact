package cli

import (
	"github.com/RakutenReady/terraform-impact/impact"
)

func createStateLister(opts impactOptions) impact.StateLister {
	return impact.NewDiscoveryStateLister(opts.RootDir, opts.Pattern)
}
