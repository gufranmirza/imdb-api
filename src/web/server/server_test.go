package server

import (
	"log"
	"os"
	"testing"

	"github.com/gufranmirza/imdb-api/config"
	"github.com/gufranmirza/imdb-api/models"
	"github.com/gufranmirza/imdb-api/web/router"
)

func TestNewServer(t *testing.T) {
	config.Config = &models.AppConfig{
		JWT: &models.JWT{
			Secret: "1456465",
		},
	}

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
			if got := NewServer(); got == nil && tt.want != nil {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
