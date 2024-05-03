package properties

import (
	"testing"

	"cyclonedx-enrich/utils"

	"github.com/CycloneDX/cyclonedx-go"
)

func TestDatabaseEnricher_Skip(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)
	tests := []struct {
		name      string
		component *cyclonedx.Component
		want      bool
	}{
		//TODO: CONTINUE
		// {name: "Test with empty component", component: utils.ComponentEmpty, want: true},  //TODO: FAILING
		// {name: "Test with component with data", component: utils.ComponentWithData, want: true}, //TODO: FAILING
		{name: "Test with component without data", component: utils.ComponentWithoutData, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &DatabaseEnricher{}
			if got := e.Skip(tt.component); got != tt.want {
				t.Errorf("DatabaseEnricher.Skip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseEnricher_Enrich(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)
	tests := []struct {
		name      string
		component *cyclonedx.Component
		wantErr   bool
	}{
		//TODO: VALIDATE IF DATA WAS ADDED
		// {name: "Test with component without data", component: utils.ComponentWithoutData, wantErr: false},  //TODO: FAILING
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &DatabaseEnricher{}
			if err := e.Enrich(tt.component); (err != nil) != tt.wantErr {
				t.Errorf("DatabaseEnricher.Enrich() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
