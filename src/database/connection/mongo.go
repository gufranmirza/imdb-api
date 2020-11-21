package connection

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/gufranmirza/imdb-api/config"
	"github.com/gufranmirza/imdb-api/models"

	"github.com/gufranmirza/imdb-api/web/interfaces/v1/healthinterface"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	once   sync.Once
	db     *mongo.Database
	client *mongo.Client
)

// MongoStore ...
type mongoStore struct {
	txID   string
	logger *log.Logger
	config *models.AppConfig
}

// NewMongoStore returns new instance of datastore
func NewMongoStore() MongoStore {
	return &mongoStore{
		txID:   uuid.New().String(),
		logger: log.New(os.Stdout, "database :=> ", log.LstdFlags),
		config: config.Config,
	}
}

// Client returns mongodb client instance
func (s *mongoStore) Client() *mongo.Client {
	once.Do(func() {
		db, client = s.initialize()
	})

	return client
}

// Database returns mongodb database instance
func (s *mongoStore) Database() *mongo.Database {
	once.Do(func() {
		db, client = s.initialize()
	})

	return db
}

func (s *mongoStore) initialize() (a *mongo.Database, b *mongo.Client) {
	client, err := mongo.NewClient(options.Client().ApplyURI(s.config.Database.Host))
	if err != nil {
		s.logger.Fatalf("%s %s Failed to connect to database with error: %v", err, s.txID, DBConnectionFailed)
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(s.config.Database.QueryTimeoutInSec)*time.Second,
	)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		s.logger.Fatalf("%s %s Failed to connect to database with error: %v", err, s.txID, DBConnectionFailed)
	}

	database := s.config.Database.Database
	db := client.Database(database)
	s.logger.Printf("%s Successfully connected to database %s", database, s.txID)

	return db, client
}

func (s *mongoStore) Health() *healthinterface.OutboundInterface {
	once.Do(func() {
		db, client = s.initialize()
	})

	outbound := healthinterface.OutboundInterface{}
	outbound.TimeStampUTC = time.Now().UTC()
	outbound.ConnectionStatus = healthinterface.ConnectionActive
	outbound.ApplicationName = "MongoDB"
	outbound.URLs = []string{s.config.Database.Host}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(s.config.Database.QueryTimeoutInSec)*time.Second,
	)
	defer cancel()

	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		outbound.ConnectionStatus = healthinterface.ConnectionDisconnected
	}

	return &outbound
}
