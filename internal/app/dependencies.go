package app

import (
	"cynxhost/internal/dependencies"
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

type Dependencies struct {
	Logger *logrus.Logger
	redis  *redis.Client

	Config     *dependencies.Config
	AWSManager *dependencies.AWSManager
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
	awsManager := dependencies.NewAWSManager(config.Aws.AccessKeyId, config.Aws.AccessKeySecret)

	logger.Infoln("Connecting to JWT")
	jwtManager := dependencies.NewJWTManager(config.App.Key, time.Hour*24, logger)

	return &Dependencies{
		Config:     config,
		redis:      redis,
		Logger:     logger,
		AWSManager: awsManager,
		JWTManager: jwtManager,
	}
}
