package cli

import (
	"github.com/RakutenReady/terraform-impact/impact"
)

func createService(opts ImpactOptions) impact.ImpactService {
	stateLister := createStateLister(opts)

	if opts.ListStates {
		return impact.NewListOnlyImpactService(stateLister)
	}

	return impact.NewImpactService(stateLister)
}
