package authservice

import (
	"encoding/json"
	"net/http"

	"github.com/gufranmirza/imdb-api/models"
	"github.com/gufranmirza/imdb-api/models/authmodel"
	"github.com/gufranmirza/imdb-api/web/interfaces/v1/authinterface"
	"github.com/gufranmirza/imdb-api/web/renderers"

	"github.com/go-chi/jwtauth"
)

// @Summary Logout
// @Description It allows to logout users from account with JWT
// @Tags authentication
// @Param Authorization header string true "BEARER JWT"
// @Accept json
// @Produce json
// @Success 200
// @Failure 401 {object} errorinterface.ErrorResponse{}
// @Failure 404 {object} errorinterface.ErrorResponse{}
// @Router /authentication/logout [POST]
func (as *authservice) Logout(w http.ResponseWriter, r *http.Request) {
	txID := r.Header.Get(models.HdrRequestID)

	data := &authinterface.LoginReqInterface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		renderers.ErrorBadRequest(w, r, authmodel.ErrIncorrectDetails)
		return
	}

	if len(data.Email) == 0 {
		renderers.ErrorBadRequest(w, r, authmodel.ErrIncorrectDetails)
		return
	}

	token, _, err := jwtauth.FromContext(r.Context())
	if err != nil {
		renderers.ErrorUnauthorized(w, r, authmodel.ErrLoginToken)
		return
	}

	err = as.tokenDal.DeleteByAccessToken(token.Raw)
	if err != nil {
		as.logger.Printf("%s %s Failed to delete access token with error %v", txID, authmodel.FailedToDeleteToken, err)
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
