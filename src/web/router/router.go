// Package api configures an http server for administration and application resources.

package router

import (
	"github.com/go-chi/chi"

	"github.com/gufranmirza/imdb-api/auth/jwt"
	"github.com/gufranmirza/imdb-api/config"
	"github.com/gufranmirza/imdb-api/models"
	"github.com/gufranmirza/imdb-api/models/authmodel"
	"github.com/gufranmirza/imdb-api/web/middlewares"
	"github.com/gufranmirza/imdb-api/web/services/health"
	"github.com/gufranmirza/imdb-api/web/services/v1/authservice"
	"github.com/gufranmirza/imdb-api/web/services/v1/movieservice"
)

type router struct {
	config *models.AppConfig
	health health.Health
	auth   authservice.AuthService
	movie  movieservice.MovieService

	jwtauth jwt.TokenAuth
}

// NewRouter returns the router implementation
func NewRouter() Router {
	return &router{
		config:  config.Config,
		health:  health.NewHealth(),
		auth:    authservice.NewAuthService(),
		movie:   movieservice.NewMovieServiceService(),
		jwtauth: jwt.NewTokenAuth(),
	}
}

// Router configures application resources and routes.
func (router *router) Router() chi.Router {
	r := chi.NewRouter()

	rprivate := chi.NewRouter()
	rprivate.Use(router.jwtauth.Verifier())
	rprivate.Use(middlewares.Authenticator)
	r.Mount("/", rprivate)

	// URL router prefix
	v1Prefix := router.config.URLPrefix + router.config.APIVersionV1

	// =================  health routes ======================
	r.Get(v1Prefix+"/health/", router.health.GetHealth)

	// =================  authentication routes ======================
	r.Post(v1Prefix+"/auth/signup", router.auth.SignUp)
	r.Post(v1Prefix+"/auth/login", router.auth.Login)
	rprivate.Post(v1Prefix+"/auth/logout", router.auth.Logout)
	rprivate.Post(v1Prefix+"/auth/validate", router.auth.Validate)

	// =================  movie routes ======================
	r.Group(func(r chi.Router) {
		r.Use(router.jwtauth.Verifier())
		r.Use(middlewares.Authenticator)
		r.Use(middlewares.RequiresRole(authmodel.Role("ADMIN"))) // try changing role to something else
		r.Post(v1Prefix+"/movies", router.movie.Create)
	})

	r.Get(v1Prefix+"/movies/search", router.movie.Search)

	return r
}
