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
			Keys:    bson.M{"TokenUUID": 1},
			Options: options.Index().SetUnique(true),
		},
	}

	migrate.Register(func(db *mongo.Database) error {
		_, err := db.Collection(config.Database.JWTCollection).Indexes().CreateMany(ctx, mod)
		return err
	}, func(db *mongo.Database) error { //Down
		_, err := db.Collection(config.Database.JWTCollection).Indexes().DropOne(ctx, "TokenUUID_1")
		return err
	})
}
