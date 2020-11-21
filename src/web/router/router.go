// Package api configures an http server for administration and application resources.

package router

import (
	"net/http"

	"github.com/gufranmirza/imdb-api/config"
	"github.com/gufranmirza/imdb-api/models"
	"github.com/gufranmirza/imdb-api/web/services/health"
)

type router struct {
	config *models.AppConfig
	health health.Health
}

// NewRouter returns the router implementation
func NewRouter() Router {
	return &router{
		config: config.Config,
		health: health.NewHealth(),
	}
}

// Router configures application resources and routes.
func (router *router) Router() *http.ServeMux {
	r := http.NewServeMux()

	// URL router prefix
	urlPrefix := router.config.URLPrefix

	// =================  health routes ======================
	r.Handle(urlPrefix+"/health/", router.health.GetHealth())

	return r
}
