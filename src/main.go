package main

import (
	"log"
	"os"

	"github.com/gufranmirza/imdb-api/config"
	"github.com/gufranmirza/imdb-api/database/migrations"
	"github.com/gufranmirza/imdb-api/models"
	"github.com/gufranmirza/imdb-api/web/server"
)

// @title API Documentation
// @version 2.0
// @description API Documentation

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8001
// @BasePath /imdb-api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	log := log.New(os.Stdout, "main :=> ", log.LstdFlags)

	// Initialize Application configuration
	_, err := config.LoadConfig(models.DefaultConfigPath)
	if err != nil {
		log.Fatalf("Reading configuration from JSON (%s) failed: %s\n", models.DefaultConfigPath, err)
	}

	// Run migration if present in args
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		if len(os.Args) < 3 {
			log.Fatalf("No enough arguments to run migrations args: (%s) \n", os.Args)
		}
		action := os.Args[2]
		message := ""
		if len(os.Args) > 3 {
			message = os.Args[2]
		}
		migrations.Migrate(action, message)
		return
	}

	// Start http server
	svr := server.NewServer()
	svr.Start()
}
