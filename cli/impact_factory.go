package cli

import (
	"github.com/RakutenReady/terraform-impact/impact"
)

type impactFactory interface {
	Create(opts ImpactOptions) (impact.Impacter, impact.ImpactService, impact.Outputer)
}

type impactFactoryImpl struct {
}

func newImpactFactory() impactFactory {
	return impactFactoryImpl{}
}

func (factory impactFactoryImpl) Create(opts ImpactOptions) (impact.Impacter, impact.ImpactService, impact.Outputer) {
	stateLister := createStateLister(opts)
	impacter := createImpacter(opts)
	outputer := createOutputer(opts)

	service := impact.NewImpactService(stateLister)

	return impacter, service, outputer
}
