package dbmodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Movie model represents the movie collection in database
type Movie struct {
	ID                  primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedTimestampUTC time.Time          `json:"created_timestamp_utc,omitempty" bson:"CreatedTimestampUTC,omitempty"`
	UpdatedTimestampUTC time.Time          `json:"updated_timestamp_utc,omitempty" bson:"UpdatedTimestampUTC,omitempty"`
	Popularity99        float64            `json:"99popularity,omitempty" bson:"Popularity99,omitempty"`
	Director            string             `json:"director,omitempty" bson:"Director,omitempty"`
	Genre               []string           `json:"genre,omitempty" bson:"Genre,omitempty"`
	IMDBScore           float64            `json:"imdb_score,omitempty" bson:"IMDBScore,omitempty"`
	Name                string             `json:"name,omitempty" bson:"Name,omitempty"`
}
