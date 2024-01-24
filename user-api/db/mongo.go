package db

import (
	"context"
	"os"
	"sync"

	"github.com/yosa12978/twitter/user-api/logging"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db          *mongo.Database
	connectOnce sync.Once
	logger      logging.Logger
)

func GetDB(ctx context.Context) *mongo.Database {
	connectOnce.Do(func() {
		logger = logging.New("mongodb")
		uri := os.Getenv("MONGO_URI")
		if uri == "" {
			logger.Fatalf("MONGO_URI env variable is not provided herehere")
		}
		dbName := os.Getenv("MONGO_DBNAME")
		if dbName == "" {
			logger.Fatalf("MONGO_DBNAME env variable is not provided")
		}
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
		if err != nil {
			logger.Fatalf(err.Error())
		}
		db = client.Database(dbName)
	})
	return db
}
