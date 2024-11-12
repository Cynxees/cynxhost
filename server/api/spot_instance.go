package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/ssh"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/types/known/timestamppb"

	microIp "mchost-spot-instance/server/lib/stubs/mchost-ip"
	models "mchost-spot-instance/server/models"
	pb "mchost-spot-instance/server/pb"
	payload "mchost-spot-instance/server/queue/payload"
)

func (s *Server) CreateTemplate(ctx context.Context, request *pb.CreateTemplateRequest) (*pb.GetTemplateResponse, error) {

  s.Logger.Info("Creating Template")
	spotInstanceTemplate := &models.SpotInstanceTemplate{
		FleetRequestId: nil,
		InstanceId: nil,
		UserId: 1,
		Name: request.Name,
		Status: "PENDING",
		InstanceType: "t2.TODO",
		IpId: request.IpId,
		Region: "ap-southeast-3",
	}

	if err := s.Db.Create(spotInstanceTemplate).Error; err != nil {
		return nil, err
	}

	microIpClient := *s.IpServiceClient

	s.Logger.Info("Reserving IP: ")
	_, err := microIpClient.ReserveIp(ctx, &microIp.ReserveIpRequest{
		IpId: int64(request.IpId),
		SpotInstanceTemplateId: uint64(spotInstanceTemplate.Id),
	})

	if err != nil {
		return nil, err
	}
	
	s.Logger.Info("Ip Reserved")
	return &pb.GetTemplateResponse{
		Error: false,
		Code: http.StatusOK,
		Message: "Success",
		Template: &pb.SpotInstanceTemplate{
			Id: uint64(spotInstanceTemplate.Id),
			UserId: uint64(spotInstanceTemplate.UserId),
			Name: spotInstanceTemplate.Name,
			Status: spotInstanceTemplate.Status,
			InstanceType: spotInstanceTemplate.InstanceType,
			CreatedAt: timestamppb.New(spotInstanceTemplate.CreatedAt),
			UpdatedAt: timestamppb.New(spotInstanceTemplate.UpdatedAt),
			AmiId: spotInstanceTemplate.AmiId,
			IpId: uint64(spotInstanceTemplate.IpId),
			Region: spotInstanceTemplate.Region,
		},
	}, nil
}

func (s *Server) LaunchSpotFleet(ctx context.Context, request *pb.LaunchTemplateRequest) (*ec2.RequestSpotFleetOutput, error) {

  result, err := s.GetTemplate(ctx, &pb.GetTemplateRequest{SpotInstanceTemplateId: request.SpotInstanceTemplateId});
  if err != nil {
    return nil, err
  }

  template := result.Template

  if(template.Status == "ACTIVE") {
    return nil, errors.New("template is already active")
  }

	client := s.AWSManager.EC2Client

  s.Logger.Info(template)

  // Basicly user ubuntu with password "ubuntu:password" and run the following script
  base64UserData := "IyEvYmluL2Jhc2gKZWNobyAidWJ1bnR1OnBhc3N3b3JkIiB8IHN1ZG8gY2hwYXNzd2QKc2VkIC1pICJzL15QYXNzd29yZEF1dGhlbnRpY2F0aW9uIG5vL1Bhc3N3b3JkQXV0aGVudGljYXRpb24geWVzLyIgL2V0Yy9zc2gvc3NoZF9jb25maWcKc3lzdGVtY3RsIHJlc3RhcnQgc3NoZAo="

	spotRequestInput := &ec2.RequestSpotFleetInput{
		SpotFleetRequestConfig: &types.SpotFleetRequestConfigData{

			IamFleetRole:   aws.String("arn:aws:iam::071412439153:role/aws-ec2-spot-fleet-tagging-role"),
			TargetCapacity: aws.Int32(1),
      TagSpecifications: []types.TagSpecification{
        {
          ResourceType: types.ResourceTypeSpotFleetRequest,
          Tags: []types.Tag{
            {
              Key: aws.String("project"),
              Value: aws.String("mchost"),
            },
          },
        },
      },

			InstanceInterruptionBehavior: types.InstanceInterruptionBehaviorStop,
			LaunchSpecifications: []types.SpotFleetLaunchSpecification{
				{
					ImageId:      aws.String(template.AmiId),
					InstanceType: types.InstanceTypeT32xlarge,
					KeyName:      aws.String("minecraft-server"),
					SecurityGroups: []types.GroupIdentifier{
						{
							GroupId: aws.String("sg-06f8c1349d5087902"),
						},
					},

					IamInstanceProfile: &types.IamInstanceProfileSpecification{
						Arn: aws.String("arn:aws:iam::071412439153:instance-profile/EC2-S3-FullAccess"),
					},
          UserData: aws.String(base64UserData),
				},
			},
		},
	}

	fleetRequest, err := client.RequestSpotFleet(ctx, spotRequestInput)
	if err != nil {
		s.Logger.Fatal(err)
	}

  if err := s.Db.Model(&models.SpotInstanceTemplate{}).
    Where("id = ?", request.SpotInstanceTemplateId).
    UpdateColumns(map[string] interface{}{
      "fleet_request_id": fleetRequest.SpotFleetRequestId,
      "status": "ACTIVE",
    }).Error; err != nil {
    return nil, err
    }

	s.Logger.Info(fleetRequest)

	microIpClient := *s.IpServiceClient

	ipResp, err := microIpClient.GetIp(ctx, &microIp.GetIpRequest{
	  IpId: int64(template.IpId),
  })

	if err != nil {
		return nil, err
	}

	payload := payload.OnProvisionInstancePayload{
		FleetRequestId:  *fleetRequest.SpotFleetRequestId,
		EipAllocationId: &ipResp.Ip.AllocationId,
	}
	
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		s.Logger.Error("Failed to marshal payload:", err)
		return nil, err
	}

	delay := time.Now().Add(20 * time.Second).Unix()
	err = s.Redis.ZAdd(ctx, "spot_instance_queue", redis.Z{
		Score: float64(delay),
		Member: payloadJSON,
	}).Err()
	if err != nil {
		return nil, errors.New("failed to push to queue");
	}

	return fleetRequest, nil
}

func (s *Server) GetTemplate(ctx context.Context, request *pb.GetTemplateRequest) (*pb.GetTemplateResponse, error) {
  
  template := &models.SpotInstanceTemplate{}
  if err := s.Db.Where("id = ?", request.SpotInstanceTemplateId).First(template).Error; err != nil {
    return nil, err
  }

  if template.FleetRequestId != nil {

    //TODO: maybe redundant
    fleetRequest, err := s.AWSManager.EC2Client.DescribeSpotFleetRequests(ctx, &ec2.DescribeSpotFleetRequestsInput{
      SpotFleetRequestIds: []string{*template.FleetRequestId},
    })
    if err != nil {
      return nil, err
    }

    if len(fleetRequest.SpotFleetRequestConfigs) > 0 {

      s.Logger.Info("getting instances")
      instances, err := s.AWSManager.EC2Client.DescribeSpotFleetInstances(ctx, &ec2.DescribeSpotFleetInstancesInput{
        SpotFleetRequestId: template.FleetRequestId,
      })
      if err != nil {
        return nil, err
      }

      if len(instances.ActiveInstances) > 0 {
        s.Logger.Info("getting active instance")
        firstInstanceId := instances.ActiveInstances[0].InstanceId
        template.InstanceId = firstInstanceId
        s.Logger.Info("First instance ID:", *firstInstanceId)
      } else {
        s.Logger.Warn("No active instances found for fleet request:", *template.FleetRequestId)
      }
    }
  }

  return &pb.GetTemplateResponse{
    Error: false,
    Code: http.StatusOK,
    Message: "Success",
    Template: &pb.SpotInstanceTemplate{
      Id: uint64(template.ID),
      FleetRequestId: func() string {
          if template.FleetRequestId != nil {
              return *template.FleetRequestId
          }
          return ""
      }(),
      InstanceId: func() string {
          if template.InstanceId != nil {
              return *template.InstanceId
          }
          return ""
      }(),
      UserId: uint64(template.UserId),
      Name: template.Name,
      Status: template.Status,
      InstanceType: template.InstanceType,
      CreatedAt: timestamppb.New(template.CreatedAt),
      UpdatedAt: timestamppb.New(template.UpdatedAt),
      AmiId: template.AmiId,
			Region: template.Region,
			IpId: uint64(template.IpId),
    },
  }, nil
}

func (s *Server) StopTemplate (ctx context.Context, request *pb.StopTemplateRequest) (*pb.StopTemplateResponse, error) {
  
  template := &models.SpotInstanceTemplate{}
  if err := s.Db.Where("id = ?", request.SpotInstanceTemplateId).First(template).Error; err != nil {
    return nil, err
  }

  if template.Status != "ACTIVE" {
    return nil, errors.New("template is not active")
  }

  if template.FleetRequestId == nil {
    return nil, errors.New("fleet request ID is nil")
  }
  
	if err := s.runStopScript(ctx, *template.InstanceId); err != nil {
    s.Logger.Error("Failed to run Stop Script:", err)
		return nil, errors.New("failed to run stop script")
	}
  
  _, err := s.AWSManager.EC2Client.CancelSpotFleetRequests(ctx, &ec2.CancelSpotFleetRequestsInput{
    TerminateInstances: aws.Bool(true),
    SpotFleetRequestIds: []string{*template.FleetRequestId},
  })

  if err != nil {
    return nil, err
  }

	if err := s.Db.Model(&models.SpotInstanceTemplate{}).
    Where("id = ?", request.SpotInstanceTemplateId).
    UpdateColumns(map[string] interface{}{
      "fleet_request_id": nil,
			"instance_id": nil,
      "status": "PENDING",
    }).Error; err != nil {
    return nil, err
  }

  microIpClient := *s.IpServiceClient

  microIpClient.UnuseIp(ctx, &microIp.UnuseIpRequest{
    IpId: int64(template.IpId),
  })
  
  return &pb.StopTemplateResponse{
    Error: false,
    Code: http.StatusOK,
    Message: "Success",
  }, nil
}

// runStopScript connects to the EC2 instance and uploads the Minecraft data to S3
func (s *Server) runStopScript(ctx context.Context, instanceId string) error {
	// Fetch instance details to get public IP
	instanceDetails, err := s.AWSManager.EC2Client.DescribeInstances(ctx, &ec2.DescribeInstancesInput{
		InstanceIds: []string{instanceId},
	})
	if err != nil {
		return err
	}

	if len(instanceDetails.Reservations) == 0 || len(instanceDetails.Reservations[0].Instances) == 0 {
		return errors.New("instance not found")
	}
	instance := instanceDetails.Reservations[0].Instances[0]
	publicIP := instance.PublicIpAddress

	// SSH into the instance
	sshClient, err := s.connectSSH(*publicIP)
	if err != nil {
		return err
	}
	defer sshClient.Close()

	// Stop, Upload data files to S3
	cmd := "/home/ubuntu/mchost/mchost-config/scripts/stop.sh"
	// cmd := fmt.Sprintf("aws s3 sync /homemchost-%s", *instance.Placement.AvailabilityZone)
	session, err := sshClient.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

  s.Logger.Info("Running stop command:", cmd)
	output, err := session.CombinedOutput(cmd); 
  
  if err != nil {
		return fmt.Errorf("failed to run stop command: %v", err)
	}

	fmt.Println("Stop command Output:", string(output))
	return nil
}

// connectSSH creates an SSH connection to the EC2 instance
func (s *Server) connectSSH(host string) (*ssh.Client, error) {

	config := &ssh.ClientConfig{
		User: "ubuntu",
		Auth: []ssh.AuthMethod{
			ssh.Password("password"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	addr := fmt.Sprintf("%s:%d", host, 22)
  s.Logger.Info("Connecting to SSH:", addr)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to SSH: %w", err)
	}

  s.Logger.Info("Connected to SSH:", addr)
	return client, nil
}