package authservice

import (
	"log"
	"os"

	"github.com/gufranmirza/imdb-api/config"
	"github.com/gufranmirza/imdb-api/models"

	"github.com/gufranmirza/imdb-api/auth/jwt"
	"github.com/gufranmirza/imdb-api/dal/tokendal"
	"github.com/gufranmirza/imdb-api/dal/userdal"
)

type authservice struct {
	tokenDal  tokendal.TokenDal
	userDal   userdal.UserDal
	tokenAuth jwt.TokenAuth
	logger    *log.Logger
	config    *models.AppConfig
}

// NewAuthService returns service implementation
func NewAuthService() AuthService {
	return &authservice{
		tokenDal:  tokendal.NewTokenDal(),
		userDal:   userdal.NewUserDal(),
		tokenAuth: jwt.NewTokenAuth(),
		config:    config.Config,
		logger:    log.New(os.Stdout, "authservice :=> ", log.LstdFlags),
	}
}
