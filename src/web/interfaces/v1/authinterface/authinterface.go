package authinterface

import "github.com/gufranmirza/imdb-api/models/authmodel"

// LoginReqInterface holds the login details
type LoginReqInterface struct {
	Email string `json:"email"`
}

// SignUpReqInterface sign up request details
type SignUpReqInterface struct {
	Email     string           `json:"email,omitempty"`
	FirstName string           `json:"firstName,omitempty"`
	Roles     []authmodel.Role `json:"roles,omitempty"`
}

// AuthenticateReqInterface holds the login details
type AuthenticateReqInterface struct {
	Token string `json:"token"`
}

// AuthenticateResInterface holds the token pair
type AuthenticateResInterface struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
