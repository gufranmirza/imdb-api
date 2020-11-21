package migrations

import (
	"context"
	"log"
	"time"

	"github.com/gufranmirza/imdb-api/config"
	"github.com/gufranmirza/imdb-api/models"
	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
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
			Keys: bson.M{"Name": 1},
		},
		{
			Keys: bson.M{"Director": 1},
		},
	}

	migrate.Register(func(db *mongo.Database) error {
		_, err := db.Collection(config.Database.MovieCollection).Indexes().CreateMany(ctx, mod)
		return err
	}, func(db *mongo.Database) error { //Down
		_, err := db.Collection(config.Database.MovieCollection).Indexes().DropOne(ctx, "Name_1")
		if err != nil {
			return err
		}
		_, err = db.Collection(config.Database.MovieCollection).Indexes().DropOne(ctx, "Director_1")
		return err
	})
}
