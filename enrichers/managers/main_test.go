package managers

import (
	"testing"

	"github.com/CycloneDX/cyclonedx-go"
)

func Test_skip(t *testing.T) {
	type args struct {
		component *cyclonedx.Component
		prefix    string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := skip(tt.args.component, tt.args.prefix); got != tt.want {
				t.Errorf("skip() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_enrich(t *testing.T) {
// 	type args struct {
// 		component  *cyclonedx.Component
// 		licenses   []string
// 		hashes     map[string]string
// 		references map[string]string
// 		properties map[string]string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			enrich(tt.args.component, tt.args.licenses, tt.args.hashes, tt.args.references, tt.args.properties)
// 		})
// 	}
// }

// func Test_request(t *testing.T) {
// 	type args struct {
// 		url     string
// 		headers map[string]string
// 		fn      func(io.ReadCloser) (*T, error)
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    *T
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := request(tt.args.url, tt.args.headers, tt.args.fn)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("request() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("request() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_parseJSON(t *testing.T) {
// 	type args struct {
// 		input io.ReadCloser
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    *T
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := parseJSON(tt.args.input)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("parseJSON() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("parseJSON() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
