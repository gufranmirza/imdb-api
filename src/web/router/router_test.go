// Package api configures an http server for administration and application resources.

package router

import (
	"reflect"
	"testing"

	"github.com/gufranmirza/imdb-api/config"
	"github.com/gufranmirza/imdb-api/web/services/health"
)

func TestNewRouter(t *testing.T) {
	tests := []struct {
		name string
		want Router
	}{
		{
			name: "Happy Path",
			want: &router{
				config: config.Config,
				health: health.NewHealth(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
