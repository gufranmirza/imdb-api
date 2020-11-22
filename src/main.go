package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/gufranmirza/imdb-api/dal/moviedal"
	"github.com/gufranmirza/imdb-api/database/dbmodels"

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
		migrate(os.Args)
		return
	}

	// load data from json file
	if len(os.Args) > 1 && os.Args[1] == "dataloader" {
		dataloader(os.Args)
		return
	}

	// Start http server
	svr := server.NewServer()
	svr.Start()
}

func migrate(args []string) {
	if len(os.Args) < 3 {
		log.Printf("No enough arguments to run migrations args: (%s) \n", os.Args)
		log.Println("USAGE ./imdb-api migrate up/down/new `message here`")
		os.Exit(1)
	}

	action := os.Args[2]
	message := ""
	if len(os.Args) > 3 {
		message = os.Args[2]
	}
	migrations.Migrate(action, message)
}

func dataloader(args []string) {
	if len(os.Args) < 3 {
		log.Printf("No enough arguments to dataloader args: (%s) \n", os.Args)
		log.Println("USAGE ./imdb-api dataloader ./data.json")
		os.Exit(1)
	}

	path := args[2]

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to read file %s with error: %v \n", path, err)
	}
	defer file.Close()

	buff, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file %s with error: %v \n", path, err)
	}

	var movies []dbmodels.Movie
	moviedal := moviedal.NewMovieDal()
	err = json.Unmarshal(buff, &movies)
	if err != nil {
		log.Fatalf("Failed to unmarshal with error: %v \n", err)
	}

	for _, movie := range movies {
		err = moviedal.Create("", &movie)
		if err != nil {
			log.Printf("Failed to insert movie with error: %v \n", err)
			continue
		}
	}
}
