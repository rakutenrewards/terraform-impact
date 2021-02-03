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
	impacter := createImpacter(opts)
	service := createService(opts)
	outputer := createOutputer(opts)

	return impacter, service, outputer
}
