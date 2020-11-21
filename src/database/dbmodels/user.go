package dbmodels

import (
	"time"

	"github.com/gufranmirza/imdb-api/models/authmodel"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User model represents the user collection in database
type User struct {
	ID                  primitive.ObjectID `json:"ID,omitempty" bson:"_id,omitempty"`
	CreatedTimestampUTC time.Time          `json:"CreatedTimestampUTC,omitempty" bson:"CreatedTimestampUTC,omitempty"`
	UpdatedTimestampUTC time.Time          `json:"UpdatedTimestampUTC,omitempty" bson:"UpdatedTimestampUTC,omitempty"`
	LastLogin           time.Time          `json:"LastLogin,omitempty" bson:"LastLogin,omitempty"`
	Email               string             `json:"Email,omitempty" bson:"Email,omitempty"`
	FirstName           string             `json:"FirstName,omitempty" bson:"FirstName,omitempty"`
	LastName            string             `json:"LastName,omitempty" bson:"LastName,omitempty"`
	Active              bool               `json:"Active,omitempty" bson:"Active,omitempty"`
	Roles               []authmodel.Role   `json:"Roles,omitempty" bson:"Roles,omitempty"`
}

// CanLogin returns true if user is allowed to login.
func (r *User) CanLogin() bool {
	return r.Active
}

// Claims returns the account's claims to be signed
func (r *User) Claims() authmodel.AppClaims {
	return authmodel.AppClaims{
		ID:    r.ID.String(),
		Sub:   r.Email,
		Roles: r.Roles,
	}
}
