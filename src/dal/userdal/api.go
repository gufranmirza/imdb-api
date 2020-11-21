package userdal

import (
	"github.com/gufranmirza/imdb-api/database/dbmodels"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserDal interface
type UserDal interface {
	Create(txID string, account *dbmodels.User) (*dbmodels.User, error)
	GetByEmail(email string) (*dbmodels.User, error)
	GetByID(id primitive.ObjectID) (*dbmodels.User, error)
	Update(user *dbmodels.User) error
}
