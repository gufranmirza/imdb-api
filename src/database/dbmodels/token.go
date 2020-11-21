package dbmodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Token model represents the Token collection in database
type Token struct {
	ID                  primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedTimestampUTC time.Time          `json:"createdTimestampUTC,omitempty" bson:"CreatedTimestampUTC,omitempty"`
	UpdatedTimestampUTC time.Time          `json:"updatedTimestampUTC,omitempty" bson:"UpdatedTimestampUTC,omitempty"`

	ExpiryTimestampUTC time.Time          `json:"-" bson:"ExpiryTimestampUTC,omitempty"`
	AccountID          primitive.ObjectID `json:"-" bson:"AccountID,omitempty"`
	AccessToken        string             `json:"-" bson:"AccessToken,omitempty"`
	ResfreshToken      string             `json:"-" bson:"ResfreshToken,omitempty"`
	Claimed            bool               `json:"-" bson:"Claimed"`

	Mobile    bool   `json:"mobile,omitempty" bson:"Mobile,omitempty"`
	TokenUUID string `json:"tokenUUID,omitempty" bson:"TokenUUID,omitempty"`
	UserAgent string `json:"userAgent,omitempty" bson:"UserAgent,omitempty"`
}
