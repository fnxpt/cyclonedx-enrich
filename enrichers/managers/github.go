package managers

import (
	"cyclonedx-enrich/models"
	"cyclonedx-enrich/utils"
	"fmt"
	"strings"

	"github.com/CycloneDX/cyclonedx-go"
)

var githubEndpoint = "https://api.github.com/repos"
var githubToken = utils.Getenv("GITHUB_TOKEN", "")

type GithubEnricher struct {
	models.Enricher
}

func (e *GithubEnricher) Skip(component *cyclonedx.Component) bool {
	return skip(component, "pkg:") || len(githubToken) == 0 || githubURL(component) == nil
}

func (e *GithubEnricher) Enrich(component *cyclonedx.Component) error {
	githubRef := githubURL(component)

	if githubRef == nil {
		return fmt.Errorf("unable to extract github reference")
	}

	githubData := strings.Split(*githubRef, "/")
	if len(githubData) < 3 {
		return fmt.Errorf("unable to extract valid github data")
	}

	url := fmt.Sprintf("%s/%s/%s", githubEndpoint, githubData[1], githubData[2])

	item, err := request[GithubPackage](url, nil, parseJSON)

	if err != nil {
		return err
	}

	if item != nil {
		licenses := make([]string, 0)
		hashes := make(map[string]string)
		references := make(map[string]string)
		properties := make(map[string]string)

		if item.License != nil {
			licenses = append(licenses, item.License.SpdxID)
		}

		enrich(component, licenses, hashes, references, properties)
	}

	//TODO: USE MORE DATA

	return nil
}

func githubURL(component *cyclonedx.Component) *string {
	if component != nil {
		if strings.HasPrefix(component.Name, "github.com/") {
			return &component.Name
		}
		if item := hasKey(*component.ExternalReferences, "vcs"); item != nil {
			return &item.URL
		}
	}
	return nil
}

func hasKey(items []cyclonedx.ExternalReference, key string) *cyclonedx.ExternalReference {
	for _, item := range items {
		if string(item.Type) == key {
			return &item
		}
	}

	return nil
}
