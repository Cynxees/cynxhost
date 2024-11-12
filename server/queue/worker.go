package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"mchost-spot-instance/server/api"
	"mchost-spot-instance/server/models"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/redis/go-redis/v9"
	payload "mchost-spot-instance/server/queue/payload"
	ipPb "mchost-spot-instance/server/lib/stubs/mchost-ip"
)

func StartSpotInstanceWorker(server *api.Server) {
	go func() {
		ctx := context.Background()

		for {
			now := time.Now().Unix()

			// Fetch tasks that are due for processing
			onProvisionInstancePayloads, err := server.Redis.ZRangeByScore(ctx, "spot_instance_queue", &redis.ZRangeBy{
				Min: "-inf",
				Max: fmt.Sprintf("%d", now),
				Offset: 0,
				Count:  10,
			}).Result()
			if err != nil {
				server.Logger.Error("Failed to fetch from Redis:", err)
				time.Sleep(5 * time.Second)
				continue
			}

			for _, payloadStr := range onProvisionInstancePayloads {
				// Fetch template by FleetRequestId
				var payload payload.OnProvisionInstancePayload
				if err := json.Unmarshal([]byte(payloadStr), &payload); err != nil {
					server.Logger.Error("Failed to parse payload:", err)
					continue
				}

				var template models.SpotInstanceTemplate
				if err := server.Db.Where("fleet_request_id = ?", payload.FleetRequestId).First(&template).Error; err != nil {
					server.Logger.Error("Failed to find template:", err)
					continue
				}

				// Check if the Spot Fleet request has been fulfilled
				fleetInstances, err := server.AWSManager.EC2Client.DescribeSpotFleetInstances(ctx, &ec2.DescribeSpotFleetInstancesInput{
					SpotFleetRequestId: &payload.FleetRequestId,
				})
				if err != nil {
					server.Logger.Error("Failed to describe Spot Fleet instances:", err)
					continue
				}

				// If instances are found, update the database
				if len(fleetInstances.ActiveInstances) > 0 {
					instanceId := fleetInstances.ActiveInstances[0].InstanceId
					template.InstanceId = instanceId
					template.Status = "ACTIVE"

					if err := server.Db.Save(&template).Error; err != nil {
						server.Logger.Error("Failed to update template:", err)
						continue
					}

					// Log the instance details
					instanceDetails, _ := json.Marshal(fleetInstances)
					server.Logger.Info("Spot instance provisioned:", string(instanceDetails))

					if payload.EipAllocationId != nil {
						assignElasticIP(ctx, server, &assignIpData{
							InstanceId: *instanceId,
							EipAllocationId: *payload.EipAllocationId,
							IpId: template.IpId,
						});
					}

					// Remove the processed task from the queue
					server.Logger.Info("Removing task from queue:", payloadStr)
					server.Redis.ZRem(ctx, "spot_instance_queue", payloadStr)
				}
			}

			// Sleep for a short interval before checking again
			time.Sleep(10 * time.Second)
		}
	}()
}

type assignIpData struct {
	InstanceId     string
	EipAllocationId string
	IpId					 uint64
}

func assignElasticIP(
	ctx context.Context, 
	s *api.Server,
	data *assignIpData,	
	) error {	
	// eipAllocationID := "eipalloc-0d3131e17bfd77974"

	// Associate the Elastic IP with the instance
	_, err := s.AWSManager.EC2Client.AssociateAddress(ctx, &ec2.AssociateAddressInput{
		InstanceId:   aws.String(data.InstanceId),
		AllocationId: aws.String(data.EipAllocationId),
		AllowReassociation: aws.Bool(true),
	})

	microIpClient := *s.IpServiceClient
	microIpClient.UseIp(ctx, &ipPb.UseIpRequest{
		IpId: int64(data.IpId),
		InstanceId: data.InstanceId,
	})

	if err != nil {
		return fmt.Errorf("failed to associate Elastic IP: %w", err)
	}

	s.Logger.Info("Elastic IP associated with instance:", data.InstanceId)
	return nil
}