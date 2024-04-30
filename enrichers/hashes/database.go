package hashes

import (
	"cyclonedx-enrich/models"
	"fmt"
	"log/slog"
	"os"

	"cyclonedx-enrich/utils"

	"github.com/CycloneDX/cyclonedx-go"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseEnricher struct {
	models.Enricher
	isInitialized bool
	database      *gorm.DB
}

func connect() (*gorm.DB, error) {
	filename := os.Getenv("DATABASE_FILE")
	return gorm.Open(sqlite.Open(filename), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
}

func (e *DatabaseEnricher) Category() models.EnricherCategory {
	return "hashes" //TODO: REFLECT package name
}

func (e *DatabaseEnricher) Skip(component *cyclonedx.Component) bool {
	return component.Hashes != nil
}

func (e *DatabaseEnricher) Enrich(component *cyclonedx.Component) error {
	if !e.isInitialized {
		db, err := connect()

		if err != nil {
			log.Warn("unable to connect to database",
				slog.String("error", err.Error()))
		}

		e.database = db
		e.isInitialized = true
	}

	if e.database == nil {
		return nil
	}

	var item *models.Component
	e.database.Where("purl = ?", utils.GetRealPurl(component.PackageURL)).Preload("Hashes").First(&item)

	if item != nil {

		for _, hash := range item.Hashes {
			if !hasKey(*component.Hashes, hash.Name) {
				*component.Hashes = append(*component.Hashes, cyclonedx.Hash{
					Algorithm: cyclonedx.HashAlgorithm(hash.Name),
					Value:     hash.Value,
				})
			}
		}
	}

	return fmt.Errorf("component doesn't met criteria")
}
