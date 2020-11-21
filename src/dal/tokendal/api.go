package tokendal

import (
	"github.com/gufranmirza/imdb-api/database/dbmodels"
)

// TokenDal interface
type TokenDal interface {
	Create(txID string, token *dbmodels.Token) (*dbmodels.Token, error)
	GetByUUID(uuid string) (*dbmodels.Token, error)
	Update(token *dbmodels.Token) error
	DeleteByAccessToken(token string) error
}
