package dbmodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Movie model represents the movie collection in database
type Movie struct {
	ID                  primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedTimestampUTC time.Time          `json:"created_timestamp_utc,omitempty" bson:"created_timestamp_utc,omitempty"`
	UpdatedTimestampUTC time.Time          `json:"updated_timestamp_utc,omitempty" bson:"updated_timestamp_utc,omitempty"`
	Popularity99        float64            `json:"99popularity,omitempty" bson:"99popularity,omitempty"`
	Director            string             `json:"director,omitempty" bson:"director,omitempty"`
	Genre               []string           `json:"genre,omitempty" bson:"genre,omitempty"`
	IMDBScore           float64            `json:"imdb_score,omitempty" bson:"imdb_score,omitempty"`
	Name                string             `json:"name,omitempty" bson:"name,omitempty"`
}
