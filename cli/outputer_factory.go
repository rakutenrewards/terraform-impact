package cli

import (
	"fmt"
	"strings"

	"github.com/RakutenReady/terraform-impact/impact"
)

func createOutputer(opts ImpactOptions) impact.Outputer {
	if len(opts.Output) == 0 {
		return impact.NewStdOutOutputer()
	}

	if strings.HasSuffix(opts.Output, ".json") {
		return impact.NewJsonOutputer(opts.Output)
	}

	panic(fmt.Errorf("Unknown output format [%v]", opts.Output))
}
