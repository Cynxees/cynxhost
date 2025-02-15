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
	"cynxhost/internal/repository/externalapi/awsmanager"
	"cynxhost/internal/repository/externalapi/porkbunmanager"
	"cynxhost/internal/usecase"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

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

	log    *logrus.Logger
	config *dependencies.Config

	awsManager     *awsmanager.AWSManager
	porkbunManager *porkbunmanager.PorkbunManager
}

func New(tblPersistentNode database.TblPersistentNode, tblInstance database.TblInstance, tblInstanceType database.TblInstanceType, tblStorage database.TblStorage, tblServerTemplate database.TblServerTemplate, awsManager *awsmanager.AWSManager, logger *logrus.Logger, config *dependencies.Config, porkbunManager *porkbunmanager.PorkbunManager) usecase.PersistentNodeUseCase {

	return &PersistentNodeUseCaseImpl{
		tblPersistentNode: tblPersistentNode,
		tblStorage:        tblStorage,
		tblServerTemplate: tblServerTemplate,
		tblInstance:       tblInstance,
		tblInstanceType:   tblInstanceType,

		log:    logger,
		config: config,

		awsManager:     awsManager,
		porkbunManager: porkbunManager,
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

func (usecase *PersistentNodeUseCaseImpl) GetPersistentNode(ctx context.Context, req request.IdRequest, resp *response.APIResponse) context.Context {
	_, persistentNodes, err := usecase.tblPersistentNode.GetPersistentNodes(ctx, "id", strconv.Itoa(req.Id))
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

	// Get Script
	_, serverTemplate, err := usecase.tblServerTemplate.GetServerTemplate(ctx, "id", strconv.Itoa(req.ServerTemplateId))
	if err != nil {
		resp.Code = responsecode.CodeTblServerTemplateError
		resp.Error = err.Error()
		return ctx
	}

	if serverTemplate == nil {
		resp.Code = responsecode.CodeNotFound
		resp.Error = "Server template not found"
		return ctx
	}

	// Change struct to map
	marshalledVariables, err := helper.StructToMapStringArray(req.Variables)
	if err != nil {
		resp.Code = responsecode.CodeFailJSON
		resp.Error = err.Error()
		return ctx
	}

	// Check if server variable is valid
	_, err = helper.FormatScriptVariables(serverTemplate.Script.Variables, req.Variables)
	if err != nil {
		resp.Code = responsecode.CodeFailJSON
		resp.Error = err.Error()
		return ctx
	}

	hash := fmt.Sprintf("%d-%d-%s", contextUser.Id, req.InstanceTypeId, req.Name)
	callbackBaseUrl := fmt.Sprintf("%s:%d", usecase.config.App.PrivateIp, usecase.config.App.Port)

	userDataVariables := map[string]string{
		"LAUNCH_SUCCESS_CALLBACK_URL": fmt.Sprintf("http://%s/api/v1/persistent-node/callback/launch", callbackBaseUrl),
		"LAUNCH_SUCCESS_TYPE":         string(types.LaunchCallbackPersistentNodeTypeInitialLaunch),
		"SETUP_SUCCESS_CALLBACK_URL":  fmt.Sprintf("http://%s/api/v1/persistent-node/callback/update-status", callbackBaseUrl),
		"SETUP_SUCCESS_TYPE":          string(types.SetupSuccessCallbackPersistentNodeType),
		"JWT_SECRET":                  hash,
		"CENTRAL_PRIVATE_IP":          usecase.config.App.PrivateIp,
		"CENTRAL_PUBLIC_IP":           usecase.config.App.PublicIp,
		"CENTRAL_PORT":                strconv.Itoa(usecase.config.App.Port),
		"DOMAIN":                      fmt.Sprintf("%s.%s", req.ServerAlias, usecase.config.Porkbun.Domain),
		"CONFIG_PATH":                 serverTemplate.Script.ConfigPath,
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
					DeleteOnTermination: aws.Bool(true),
					VolumeSize:          aws.Int32(int32(req.StorageSizeGb)),
					VolumeType:          awstypes.VolumeTypeGp2,
				},
			},
		},
		ImageId:      aws.String(param.StaticParam.ParamAwsNodeId.AmiId),
		KeyName:      aws.String(param.StaticParam.ParamAwsNodeId.KeyPairName),
		InstanceType: awstypes.InstanceType(instanceType.AwsKey),
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
	ec2RunInstanceOutput, err := usecase.awsManager.EC2Client.RunInstances(ctx, ec2RunInstanceInput)
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
		SizeMb: req.StorageSizeGb,
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
		InstanceId:       &instanceId,
		Status:           types.PersistentNodeStatusCreating,
		ServerAlias:      req.ServerAlias,
		Variables:        marshalledVariables,
	}
	ctx, _, err = usecase.tblPersistentNode.CreatePersistentNode(ctx, persistentNode)
	if err != nil {
		resp.Code = responsecode.CodeTblPersistentNodeError
		resp.Error = err.Error()
		return ctx
	}

	// Create ecr repository
	err = usecase.awsManager.CreateEcrRepository(fmt.Sprintf("cynxhost-%d", persistentNode.Id))
	if err != nil && !strings.Contains(err.Error(), "RepositoryAlreadyExistsException") {
		resp.Code = responsecode.CodeECRError
		resp.Error = err.Error()
		return ctx
	}

	resp.Code = responsecode.CodeSuccess
	return ctx
}

func (usecase *PersistentNodeUseCaseImpl) ForceShutdownPersistentNode(ctx context.Context, req request.ForceShutdownPersistentNodeRequest, resp *response.APIResponse) context.Context {

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

	persistentNode := persistentNodes[0]

	// Check if owner is the same
	if persistentNode.OwnerId != req.SessionUser.Id {
		resp.Code = responsecode.CodeForbidden
		resp.Error = "You are not allowed to access this persistent node"

		return ctx
	}

	if persistentNode.Instance == nil {
		resp.Code = responsecode.CodeNotFound
		resp.Error = "No Instance Running instance found"

		return ctx
	}

	ctx, err = usecase.tblPersistentNode.UpdatePersistentNode(ctx, persistentNode.Id, entity.TblPersistentNode{
		Status: types.PersistentNodeStatusShuttingDown,
	})

	// Shutdown the instance
	terminatedInstance, err := usecase.shutdownInstance(ctx, &persistentNode)
	if err != nil {
		resp.Code = responsecode.CodeEC2Error
		resp.Error = err.Error()

		return ctx
	}

	if terminatedInstance.CurrentState.Name == awstypes.InstanceStateNameTerminated {
		resp.Code = responsecode.CodeNotAllowed
		resp.Error = "Instance already terminated"
		return ctx
	}

	// Remove instance from persistent node
	ctx, err = usecase.tblPersistentNode.UpdatePersistentNode(ctx, persistentNode.Id, entity.TblPersistentNode{
		InstanceId: nil,
		Status:     types.PersistentNodeStatusShutdown,
	})

	if err != nil {
		resp.Code = responsecode.CodeTblPersistentNodeError
		resp.Error = err.Error()
		return ctx
	}

	// Delete the instance
	if persistentNode.InstanceId == nil {
		resp.Code = responsecode.CodeNotFound
		resp.Error = "No instance found"
		return ctx
	}

	ctx, err = usecase.tblInstance.DeleteInstance(ctx, *persistentNode.InstanceId)
	if err != nil {
		resp.Code = responsecode.CodeTblInstanceError
		resp.Error = err.Error()
		return ctx
	}

	// Change the status of the persistent node ( TODO )

	resp.Code = responsecode.CodeSuccess
	return ctx
}

func (usecase *PersistentNodeUseCaseImpl) shutdownInstance(ctx context.Context, persistentNode *entity.TblPersistentNode) (*awstypes.InstanceStateChange, error) {

	response, err := usecase.awsManager.EC2Client.TerminateInstances(ctx, &ec2.TerminateInstancesInput{
		InstanceIds: []string{persistentNode.Instance.AwsInstanceId},
	})

	if len(response.TerminatingInstances) == 0 {
		return nil, fmt.Errorf("No instances terminated")
	}

	terminatedInstance := response.TerminatingInstances[0]

	// Update DNS
	usecase.porkbunManager.UpdateDNS("A", *&persistentNode.ServerAlias, usecase.config.App.PublicIp)

	return &terminatedInstance, err
}
