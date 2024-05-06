package managers

import (
	"cyclonedx-enrich/utils"
	"testing"

	"github.com/CycloneDX/cyclonedx-go"
)

func TestGithubEnricher_Skip(t *testing.T) {

	tests := []struct {
		name      string
		component *cyclonedx.Component
		want      bool
	}{
		{name: "Test without component with github data", component: utils.ComponentCocoapods, want: true},
		{name: "Test with component with github data", component: utils.ComponentWithData, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &GithubEnricher{}
			if got := e.Skip(tt.component); got != tt.want {
				t.Errorf("GithubEnricher.Skip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGithubEnricher_Enrich(t *testing.T) {
	tests := []struct {
		name      string
		component *cyclonedx.Component
		wantErr   bool
	}{
		{name: "Test with component with github data", component: utils.ComponentWithData, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &GithubEnricher{}
			if err := e.Enrich(tt.component); (err != nil) != tt.wantErr {
				t.Errorf("GithubEnricher.Enrich() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
