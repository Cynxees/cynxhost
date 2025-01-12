package persistentnodeusecase

import (
	"context"
	"cynxhost/internal/constant/types"
	"cynxhost/internal/dependencies"
	"cynxhost/internal/dependencies/param"
	"cynxhost/internal/helper"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/model/request"
	"cynxhost/internal/model/response"
	"cynxhost/internal/model/response/responsecode"
	"cynxhost/internal/model/response/responsedata"
	"cynxhost/internal/repository/database"
	"cynxhost/internal/usecase"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	awstypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/sirupsen/logrus"
)

type PersistentNodeUseCaseImpl struct {
	tblPersistentNode database.TblPersistentNode
	tblInstance       database.TblInstance
	tblInstanceType   database.TblInstanceType
	tblStorage        database.TblStorage
	tblServerTemplate database.TblServerTemplate

	awsClient *dependencies.AWSClient
	log       *logrus.Logger
	config    *dependencies.Config
}

func New(tblPersistentNode database.TblPersistentNode, tblInstance database.TblInstance, tblInstanceType database.TblInstanceType, tblStorage database.TblStorage, tblServerTemplate database.TblServerTemplate, awsClient *dependencies.AWSClient, logger *logrus.Logger, config *dependencies.Config) usecase.PersistentNodeUseCase {

	return &PersistentNodeUseCaseImpl{
		tblPersistentNode: tblPersistentNode,
		tblStorage:        tblStorage,
		tblServerTemplate: tblServerTemplate,
		tblInstance:       tblInstance,
		tblInstanceType:   tblInstanceType,

		awsClient: awsClient,
		log:       logger,
		config:    config,
	}
}

func (usecase *PersistentNodeUseCaseImpl) GetPersistentNodes(ctx context.Context, resp *response.APIResponse) {

	contextUser, ok := helper.GetUserFromContext(ctx)
	if !ok {
		resp.Code = responsecode.CodeAuthenticationError
		resp.Error = "User not found in context"
		return
	}

	_, persistentNodes, err := usecase.tblPersistentNode.GetPersistentNodes(ctx, "owner_id", strconv.Itoa(contextUser.Id))
	if err != nil {
		resp.Code = responsecode.CodeTblPersistentNodeError
		resp.Error = err.Error()
		return
	}

	resp.Code = responsecode.CodeSuccess
	resp.Data = responsedata.PaginatePersistentNodeResponseData{
		PersistentNodes: persistentNodes,
	}
}

func (usecase *PersistentNodeUseCaseImpl) GetPersistentNode(ctx context.Context, req request.GetPersistentNodeRequest, resp *response.APIResponse) context.Context {
	_, persistentNodes, err := usecase.tblPersistentNode.GetPersistentNodes(ctx, "id", strconv.Itoa(req.PersistentNodeId))
	if err != nil {
		resp.Code = responsecode.CodeTblPersistentNodeError
		resp.Error = err.Error()
		return ctx
	}

	if len(persistentNodes) == 0 {
		resp.Code = responsecode.CodeNotFound
		resp.Error = "Persistent node not found"
		return ctx
	}

	contextUser, ok := helper.GetUserFromContext(ctx)
	if !ok {
		resp.Code = responsecode.CodeAuthenticationError
		resp.Error = "User not found in context"
		return ctx
	}

	persistentNode := persistentNodes[0]
	if persistentNode.OwnerId != contextUser.Id {
		resp.Code = responsecode.CodeForbidden
		resp.Error = "You are not allowed to access this persistent node"
		return ctx
	}

	ctx = helper.SetVisibilityLevelToContext(ctx, types.VisibilityLevelPrivate)

	resp.Code = responsecode.CodeSuccess
	resp.Data = persistentNode
	return ctx
}

func (usecase *PersistentNodeUseCaseImpl) CreatePersistentNode(ctx context.Context, req request.CreatePersistentNodeRequest, resp *response.APIResponse) context.Context {

	contextUser, ok := helper.GetUserFromContext(ctx)
	if !ok {
		resp.Code = responsecode.CodeAuthenticationError
		resp.Error = "User not found in context"
		return ctx
	}

	// Get instance type
	_, instanceType, err := usecase.tblInstanceType.GetInstanceType(ctx, "id", strconv.Itoa(req.InstanceTypeId))
	if err != nil {
		resp.Code = responsecode.CodeTblInstanceTypeError
		resp.Error = err.Error()
		return ctx
	}

	userDataVariables := map[string]string{
		"LAUNCH_SUCCESS_CALLBACK_URL": fmt.Sprintf("http://%s/api/v1/persistent-node/callback/launch", usecase.config.App.PrivateIp),
		"LAUNCH_SUCCESS_TYPE":         string(types.LaunchCallbackPersistentNodeTypeInitialLaunch),
		"SETUP_SUCCESS_CALLBACK_URL":  fmt.Sprintf("http://%s/api/v1/persistent-node/callback/update-status", usecase.config.App.PrivateIp),
		"SETUP_SUCCESS_TYPE":          string(types.SetupSuccessCallbackPersistentNodeType),
	}

	userData, err := helper.ReplacePlaceholders(string(param.StaticParam.ParamAwsLaunchScript), userDataVariables)
	if err != nil {
		resp.Code = responsecode.CodeInternalError
		resp.Error = err.Error()
		return ctx
	}
	usecase.log.Infoln("User data: ", userData)
	encodedUserData := base64.StdEncoding.EncodeToString([]byte(userData))

	usecase.log.Infoln("encoded user data: ", encodedUserData)
	ec2RunInstanceInput := &ec2.RunInstancesInput{
		MinCount: aws.Int32(1),
		MaxCount: aws.Int32(1),
		IamInstanceProfile: &awstypes.IamInstanceProfileSpecification{
			Arn: aws.String("arn:aws:iam::242201306378:instance-profile/cynxhost-node-iam"),
		},
		BlockDeviceMappings: []awstypes.BlockDeviceMapping{
			{
				DeviceName: aws.String("/dev/sda1"), // The device name, typically /dev/sda1 for the root volume
				Ebs: &awstypes.EbsBlockDevice{
					DeleteOnTermination: aws.Bool(true),         // Ensures that the volume is deleted when the instance is terminated
					VolumeSize:          aws.Int32(8),           // Set the volume size in GiB (e.g., 20 GiB)
					VolumeType:          awstypes.VolumeTypeGp2, // You can also specify the volume type, such as gp2 (General Purpose SSD)
				},
			},
			{
				DeviceName: aws.String("/dev/sdb"),
				Ebs: &awstypes.EbsBlockDevice{
					DeleteOnTermination: aws.Bool(true), // TODO: Change to false
					VolumeSize:          aws.Int32(8),
					VolumeType:          awstypes.VolumeTypeGp2,
				},
			},
		},
		ImageId:      aws.String(param.StaticParam.ParamAwsNodeId.AmiId),
		KeyName:      aws.String(param.StaticParam.ParamAwsNodeId.KeyPairName),
		InstanceType: awstypes.InstanceType(instanceType.Name),
		InstanceMarketOptions: &awstypes.InstanceMarketOptionsRequest{
			MarketType: awstypes.MarketTypeSpot,
			SpotOptions: &awstypes.SpotMarketOptions{
				InstanceInterruptionBehavior: awstypes.InstanceInterruptionBehaviorTerminate,
			},
		},
		NetworkInterfaces: []awstypes.InstanceNetworkInterfaceSpecification{
			{
				AssociatePublicIpAddress: aws.Bool(true),
				DeviceIndex:              aws.Int32(0),
				Groups: []string{
					param.StaticParam.ParamAwsNodeId.SecurityGroupId, // Security group defined inside the network interface
				},
			},
		},
		UserData: aws.String(encodedUserData),
	}

	// Create instance in aws
	ec2RunInstanceOutput, err := usecase.awsClient.EC2Client.RunInstances(ctx, ec2RunInstanceInput)
	if err != nil {
		resp.Code = responsecode.CodeEC2Error
		resp.Error = err.Error()
		return ctx
	}

	if len(ec2RunInstanceOutput.Instances) == 0 {
		resp.Code = responsecode.CodeEC2Error
		resp.Error = "No instances created"
		return ctx
	}

	createdEc2 := ec2RunInstanceOutput.Instances[0]
	data, err := json.Marshal(createdEc2)
	if err != nil {
		resp.Code = responsecode.CodeInternalError
		resp.Error = err.Error()
		return ctx
	}
	usecase.log.Infoln("Created instance: ", string(data))

	storage := entity.TblStorage{
		Name:   req.Name,
		SizeMb: req.StorageSizeMb,
		Status: types.StorageStatusNew,
	}

	instance := entity.TblInstance{
		Name:           req.Name,
		AwsInstanceId:  *createdEc2.InstanceId,
		PrivateIp:      *createdEc2.PrivateIpAddress,
		InstanceTypeId: req.InstanceTypeId,
		Status:         types.InstanceStatusCreate,
	}

	ctx, storageId, err := usecase.tblStorage.CreateStorage(ctx, storage)
	if err != nil {
		resp.Code = responsecode.CodeTblStorageError
		resp.Error = err.Error()
		return ctx
	}

	ctx, instanceId, err := usecase.tblInstance.CreateInstance(ctx, instance)
	if err != nil {
		resp.Code = responsecode.CodeTblInstanceError
		resp.Error = err.Error()
		return ctx
	}

	persistentNode := entity.TblPersistentNode{
		Name:             req.Name,
		OwnerId:          contextUser.Id,
		ServerTemplateId: req.ServerTemplateId,
		InstanceTypeId:   req.InstanceTypeId,
		StorageId:        storageId,
		InstanceId:       instanceId,
		Status:           types.PersistentNodeStatusCreating,
	}
	ctx, _, err = usecase.tblPersistentNode.CreatePersistentNode(ctx, persistentNode)
	if err != nil {
		resp.Code = responsecode.CodeTblPersistentNodeError
		resp.Error = err.Error()
		return ctx
	}

	resp.Code = responsecode.CodeSuccess
	return ctx
}

func (usecase *PersistentNodeUseCaseImpl) RunPersistentNodeScript(ctx context.Context, req request.RunPersistentNodeScriptRequest, resp *response.APIResponse) {

}
