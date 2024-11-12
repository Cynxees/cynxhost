package main

import (
	"fmt"
	"mchost-spot-instance/server/api"
	awsManager "mchost-spot-instance/server/aws"
	"mchost-spot-instance/server/config"
	controller "mchost-spot-instance/server/controller"
	jwtManager "mchost-spot-instance/server/jwt"
	queue "mchost-spot-instance/server/queue"
	ipPb "mchost-spot-instance/server/lib/stubs/mchost-ip"

	// "mchost-spot-instance/server/lib/rabbitmq"
	"mchost-spot-instance/server/models"
	"mchost-spot-instance/server/pb"
	"net"

	"mchost-spot-instance/www/docs"

	// elasticLog "gopkg.in/sohlich/elogrus.v7"

	"github.com/gin-gonic/gin"
	// "github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	fmt.Println("starting server")
	appConfig := config.InitConfig(".env")

	docs.SwaggerInfo.Title = "Spot Instance API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", appConfig.App.Address, appConfig.App.Port)
	docs.SwaggerInfo.BasePath = "/api"

	esLogger := logrus.New()
	esLogger.SetFormatter(&ecslogrus.Formatter{})
	esLogger.SetLevel(logrus.InfoLevel)

	// client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"), elastic.SetSniff(false))
	// if err != nil {
	// 	esLogger.Fatalf("Failed to create elasticsearch client: %v", err)
	// }

	// hook, err := elasticLog.NewAsyncElasticHook(client, serviceName, logrus.InfoLevel, "go-auth-logs")
	// if err != nil {
	// 	logrus.Fatalf("Failed to create Elasticsearch hook: %v", err)
	// }
	// esLogger.AddHook(hook)

	dsn := "user:password@tcp(127.0.0.1:" + appConfig.Db.Port + ")/mchost_spot_instance?charset=utf8mb4&parseTime=True&loc=Local"

	esLogger.Info(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		esLogger.Fatalf("failed to connect database: %v", err)
	}

	db.AutoMigrate(&models.SpotInstanceTemplate{})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", appConfig.App.MicroservicePort))
	if err != nil {
		esLogger.Fatalf("failed to listen: %v", err)
	}

	// rabbitmq.InitRabbitMq()

	// // Start the RabbitMQ consumer in a goroutine
	// go rabbitmq.StartConsumer("orders")

	ipServiceConn, err := grpc.Dial(appConfig.Microservice.Ip.Host+":"+appConfig.Microservice.Ip.Port, grpc.WithInsecure())
	if err != nil {
		esLogger.Fatalf("failed to connect to auth service: %v", err)
	}
	defer ipServiceConn.Close()

	ipServiceClient := ipPb.NewIpServiceClient(ipServiceConn)

	grpcServer := grpc.NewServer()
	server := &api.Server{
		Db:              db,
		Logger:          esLogger,
		JWTManager:      jwtManager.NewJWTManager(appConfig.App.Key, 3600, esLogger),
		AppConfig:       appConfig,
		AWSManager:      awsManager.NewAWSManager(appConfig.Aws.AccessKeyId, appConfig.Aws.AccessKeySecret),
		Redis:           queue.NewRedisClient(),
		IpServiceClient: &ipServiceClient,
	}

	queue.StartSpotInstanceWorker(server)

	router := gin.Default()

	controller.SetupHandlers(router, server)

	go router.Run(fmt.Sprintf(":%s", appConfig.App.Port))
	pb.RegisterSpotServiceServer(grpcServer, server)

	if err := grpcServer.Serve(lis); err != nil {
		esLogger.Fatalf("failed to serve: %v", err)
	}
}
