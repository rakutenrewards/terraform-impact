package impact

type ListOnlyImpactService struct {
	lister StateLister
}

func NewListOnlyImpactService(lister StateLister) ImpactService {
	return ListOnlyImpactService{lister}
}

// Impact simply lists the state directories
func (service ListOnlyImpactService) Impact(paths []string) ([]string, error) {
	return service.lister.List(), nil
}
