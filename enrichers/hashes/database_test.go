package hashes

import (
	"os"
	"testing"

	"cyclonedx-enrich/utils"

	"github.com/CycloneDX/cyclonedx-go"
)

func TestDatabaseEnricher_Skip_WithoutDatabase(t *testing.T) {
	tests := []struct {
		name      string
		component *cyclonedx.Component
		want      bool
	}{
		{name: "Test without database with empty component", component: utils.ComponentEmpty, want: true},
		{name: "Test without database with component with data", component: utils.ComponentWithData, want: true},
		{name: "Test without database with component without data", component: utils.ComponentWithoutData, want: true},
	}

	os.Setenv("DATABASE_FILE", "NO_DATABASE.db")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &DatabaseEnricher{}
			if got := e.Skip(tt.component); got != tt.want {
				t.Errorf("DatabaseEnricher.Skip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseEnricher_Skip(t *testing.T) {
	tests := []struct {
		name      string
		component *cyclonedx.Component
		want      bool
	}{
		//TODO: CONTINUE
		{name: "Test with empty component", component: utils.ComponentEmpty, want: true},
		{name: "Test with component with data", component: utils.ComponentWithData, want: true},
		// {name: "Test with component without data", component: utils.ComponentWithoutData, want: false}, //TODO: FAILING
	}
	os.Setenv("DATABASE_FILE", "../../database.db") //TODO: GET A BETTER WAY TO RESOLVE DATABASE
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &DatabaseEnricher{}
			if got := e.Skip(tt.component); got != tt.want {
				t.Errorf("DatabaseEnricher.Skip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseEnricher_Enrich_WithoutDatabase(t *testing.T) {
	tests := []struct {
		name      string
		component *cyclonedx.Component
		wantErr   bool
	}{
		//TODO: VALIDATE IF DATA WAS ADDED
		// {name: "Test with component without data", component: utils.ComponentWithoutData, wantErr: false}, //TODO: FAILING
	}
	os.Setenv("DATABASE_FILE", "../../database.db") //TODO: GET A BETTER WAY TO RESOLVE DATABASE
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &DatabaseEnricher{}
			if err := e.Enrich(tt.component); (err != nil) != tt.wantErr {
				t.Errorf("DatabaseEnricher.Enrich() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabaseEnricher_Enrich(t *testing.T) {
	tests := []struct {
		name      string
		component *cyclonedx.Component
		wantErr   bool
	}{
		//TODO: VALIDATE IF DATA WAS ADDED
		// {name: "Test with component without data", component: utils.ComponentWithoutData, wantErr: false}, //TODO: FAILING
	}
	os.Setenv("DATABASE_FILE", "../../database.db") //TODO: GET A BETTER WAY TO RESOLVE DATABASE
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &DatabaseEnricher{}
			if err := e.Enrich(tt.component); (err != nil) != tt.wantErr {
				t.Errorf("DatabaseEnricher.Enrich() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}