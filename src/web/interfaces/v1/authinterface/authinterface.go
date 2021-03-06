package authinterface

import "github.com/gufranmirza/imdb-api/models/authmodel"

// LoginReqInterface holds the login details
type LoginReqInterface struct {
	Email string `json:"email"`
}

// SignUpReqInterface sign up request details
type SignUpReqInterface struct {
	Email     string           `json:"email,omitempty"`
	FirstName string           `json:"first_name,omitempty"`
	Roles     []authmodel.Role `json:"roles,omitempty"`
}

// AuthenticateResInterface holds the token pair
type AuthenticateResInterface struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
