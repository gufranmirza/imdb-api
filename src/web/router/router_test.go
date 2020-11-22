// Package api configures an http server for administration and application resources.

package router

import (
	"testing"

	"github.com/go-chi/chi"
	"github.com/gufranmirza/imdb-api/auth/jwt"
	"github.com/gufranmirza/imdb-api/config"
	"github.com/gufranmirza/imdb-api/models"
	_ "github.com/gufranmirza/imdb-api/web/docs"
	"github.com/gufranmirza/imdb-api/web/services/health"
	"github.com/gufranmirza/imdb-api/web/services/v1/authservice"
	"github.com/gufranmirza/imdb-api/web/services/v1/movieservice"
)

func TestNewRouter(t *testing.T) {
	config.Config = &models.AppConfig{
		JWT: &models.JWT{
			Secret: "1456465",
		},
	}
	tests := []struct {
		name string
		want Router
	}{
		{
			name: "Happy Path",
			want: &router{
				config:  config.Config,
				health:  health.NewHealth(),
				auth:    authservice.NewAuthService(),
				movie:   movieservice.NewMovieServiceService(),
				jwtauth: jwt.NewTokenAuth(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRouter(); got == nil && tt.want != nil {
				t.Errorf("NewRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_router_Router(t *testing.T) {
	config.Config = &models.AppConfig{
		JWT: &models.JWT{
			Secret: "1456465",
		},
	}

	type fields struct {
		config  *models.AppConfig
		health  health.Health
		auth    authservice.AuthService
		movie   movieservice.MovieService
		jwtauth jwt.TokenAuth
	}
	tests := []struct {
		name   string
		fields fields
		want   chi.Router
	}{
		{
			name: "Happy Path",
			fields: fields{
				config:  config.Config,
				health:  health.NewHealth(),
				auth:    authservice.NewAuthService(),
				movie:   movieservice.NewMovieServiceService(),
				jwtauth: jwt.NewTokenAuth(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := &router{
				config:  tt.fields.config,
				health:  tt.fields.health,
				auth:    tt.fields.auth,
				movie:   tt.fields.movie,
				jwtauth: tt.fields.jwtauth,
			}
			if got := router.Router(); got == nil && tt.want != nil {
				t.Errorf("router.Router() = %v, want %v", got, tt.want)
			}
		})
	}
}
