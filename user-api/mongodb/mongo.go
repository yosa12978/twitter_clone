package mongodb

import (
	"context"
	"sync"

	"github.com/yosa12978/twitter/user-api/configs"
	"github.com/yosa12978/twitter/user-api/logging"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db          *mongo.Database
	connectOnce sync.Once
	logger      logging.Logger
)

func Get(ctx context.Context) *mongo.Database {
	connectOnce.Do(func() {
		logger = logging.New("mongodb client")
		cfg := configs.Get()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoUri))
		if err != nil {
			logger.Fatalf(err.Error())
		}
		db = client.Database(cfg.MongoDbname)
	})
	return db
}
