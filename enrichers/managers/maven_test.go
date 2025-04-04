package managers

import (
	"testing"

	"github.com/fnxpt/cyclonedx-enrich/utils"

	"github.com/CycloneDX/cyclonedx-go"
)

func TestMavenEnricher_Skip(t *testing.T) {
	tests := []struct {
		name      string
		component *cyclonedx.Component
		want      bool
	}{
		{name: "Test with nil package", component: nil, want: true},
		{name: "Test with empty package", component: utils.ComponentEmpty, want: true},
		{name: "Test with cocoapods package", component: utils.ComponentCocoapods, want: true},
		{name: "Test with maven package", component: utils.ComponentMaven, want: false},
		{name: "Test with npm package", component: utils.ComponentNpm, want: true},
		{name: "Test with pypi package", component: utils.ComponentPypi, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &MavenEnricher{}
			if got := e.Skip(tt.component); got != tt.want {
				t.Errorf("MavenEnricher.Skip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMavenEnricher_Enrich(t *testing.T) {
	tests := []struct {
		name      string
		component *cyclonedx.Component
		wantErr   bool
	}{
		{name: "Test with maven package", component: utils.ComponentMaven, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &MavenEnricher{}
			if err := e.Enrich(tt.component); (err != nil) != tt.wantErr {
				t.Errorf("MavenEnricher.Enrich() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
