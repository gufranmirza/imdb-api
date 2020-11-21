package models

import "time"

const (
	// DefaultConfigPath is path used when config path is not provided
	DefaultConfigPath = "./app-config.json"
)

// AppConfig  is the overall config that file that our application will use
type AppConfig struct {
	Hostname        string `json:"hostname"`
	Port            int    `json:"port"`
	ServiceName     string `json:"service_name"`
	ServiceProvider string `json:"service_provider"`
	ServiceVersion  string `json:"service_version"`
	URLPrefix       string `json:"url_prefix"`
	APIVersionV1    string `json:"api_version_v1"`
	JWT             *JWT   `json:"jwt"`
	Database        *DB    `json:"database"`
}

// JWT holds the config related to jwt authentication
type JWT struct {
	Secret           string `json:"secret"`
	JwtExpiry        int64  `json:"jwt_expiry_in_sec"`
	JwtRefreshExpiry int64  `json:"jwt_refresh_expiry_in_sec"`
}

// DB holds the config related to database
type DB struct {
	Host              string        `json:"host"`
	Database          string        `json:"database"`
	QueryTimeoutInSec time.Duration `json:"query_timeout_in_sec"`
	UserCollection    string        `json:"user_collection"`
	JWTCollection     string        `json:"jwt_collection"`
	MovieCollection   string        `json:"movie_collection"`
}
