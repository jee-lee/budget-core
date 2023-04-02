package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"go.uber.org/zap"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var (
	Logger    *zap.Logger
	dbPool    *sqlx.DB
	oneDBPool sync.Once
)

func init() {
	// If appEnv is not set, set it to "DEVELOPMENT".
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "DEVELOPMENT"
	}
	if appEnv == "DEVELOPMENT" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	configureLogger(appEnv)
	Logger.Info("logger construction succeeded")
}

func configureLogger(appEnv string) {
	if appEnv == "TEST" || appEnv == "DEVELOPMENT" {
		Logger = zap.Must(zap.NewDevelopment())
		defer Logger.Sync()
		return
	}

	rawJSON := []byte(`{
	  "level": "info",
	  "encoding": "json",
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)
	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	Logger = zap.Must(cfg.Build())
	defer Logger.Sync()

}

// GetDB will return only one dbPool, no matter how many times it is called
func GetDB() *sqlx.DB {
	oneDBPool.Do(
		func() {
			var err error
			dbURL := mustGetEnv("DATABASE_URL")
			dbPool, err = sqlx.Open("postgres", dbURL)
			if err != nil {
				Logger.Fatal(err.Error())
			}
		},
	)
	return dbPool
}

// Always use mustGetEnv where we want to fail if an environment variable is not loaded
func mustGetEnv(key string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		msg := fmt.Sprintf("Cannot retrieve %v", key)
		Logger.Fatal(msg)
	}
	return
}
