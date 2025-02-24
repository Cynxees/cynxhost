package cloudflaremanager

import (
	"cynxhost/internal/dependencies"
	"cynxhost/internal/helper"
	cloudflaremodel "cynxhost/internal/model/externalapi/cloudflare"
	"encoding/json"
	"fmt"
)

type CloudflareManager struct {
	Config *dependencies.ConfigCloudflare
}

func New(config *dependencies.ConfigCloudflare) *CloudflareManager {
	return &CloudflareManager{
		Config: config,
	}
}

// CreateDNS will create a new DNS record for the subdomain.
func (c *CloudflareManager) CreateDNSRecord(recordType string, subdomain string, ip string) (cloudflaremodel.UpsertDNSRecordResponse, error) {
	url := "https://api.cloudflare.com/client/v4/zones/" + c.Config.ZoneId + "/dns_records"

	proxied := false

	// Create the request body
	request := cloudflaremodel.UpsertDNSRecordRequest{
		Type:    recordType,
		Name:    subdomain,
		Content: ip,
		TTL:     3600,
		Proxied: proxied,
	}

	// Send the POST request
	resp, err := helper.SendHttpRequest("POST", url, request, map[string]string{
		"Authorization": "Bearer " + c.Config.ApiToken,
		"X-Auth-Email":  c.Config.Email,
		"Content-Type":  "application/json",
	})
	if err != nil {
		return cloudflaremodel.UpsertDNSRecordResponse{}, fmt.Errorf("failed to create DNS record: %v", err)
	}

	// Parse the response to check if the DNS creation was successful
	var response cloudflaremodel.UpsertDNSRecordResponse
	if err := json.Unmarshal([]byte(resp), &response); err != nil {
		return cloudflaremodel.UpsertDNSRecordResponse{}, fmt.Errorf("failed to parse response: %v", err)
	}

	// Check if there is an error in the response from the Cloudflare API
	if !response.Success {
		return cloudflaremodel.UpsertDNSRecordResponse{}, fmt.Errorf("failed to create DNS record, errors: %v", response.Errors)
	}

	return response, nil
}

func (c *CloudflareManager) UpdateDNS(recordId string, recordType string, subdomain string, content string) error {
	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s", c.Config.ZoneId, recordId)

	request := cloudflaremodel.UpsertDNSRecordRequest{
		Content: content,
		Type:    recordType,
		Name:    subdomain,
		Comment: &subdomain,
		Proxied: true,
	}

	// Send the PUT request
	resp, err := helper.SendHttpRequest("PATCH", url, request, map[string]string{
		"Authorization": "Bearer " + c.Config.ApiToken,
		"X-Auth-Email":  c.Config.Email,
		"Content-Type":  "application/json",
	})
	if err != nil {
		return fmt.Errorf("failed to update DNS record: %v", err)
	}

	// Parse the response to check if the DNS update was successful
	var response cloudflaremodel.UpsertDNSRecordResponse
	if err := json.Unmarshal([]byte(resp), &response); err != nil {
		return fmt.Errorf("failed to parse response: %v", err)
	}

	// Check if there is an error in the response from the Cloudflare API
	if !response.Success {
		return fmt.Errorf("failed to update DNS record, errors: %v", response.Errors)
	}

	return nil
}
