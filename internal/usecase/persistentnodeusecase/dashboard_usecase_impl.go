package persistentnodeusecase

import (
	"context"
	"cynxhost/internal/model/request"
	"cynxhost/internal/model/response"
	"cynxhost/internal/model/response/responsecode"
	"fmt"
	"strconv"

	"github.com/gorcon/rcon"
)

func (usecase *PersistentNodeUseCaseImpl) SendCommandPersistentNode(ctx context.Context, req request.SendCommandPersistentNodeRequest, resp *response.APIResponse) context.Context {

	// Get persistent node
	ctx, persistentNodes, err := usecase.tblPersistentNode.GetPersistentNodes(ctx, "id", strconv.Itoa(req.PersistentNodeId))
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

	if persistentNode.OwnerId != req.SessionUser.Id {
		resp.Code = responsecode.CodeForbidden
		resp.Error = "You are not allowed to access this persistent node"
		return ctx
	}

	// Get instance
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

	// Prepare to send command
	rconHost := instance.PublicIp // TODO: Change to private IP
	rconPort := 25575             //
	rconPassword := "pass"        // TODO: implement pass in db

	// Connect to RCON
	rconClient, err := rcon.Dial(fmt.Sprintf("%s:%d", rconHost, rconPort), rconPassword)
	if err != nil {
		resp.Code = responsecode.CodeRCONError
		resp.Error = "Failed to connect to RCON: " + err.Error()
		return ctx
	}
	defer rconClient.Close()

	// Send command
	_, err = rconClient.Execute(req.Command)
	if err != nil {
		resp.Code = responsecode.CodeRCONError
		resp.Error = "Failed to send command: " + err.Error()
	}

	resp.Code = responsecode.CodeSuccess
	return ctx
}
