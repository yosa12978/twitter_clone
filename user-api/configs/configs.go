package configs

import "os"

type Config struct {
	Addr        string
	MongoUri    string
	MongoDbname string
}

var (
	config Config
)

func LoadConfig() Config {
	config = Config{
		Addr:        "0.0.0.0:5000",
		MongoUri:    "mongodb://localhost:27017",
		MongoDbname: "twitter-user-db",
	}
	if val, ok := os.LookupEnv("MONGO_URI"); ok {
		config.MongoUri = val
	}
	if val, ok := os.LookupEnv("MONGO_DBNAME"); ok {
		config.MongoDbname = val
	}
	if val, ok := os.LookupEnv("ADDR"); ok {
		config.Addr = val
	}
	return config
}

func Get() Config {
	return config
}
