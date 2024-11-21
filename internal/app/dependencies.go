package app

import (
	"fmt"
	"mchost/internal/dependencies"
	"time"

	"github.com/sirupsen/logrus"
)

type Dependencies struct {
	Logger *logrus.Logger

	Config     *dependencies.Config
	AWSManager *dependencies.AWSManager
	JWTManager *dependencies.JWTManager
}

func NewDependencies(configPath string) *Dependencies {

	fmt.Println("loading config from: ", configPath)
	config, err := dependencies.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}

	logger := dependencies.NewLogger(config)

	awsManager := dependencies.NewAWSManager(config.Aws.AccessKeyId, config.Aws.AccessKeySecret)
	jwtManager := dependencies.NewJWTManager(config.App.Key, time.Hour*24, logger)

	return &Dependencies{
		Config:   	config,
		Logger:			logger,
		AWSManager: awsManager,
		JWTManager: jwtManager,
	}
}
