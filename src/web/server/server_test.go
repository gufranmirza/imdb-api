package server

import (
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/gufranmirza/imdb-api/config"
	"github.com/gufranmirza/imdb-api/web/router"
)

func TestNewServer(t *testing.T) {
	tests := []struct {
		name string
		want Server
	}{
		{
			name: "Happy Path",
			want: &server{
				logger: log.New(os.Stdout, "server :=> ", log.LstdFlags),
				config: config.Config,
				router: router.NewRouter(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
