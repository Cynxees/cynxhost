package persistentnodeusecase

import (
	"context"
	"cynxhost/internal/constant/types"
	"cynxhost/internal/helper"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/model/request"
	"cynxhost/internal/model/response"
	"cynxhost/internal/model/response/responsecode"
	"cynxhost/internal/model/response/responsedata"
	"encoding/base64"
	"fmt"
	"strconv"

	awstypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func (usecase *PersistentNodeUseCaseImpl) LaunchCallbackPersistentNode(ctx context.Context, req request.LaunchCallbackPersistentNodeRequest, resp *response.APIResponse) context.Context {

	// Get instances
	ctx, instances, err := usecase.tblInstance.GetInstances(ctx, "aws_instance_id", req.AwsInstanceId)
	if err != nil {
		resp.Code = responsecode.CodeTblInstanceError
		resp.Error = "Error getting instance: " + err.Error()
		return ctx
	}

	if len(instances) == 0 {
		resp.Code = responsecode.CodeNotFound
		resp.Error = "Instance not found"
		return ctx
	}

	instance := instances[0]

	// Compare IP
	if req.ClientIp != instance.PublicIp && req.ClientIp != instance.PrivateIp {
		resp.Code = responsecode.CodeForbidden
		resp.Error = fmt.Sprintf("Client IP does not match instance IP: %s", req.ClientIp)
		return ctx
	}

	ctx, persistentNodes, err := usecase.tblPersistentNode.GetPersistentNodes(ctx, "instance_id", strconv.Itoa(instance.Id))
	if err != nil {
		resp.Code = responsecode.CodeTblPersistentNodeError
		resp.Error = "Error getting persistent node: " + err.Error()
		return ctx
	}

	if len(persistentNodes) == 0 {
		resp.Code = responsecode.CodeNotFound
		resp.Error = "Persistent node not found"
		return ctx
	}

	persistentNode := persistentNodes[0]

	variables, err := persistentNode.Variables.ToScriptVariables()
	if err != nil {
		resp.Code = responsecode.CodeFailJSON
		resp.Error = err.Error()
		return ctx
	}

	result, err := helper.FormatScriptVariables(persistentNode.ServerTemplate.Script.Variables, variables)
	if err != nil {
		resp.Code = responsecode.CodeFailJSON
		resp.Error = err.Error()
		return ctx
	}

	script, err := helper.ReplacePlaceholders(string(persistentNode.ServerTemplate.Script.SetupScript), result)
	encodedScript := base64.StdEncoding.EncodeToString([]byte(script))

	// Set DNS
	dnsRecordId := persistentNode.DnsRecordId

	if dnsRecordId == nil {
		cloudflareResp, err := usecase.cloudflareManager.CreateDNSRecord(persistentNode.ServerAlias, req.PublicIp)
		if err != nil {
			resp.Code = responsecode.CodeCloudflareError
			resp.Error = "Error creating DNS: " + err.Error()
			return ctx
		}

		newId := cloudflareResp.Result.ID
		dnsRecordId = &newId

	} else {
		err = usecase.cloudflareManager.UpdateDNS(*dnsRecordId, "AAAA", persistentNode.ServerAlias, req.PublicIp)
		if err != nil {
			resp.Code = responsecode.CodeCloudflareError
			resp.Error = "Error updating DNS: " + err.Error()
			return ctx
		}
	}

	// Update persistent node
	ctx, err = usecase.tblPersistentNode.UpdatePersistentNode(ctx, persistentNode.Id, entity.TblPersistentNode{
		Status:      types.PersistentNodeStatusSetup,
		DnsRecordId: dnsRecordId,
	})
	if err != nil {
		resp.Code = responsecode.CodeTblPersistentNodeError
		resp.Error = "Error updating persistent node: " + err.Error()
		return ctx
	}

	// Update Instance
	ctx, err = usecase.tblInstance.UpdateInstance(ctx, instance.Id, entity.TblInstance{
		Status:   types.InstanceStatusActive,
		PublicIp: req.PublicIp,
	})
	if err != nil {
		resp.Code = responsecode.CodeTblInstanceError
		resp.Error = "Error updating instance: " + err.Error()
	}

	// Update Storage
	ctx, err = usecase.tblStorage.UpdateStorage(ctx, persistentNode.StorageId, entity.TblStorage{
		AwsEbsId: req.EbsVolumeId,
	})
	if err != nil {
		resp.Code = responsecode.CodeTblStorageError
		resp.Error = "Error updating storage: " + err.Error()
	}

	resp.Code = responsecode.CodeSuccess
	resp.Data = responsedata.LaunchCallbackPersistentNodeResponseData{
		PersistentNodeId: persistentNode.Id,
		Script:           encodedScript,
	}

	return ctx
}

func (usecase *PersistentNodeUseCaseImpl) StatusCallbackPersistentNode(ctx context.Context, req request.StatusCallbackPersistentNodeRequest, resp *response.APIResponse) context.Context {

	// Get persistentNodes
	ctx, persistentNodes, err := usecase.tblPersistentNode.GetPersistentNodes(ctx, "id", strconv.Itoa(req.PersistentNodeId))
	if err != nil {
		resp.Code = responsecode.CodeTblPersistentNodeError
		resp.Error = "Error getting persistent node: " + err.Error()
	}

	if len(persistentNodes) == 0 {
		resp.Code = responsecode.CodeNotFound
		resp.Error = "Persistent node not found"
	}

	persistentNode := persistentNodes[0]

	// Get instances
	ctx, instances, err := usecase.tblInstance.GetInstances(ctx, "id", strconv.Itoa(*persistentNode.InstanceId))
	if err != nil {
		resp.Code = responsecode.CodeTblInstanceError
		resp.Error = "Error getting instance: " + err.Error()
		return ctx
	}

	if len(instances) == 0 {
		resp.Code = responsecode.CodeNotFound
		resp.Error = "Instance not found"
		return ctx
	}

	instance := instances[0]

	// Compare IP
	if req.ClientIp != instance.PublicIp && req.ClientIp != instance.PrivateIp {
		resp.Code = responsecode.CodeForbidden
		resp.Error = fmt.Sprintf("Client IP does not match instance IP: %s", req.ClientIp)
		return ctx
	}

	// Update persistent node
	ctx, err = usecase.tblPersistentNode.UpdatePersistentNode(ctx, persistentNode.Id, entity.TblPersistentNode{
		Status: types.PersistentNodeStatusRunning,
	})
	if err != nil {
		resp.Code = responsecode.CodeTblPersistentNodeError
		resp.Error = "Error updating persistent node: " + err.Error()
	}

	resp.Code = responsecode.CodeSuccess
	return ctx
}

func (usecase *PersistentNodeUseCaseImpl) ShutdownCallbackPersistentNode(ctx context.Context, req request.ShutdownCallbackPersistentNodeRequest, resp *response.APIResponse) context.Context {

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

	if persistentNode.Instance == nil {
		resp.Code = responsecode.CodeNotFound
		resp.Error = "No Instance Running instance found"

		return ctx
	}

	// Check if the persistent node ip is the same
	if persistentNode.Instance.PrivateIp != req.ClientIp && persistentNode.Instance.PublicIp != req.ClientIp {
		resp.Code = responsecode.CodeForbidden
		resp.Error = "You are not allowed to access this persistent node"

		return ctx
	}

	// Shutdown the instance
	terminatedInstance, err := usecase.shutdownInstance(ctx, &persistentNode)
	if err != nil {
		resp.Code = responsecode.CodeEC2Error
		resp.Error = err.Error()

		return ctx
	}

	if terminatedInstance.CurrentState.Name == awstypes.InstanceStateNameTerminated || terminatedInstance.CurrentState.Name == awstypes.InstanceStateNameShuttingDown {
		resp.Code = responsecode.CodeNotAllowed
		resp.Error = "Instance already terminated or is shutting down"
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
