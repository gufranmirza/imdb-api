package authservice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gufranmirza/imdb-api/database/dbmodels"
	"github.com/gufranmirza/imdb-api/models"
	"github.com/gufranmirza/imdb-api/models/authmodel"
	"github.com/gufranmirza/imdb-api/web/interfaces/v1/authinterface"
	"github.com/gufranmirza/imdb-api/web/renderers"
	"github.com/mssola/user_agent"
)

// @Summary Login to account with email
// @Description It allows to login to account with email address registered
// @Tags authentication
// @Accept json
// @Produce json
// @Param * body authinterface.LoginReqInterface{} true "login with email"
// @Success 200 {object} authinterface.AuthenticateResInterface{}
// @Failure 400 {object} errorinterface.ErrorResponse{}
// @Failure 404 {object} errorinterface.ErrorResponse{}
// @Failure 500 {object} errorinterface.ErrorResponse{}
// @Router /login [POST]
func (as *authservice) Login(w http.ResponseWriter, r *http.Request) {
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

	account, _ := as.userDal.GetByEmail(data.Email)
	if account == nil {
		renderers.ErrorBadRequest(w, r, authmodel.ErrUnknownLogin)
		return
	}

	if !account.CanLogin() {
		renderers.ErrorUnauthorized(w, r, authmodel.ErrLoginDisabled)
		return
	}

	access, refresh, err := as.tokenAuth.GenTokenPair(account.Claims(), authmodel.RefreshClaims{})
	if err != nil {
		as.logger.Printf("%s %s Failed to generate token with error %v", txID, authmodel.FailedToAuthenticateToken, err)
		renderers.ErrorInternalServerError(w, r, authmodel.ErrServerError)
		return
	}

	ua := user_agent.New(r.UserAgent())
	browser, _ := ua.Browser()
	token := &dbmodels.Token{}
	token.TokenUUID = uuid.New().String()
	token.AccessToken = access
	token.ResfreshToken = refresh
	token.Claimed = true
	token.ExpiryTimestampUTC = time.Now().UTC().Add(time.Duration(as.config.JWT.JwtExpiry) * time.Second)
	token.CreatedTimestampUTC = time.Now().UTC()
	token.UpdatedTimestampUTC = time.Now().UTC()
	token.UserAgent = fmt.Sprintf("%s on %s", browser, ua.OS())
	token.Mobile = ua.Mobile()

	_, err = as.tokenDal.Create(txID, token)
	if err != nil {
		as.logger.Printf("%s %s Failed to create access token with error %v", txID, authmodel.FailedToCreateAccessToken, err)
		renderers.ErrorInternalServerError(w, r, authmodel.ErrServerError)
		return
	}

	renderers.RenderJSON(w, r, txID, &authinterface.AuthenticateResInterface{
		AccessToken:  access,
		RefreshToken: refresh,
	})
	return
}
