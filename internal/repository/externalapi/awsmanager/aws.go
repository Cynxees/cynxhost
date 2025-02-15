package awsmanager

import (
	"context"
	"cynxhost/internal/dependencies"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AWSManager struct {
	AppConfig *dependencies.ConfigAws
	Config    *aws.Config

	EC2Client *ec2.Client
	S3Client  *s3.Client
	ECRClient *ecr.Client
}

func NewAWSManager(appConfig *dependencies.ConfigAws) *AWSManager {

	config := newAWSConfig(appConfig.AccessKeyId, appConfig.AccessKeySecret)

	return &AWSManager{
		AppConfig: appConfig,
		Config:    config,
		EC2Client: newEC2Client(*config),
		S3Client:  newS3Client(*config),
		ECRClient: newECRClient(*config),
	}
}

func newEC2Client(config aws.Config) *ec2.Client {
	return ec2.NewFromConfig(config)
}

func newS3Client(config aws.Config) *s3.Client {
	return s3.NewFromConfig(config)
}

func newECRClient(config aws.Config) *ecr.Client {
	return ecr.NewFromConfig(config)
}

func newAWSConfig(accessKeyId string, secret string) *aws.Config {

	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				accessKeyId,
				secret,
				"", // Optional session token (use if MFA is enabled, otherwise leave it empty)
			),
		),
		config.WithRegion("ap-southeast-1"),
	)

	if err != nil {
		panic(err)
	}

	return &cfg
}

// GetSignedURLs generates signed URLs for a list of S3 keys in bulk.
func (client *AWSManager) GetSignedURL(key string) (*string, error) {
	presignClient := s3.NewPresignClient(client.S3Client)

	req, err := presignClient.PresignGetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(client.AppConfig.S3.Bucket),
		Key:    aws.String(key),
	}, func(options *s3.PresignOptions) {
		options.Expires = *aws.Duration(time.Minute * time.Duration(client.AppConfig.S3.Ttl))
	})
	if err != nil {
		return nil, err
	}

	return &req.URL, nil
}

func (client *AWSManager) GetUnsignedURL(key string) (*string, error) {
	return aws.String("https://" + client.AppConfig.S3.Bucket + ".s3." + client.AppConfig.Region + ".amazonaws.com/" + key), nil
}

func (manager *AWSManager) CreateEcrRepository(repositoryName string) error {

	_, err := manager.ECRClient.CreateRepository(context.Background(), &ecr.CreateRepositoryInput{
		RepositoryName: aws.String(repositoryName),
	})

	return err
}
