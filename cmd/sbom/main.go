package sbom

import (
	"cyclonedx-enrich/models"
	"log/slog"
)

var log = slog.Default()

type SbomCMD struct {
	models.Commandable
}

func (c SbomCMD) Commands() []models.Command {

	return []models.Command{
		{Flag: "sbom-enrich", Description: "Enrichs sbom", NeedsValue: true, Handler: func(value string) error {
			return enrichFiles(value)
		}},
		{Flag: "sbom-validate", Description: "Validates sbom", NeedsValue: true, Handler: func(value string) error {
			return validateFiles(value)
		}},
	}
}
