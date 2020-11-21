package tokendal

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gufranmirza/imdb-api/config"
	"github.com/gufranmirza/imdb-api/models"

	"github.com/gufranmirza/imdb-api/database/connection"
	"github.com/gufranmirza/imdb-api/database/dbmodels"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type token struct {
	db     connection.MongoStore
	config *models.AppConfig
}

// NewTokenDal returns the implementation
func NewTokenDal() TokenDal {
	return &token{
		db:     connection.NewMongoStore(),
		config: config.Config,
	}
}

// Create creates a new account.
func (t *token) Create(txID string, token *dbmodels.Token) (*dbmodels.Token, error) {
	tc := t.db.Database().Collection(t.config.Database.JWTCollection)
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(t.config.Database.QueryTimeoutInSec)*time.Second,
	)
	defer cancel()

	insertResult, err := tc.InsertOne(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("Failed to create user with error %v", err)
	}

	token.ID = insertResult.InsertedID.(primitive.ObjectID)
	return token, nil
}

func (t *token) Update(token *dbmodels.Token) error {
	tc := t.db.Database().Collection(t.config.Database.JWTCollection)
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(t.config.Database.QueryTimeoutInSec)*time.Second,
	)
	defer cancel()

	if _, err := tc.ReplaceOne(ctx, bson.M{"_id": token.ID}, token); err != nil {
		return err
	}

	return nil
}

func (t *token) DeleteByAccessToken(token string) error {
	tc := t.db.Database().Collection(t.config.Database.JWTCollection)
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(t.config.Database.QueryTimeoutInSec)*time.Second,
	)
	defer cancel()

	if _, err := tc.DeleteOne(ctx, bson.M{"AccessToken": strings.TrimSpace(token)}); err != nil {
		return err
	}

	return nil
}

func (t *token) GetByUUID(uuid string) (*dbmodels.Token, error) {
	tc := t.db.Database().Collection(t.config.Database.JWTCollection)
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(t.config.Database.QueryTimeoutInSec)*time.Second,
	)
	defer cancel()

	var rec dbmodels.Token
	if err := tc.FindOne(ctx, bson.M{"TokenUUID": uuid}).Decode(&rec); err != nil {
		return nil, err
	}

	return &rec, nil
}
