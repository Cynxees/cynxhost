package api

import (
	aws "mchost-spot-instance/server/aws"
	"mchost-spot-instance/server/config"
	manager "mchost-spot-instance/server/jwt"
	mchostIp "mchost-spot-instance/server/lib/stubs/mchost-ip"
	pb "mchost-spot-instance/server/pb"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Server struct {
	pb.SpotServiceServer
	Db         *gorm.DB
	Logger     *logrus.Logger
	JWTManager *manager.JWTManager
	AppConfig  *config.Config
	AWSManager *aws.AWSManager
	Redis 		 *redis.Client

	IpServiceClient *mchostIp.IpServiceClient
}
