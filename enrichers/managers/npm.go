package managers

import (
	"fmt"

	"github.com/fnxpt/cyclonedx-enrich/models"
	"github.com/fnxpt/cyclonedx-enrich/utils"

	"github.com/CycloneDX/cyclonedx-go"
)

var npmEndpoint = "https://registry.npmjs.org"

type NPMEnricher struct {
	models.Enricher
}

func (e *NPMEnricher) Skip(component *cyclonedx.Component) bool {
	return skip(component, "pkg:npm/")
}

func (e *NPMEnricher) Enrich(component *cyclonedx.Component) error {
	var url string

	if len(component.Group) > 0 {
		url = fmt.Sprintf("%s/%s/%s/%s", npmEndpoint, utils.Decoded(component.Group), component.Name, component.Version)
	} else {
		url = fmt.Sprintf("%s/%s/%s", npmEndpoint, component.Name, component.Version)
	}

	item, err := request[NpmPackage](url, nil, parseJSON)

	if err != nil {
		return err
	}

	if item != nil {
		licenses := make([]string, 0)
		hashes := make(map[string]string)
		references := make(map[string]string)
		properties := make(map[string]string)

		if item.License != nil {
			licenses = append(licenses, *item.License)
		}

		enrich(component, licenses, hashes, references, properties)
	}

	//TODO: USE MORE DATA

	return nil
}
