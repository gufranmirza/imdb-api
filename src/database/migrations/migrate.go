package migrations

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gufranmirza/imdb-api/config"
	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Migrate applies or create a new migration for database
func Migrate(action, message string) {
	config := config.Config
	opt := options.Client().ApplyURI(config.Database.Host)
	client, err := mongo.NewClient(opt)
	if err != nil {
		log.Fatal(err.Error())
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(config.Database.QueryTimeoutInSec)*time.Second,
	)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	db := client.Database(config.Database.Database)
	migrate.SetDatabase(db)
	migrate.SetMigrationsCollection("migrations")

	switch action {
	case "new":
		if len(message) == 0 {
			log.Fatal("Provide message for new migration")
		}
		fName := fmt.Sprintf("./database/migrations/%s_%s.go", time.Now().Format("20060102150405"), strings.ReplaceAll(message, " ", "_"))
		from, err := os.Open("./database/migrations/template.go")
		if err != nil {
			log.Fatal("Migration template not found")
		}
		defer from.Close()

		to, err := os.OpenFile(fName, os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer to.Close()

		_, err = io.Copy(to, from)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Printf("New migration created: %s\n", fName)
	case "up":
		err = migrate.Up(migrate.AllAvailable)
	case "down":
		err = migrate.Down(migrate.AllAvailable)
	default:
		log.Fatal("Invalid operation")
	}

	if err != nil {
		log.Fatal(err.Error())
	}
}
