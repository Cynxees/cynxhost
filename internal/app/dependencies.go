package app

import (
	"cynxhost/internal/dependencies"
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

type Dependencies struct {
	Logger *logrus.Logger
	Config *dependencies.Config

	RedisClient    *redis.Client
	DatabaseClient *dependencies.DatabaseClient
	AWSClient *dependencies.AWSClient

	JWTManager *dependencies.JWTManager
}

func NewDependencies(configPath string) *Dependencies {

	config, err := dependencies.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}

	logger := dependencies.NewLogger(config)

	logger.Infoln("Connecting to Redis")
	redis := dependencies.NewRedisClient(config)

	logger.Infoln("Connecting to AWS")
	awsManager := dependencies.NewAWSClient(config.Aws.AccessKeyId, config.Aws.AccessKeySecret)

	logger.Infoln("Connecting to JWT")
	jwtManager := dependencies.NewJWTManager(config.App.Key, time.Hour*24)

	logger.Infoln("Connecting to Database")
	databaseClient, err := dependencies.NewDatabaseClient(config)
	if err != nil {
		logger.Fatalln("Failed to connect to database: ", err)
	}

	logger.Infoln("Dependencies initialized")
	return &Dependencies{
		Config:      config,
		DatabaseClient: databaseClient,
		RedisClient: redis,
		Logger:      logger,
		AWSClient:  awsManager,
		JWTManager:  jwtManager,
	}
}
