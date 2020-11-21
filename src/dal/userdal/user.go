package userdal

import (
	"context"
	"fmt"
	"time"

	"github.com/gufranmirza/imdb-api/config"
	"github.com/gufranmirza/imdb-api/models"

	"github.com/gufranmirza/imdb-api/database/connection"
	"github.com/gufranmirza/imdb-api/database/dbmodels"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type userdal struct {
	db     connection.MongoStore
	config *models.AppConfig
}

// NewUserDal returns the implementation
func NewUserDal() UserDal {
	return &userdal{
		db:     connection.NewMongoStore(),
		config: config.Config,
	}
}

// Create creates a new account.
func (r *userdal) Create(txID string, account *dbmodels.User) (*dbmodels.User, error) {
	rc := r.db.Database().Collection(r.config.Database.UserCollection)
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(r.config.Database.QueryTimeoutInSec)*time.Second,
	)
	defer cancel()

	insertResult, err := rc.InsertOne(ctx, account)
	if err != nil {
		return nil, fmt.Errorf("Failed to create user with error %v", err)
	}

	account.ID = insertResult.InsertedID.(primitive.ObjectID)
	return account, nil
}

func (r *userdal) Update(user *dbmodels.User) error {
	rc := r.db.Database().Collection(r.config.Database.UserCollection)
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(r.config.Database.QueryTimeoutInSec)*time.Second,
	)
	defer cancel()

	if _, err := rc.ReplaceOne(ctx, bson.M{"_id": user.ID}, user); err != nil {
		return err
	}

	return nil
}

func (r *userdal) GetByEmail(email string) (*dbmodels.User, error) {
	rc := r.db.Database().Collection(r.config.Database.UserCollection)
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(r.config.Database.QueryTimeoutInSec)*time.Second,
	)
	defer cancel()

	var rec dbmodels.User
	if err := rc.FindOne(ctx, bson.M{"Email": email}).Decode(&rec); err != nil {
		return nil, err
	}

	return &rec, nil
}

func (r *userdal) GetByID(id primitive.ObjectID) (*dbmodels.User, error) {
	rc := r.db.Database().Collection(r.config.Database.UserCollection)
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(r.config.Database.QueryTimeoutInSec)*time.Second,
	)
	defer cancel()

	var rec dbmodels.User
	if err := rc.FindOne(ctx, bson.M{"_id": id}).Decode(&rec); err != nil {
		return nil, err
	}

	return &rec, nil
}
