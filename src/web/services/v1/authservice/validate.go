package authservice

import (
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/gufranmirza/imdb-api/models/authmodel"
	"github.com/gufranmirza/imdb-api/web/renderers"
)

// @Summary Validate
// @Description It allows valdidate jwt token
// @Tags authentication
// @Param Authorization header string true "BEARER JWT"
// @Accept json
// @Produce json
// @Success 200
// @Failure 401 {object} errorinterface.ErrorResponse{}
// @Failure 404 {object} errorinterface.ErrorResponse{}
// @Router /authentication/validate [POST]
func (as *authservice) Validate(w http.ResponseWriter, r *http.Request) {
	_, _, err := jwtauth.FromContext(r.Context())
	if err != nil {
		renderers.ErrorUnauthorized(w, r, authmodel.ErrLoginToken)
		return
	}

	w.WriteHeader(http.StatusOK)
}
