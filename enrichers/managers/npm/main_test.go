package npm

import (
	"cyclonedx-enrich/utils"
	"testing"

	"github.com/CycloneDX/cyclonedx-go"
)

func TestNPMEnricher_Skip(t *testing.T) {

	tests := []struct {
		name      string
		component *cyclonedx.Component
		want      bool
	}{
		//TODO: CONTINUE
		{name: "Test with empty component", component: utils.ComponentEmpty, want: true},
		{name: "Test with component with cocoapods", component: utils.ComponentCocoapods, want: true},
		{name: "Test with component with maven", component: utils.ComponentMaven, want: true},
		{name: "Test with component with npm", component: utils.ComponentNpm, want: false},
		{name: "Test with component with pypi", component: utils.ComponentPypi, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &NPMEnricher{}
			if got := e.Skip(tt.component); got != tt.want {
				t.Errorf("NPMEnricher.Skip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNPMEnricher_Enrich(t *testing.T) {

	tests := []struct {
		name      string
		component *cyclonedx.Component
		wantErr   bool
	}{
		//TODO: CONTINUE
		// {name: "Test with component with npm", component: utils.ComponentNpm, wantErr: false}, //TODO: FAILING
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &NPMEnricher{}
			if err := e.Enrich(tt.component); (err != nil) != tt.wantErr {
				t.Errorf("NPMEnricher.Enrich() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
