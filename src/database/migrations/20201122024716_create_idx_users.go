package migrations

import (
	"context"
	"log"
	"time"

	"github.com/gufranmirza/imdb-api/config"
	"github.com/gufranmirza/imdb-api/models"
	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	// Initialize Application configuration
	_, err := config.LoadConfig(models.DefaultConfigPath)
	if err != nil {
		log.Fatalf("Reading configuration from JSON (%s) failed: %s\n", models.DefaultConfigPath, err)
	}

	config := config.Config
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	mod := []mongo.IndexModel{
		{
			Keys:    bson.M{"Email": 1},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.M{"FirstName": 1},
		},
		{
			Keys: bson.M{"LastName": 1},
		},
	}

	migrate.Register(func(db *mongo.Database) error {
		_, err := db.Collection(config.Database.UserCollection).Indexes().CreateMany(ctx, mod)
		return err
	}, func(db *mongo.Database) error { //Down
		_, err := db.Collection(config.Database.UserCollection).Indexes().DropOne(ctx, "Email_1")
		if err != nil {
			return err
		}
		_, err = db.Collection(config.Database.UserCollection).Indexes().DropOne(ctx, "FirstName_1")
		if err != nil {
			return err
		}
		_, err = db.Collection(config.Database.UserCollection).Indexes().DropOne(ctx, "LastName_1")
		return err
	})
}
