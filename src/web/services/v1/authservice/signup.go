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

// @Summary Sign up with email
// @Description It allows to sign up with email address and personal details
// @Tags authentication
// @Accept json
// @Produce json
// @Param * body authinterface.SignUpReqInterface{} true "signup with email"
// @Success 200 {object} authinterface.AuthenticateResInterface{}
// @Failure 400 {object} errorinterface.ErrorResponse{}
// @Failure 404 {object} errorinterface.ErrorResponse{}
// @Failure 500 {object} errorinterface.ErrorResponse{}
// @Router /authentication/signup [POST]
func (as *authservice) SignUp(w http.ResponseWriter, r *http.Request) {
	txID := r.Header.Get(models.HdrRequestID)

	data := &authinterface.SignUpReqInterface{}
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
	if account != nil {
		renderers.ErrorBadRequest(w, r, authmodel.ErrAlreadyRegistered)
		return
	}

	user := &dbmodels.User{}
	user.Roles = data.Roles
	user.Active = true
	user.CreatedTimestampUTC = time.Now().UTC()
	user.UpdatedTimestampUTC = time.Now().UTC()
	user.Email = data.Email
	user.FirstName = data.FirstName

	account, err = as.userDal.Create(txID, user)
	if err != nil {
		as.logger.Printf("%s %s Failed to create user record with error %v", txID, authmodel.FailedToSignUp, err)
		renderers.ErrorInternalServerError(w, r, authmodel.ErrServerError)
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
