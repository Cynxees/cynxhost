package persistentnodeusecase

import (
	"context"
	"cynxhost/internal/constant/types"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/model/request"
	"cynxhost/internal/model/response"
	"cynxhost/internal/model/response/responsecode"
	"cynxhost/internal/model/response/responsedata"
	"encoding/base64"
	"fmt"
	"strconv"
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

	// Set DNS
	dnsRecordId := persistentNode.DnsRecordId

	if dnsRecordId == nil {
		porkbunResp, err := usecase.porkbunManager.CreateDNS(persistentNode.ServerAlias, req.PublicIp)
		if err != nil {
			resp.Code = responsecode.CodePorkbunError
			resp.Error = "Error creating DNS: " + err.Error()
			return ctx
		}

		newId := strconv.Itoa(porkbunResp.Id)
		dnsRecordId = &newId

	} else {
		err = usecase.porkbunManager.UpdateDNS(*dnsRecordId, persistentNode.ServerAlias, req.PublicIp)
		if err != nil {
			resp.Code = responsecode.CodePorkbunError
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

	// Get Script
	ctx, serverTemplate, err := usecase.tblServerTemplate.GetServerTemplate(ctx, "id", strconv.Itoa(persistentNode.ServerTemplateId))
	if err != nil {
		resp.Code = responsecode.CodeTblServerTemplateError
		resp.Error = "Error getting server template: " + err.Error()
		return ctx
	}

	encodedScript := base64.StdEncoding.EncodeToString([]byte(serverTemplate.Script.SetupScript))

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
