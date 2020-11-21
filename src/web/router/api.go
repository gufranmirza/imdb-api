package router

import (
	"github.com/go-chi/chi"
)

// Router interface
type Router interface {
	Router() chi.Router
}
