package app

import (
	"cynxhost/internal/dependencies"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

type Dependencies struct {
	Logger *logrus.Logger
	Config *dependencies.Config

	Validator *validator.Validate

	RedisClient    *redis.Client
	DatabaseClient *dependencies.DatabaseClient

	JWTManager *dependencies.JWTManager
}

func NewDependencies(configPath string) *Dependencies {

	log.Println("Loading Config")
	config, err := dependencies.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}

	log.Println("Initializing Logger")
	logger := dependencies.NewLogger(config)

	logger.Infoln("Initializing Validator")
	validator := validator.New()

	logger.Infoln("Connecting to Redis")
	redis := dependencies.NewRedisClient(config)

	logger.Infoln("Connecting to JWT")
	jwtManager := dependencies.NewJWTManager(config.App.Key, time.Hour*24)

	logger.Infoln("Connecting to Database")
	databaseClient, err := dependencies.NewDatabaseClient(config)
	if err != nil {
		logger.Fatalln("Failed to connect to database: ", err)
	}

	logger.Infoln("Dependencies initialized")
	return &Dependencies{
		Config:         config,
		DatabaseClient: databaseClient,
		Validator:      validator,
		RedisClient:    redis,
		Logger:         logger,
		JWTManager:     jwtManager,
	}
}
