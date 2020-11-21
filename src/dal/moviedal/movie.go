package moviedal

import (
	"context"
	"fmt"
	"time"

	"github.com/gufranmirza/imdb-api/config"
	"github.com/gufranmirza/imdb-api/models"

	"github.com/gufranmirza/imdb-api/database/connection"
	"github.com/gufranmirza/imdb-api/database/dbmodels"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type movie struct {
	db     connection.MongoStore
	config *models.AppConfig
}

// NewMovieDal returns the implementation
func NewMovieDal() MovieDal {
	return &movie{
		db:     connection.NewMongoStore(),
		config: config.Config,
	}
}

// Create creates a new movie into db.
func (r *movie) Create(txID string, movie *dbmodels.Movie) error {
	rc := r.db.Database().Collection(r.config.Database.MovieCollection)
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(r.config.Database.QueryTimeoutInSec)*time.Second,
	)
	defer cancel()

	_, err := rc.InsertOne(ctx, movie)
	if err != nil {
		return fmt.Errorf("Failed to create repo with error %v", err)
	}

	return nil
}

// Search looks for movie matching with query.
func (r *movie) Search(query string) ([]dbmodels.Movie, error) {
	rc := r.db.Database().Collection(r.config.Database.MovieCollection)
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(r.config.Database.QueryTimeoutInSec)*time.Second,
	)
	defer cancel()

	filter := bson.D{
		{"$or", []interface{}{
			bson.D{{"Director", primitive.Regex{Pattern: query, Options: "i"}}},
			bson.D{{"Name", primitive.Regex{Pattern: query, Options: "i"}}},
		}},
	}

	filterCursor, err := rc.Find(ctx, filter)
	if err != nil {
		return []dbmodels.Movie{}, fmt.Errorf("Failed to search movie with error %v", err)
	}

	var movie []dbmodels.Movie
	if err = filterCursor.All(ctx, &movie); err != nil {
		return []dbmodels.Movie{}, fmt.Errorf("Failed to search movie with error %v", err)
	}

	return movie, nil
}
