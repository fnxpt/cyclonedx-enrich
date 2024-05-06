package licenses

import (
	"cyclonedx-enrich/models"
	"cyclonedx-enrich/utils"

	"github.com/CycloneDX/cyclonedx-go"
)

type RegexpEnricher struct {
	models.Enricher
}

func (e *RegexpEnricher) Skip(component *cyclonedx.Component) bool {
	return len(utils.LoadRules()) == 0 || skip(component)
}

func (e *RegexpEnricher) Enrich(component *cyclonedx.Component) error {
	return utils.EnrichRules(component, func(item *models.RuleEntry) error {
		return enrich(component, item.Licenses)
	})
}
