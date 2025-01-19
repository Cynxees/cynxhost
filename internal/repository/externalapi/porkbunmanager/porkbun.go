package porkbunmanager

import (
	"cynxhost/internal/dependencies"
	"cynxhost/internal/helper"
	porkbunmodel "cynxhost/internal/model/externalapi/porkbun"
	"encoding/json"
	"fmt"
)

type PorkbunManager struct {
	Config *dependencies.ConfigPorkbun
}

func New(config *dependencies.ConfigPorkbun) *PorkbunManager {
	return &PorkbunManager{
		Config: config,
	}
}

// CreateDNS will create a new DNS record for the subdomain.
func (p *PorkbunManager) CreateDNS(subdomain string, ip string) (*porkbunmodel.CreateDNSResponse, error) {
	url := "https://api.porkbun.com/api/json/v3/dns/create/" + p.Config.Domain

	// Create the request body
	request := porkbunmodel.CreateDNSRequest{
		BasePorkbunRequest: porkbunmodel.BasePorkbunRequest{
			SecretApiKey: p.Config.SecretKey,
			ApiKey:       p.Config.ApiKey,
		},
		Name:    subdomain,
		Type:    "A",
		Content: ip,
	}

	// Send the POST request
	resp, err := helper.SendPostRequest(url, request, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create DNS record: %v", err)
	}

	// Parse the response to check if the DNS creation was successful
	var response porkbunmodel.CreateDNSResponse
	if err := json.Unmarshal([]byte(resp), &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	// Check if there is an error in the response from the Porkbun API
	if response.Status != "SUCCESS" {
		return &response, fmt.Errorf("failed to create DNS record, status: %s, id: %d", response.Status, response.Id)
	}

	return &response, nil
}

func (p *PorkbunManager) UpdateDNS(recordType string, subdomain string, newContent string) error {
	url := fmt.Sprintf("https://api.porkbun.com/api/json/v3/dns/editByNameType/%s/%s/%s", p.Config.Domain, recordType, subdomain)

	request := porkbunmodel.UpdateDnsRequest{
		BasePorkbunRequest: porkbunmodel.BasePorkbunRequest{
			ApiKey:       p.Config.ApiKey,
			SecretApiKey: p.Config.SecretKey,
		},
		Content: newContent,
	}

	// Send the POST request
	resp, err := helper.SendPostRequest(url, request, nil)
	if err != nil {
		return fmt.Errorf("failed to update DNS record: %v", err)
	}

	// Parse the response to check if the DNS update was successful
	var response porkbunmodel.UpdateDNSResponse
	if err := json.Unmarshal([]byte(resp), &response); err != nil {
		return fmt.Errorf("failed to parse response: %v", err)
	}

	// Check if there is an error in the response from the Porkbun API
	if response.Status != "SUCCESS" {
		return fmt.Errorf("failed to update DNS record, status: %s", response.Status)
	}

	return nil
}

/**
 * RetreiveDNSByTypeSubdomain will check if a DNS record exists for the given subdomain and record type.
 * @param recordType The type of DNS record (e.g., A, CNAME, etc.)
 * @param subdomain The subdomain for which to check the DNS record
 */
func (p *PorkbunManager) RetrieveDNSByTypeSubdomain(recordType string, subdomain string) (*porkbunmodel.RetrieveDNSByTypeSubdomain, error) {
	url := fmt.Sprintf("https://api.porkbun.com/api/json/v3/dns/retrieveByNameType/%s/%s/%s", p.Config.Domain, recordType, subdomain)

	request := porkbunmodel.BasePorkbunRequest{
		ApiKey:       p.Config.ApiKey,
		SecretApiKey: p.Config.SecretKey,
	}

	// Make a GET request to check if the DNS record exists
	resp, err := helper.SendPostRequest(url, request, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to check DNS record: %v", err)
	}

	// Parse the response to determine if the DNS record exists
	var response porkbunmodel.RetrieveDNSByTypeSubdomain
	if err := json.Unmarshal([]byte(resp), &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	if response.Status != "SUCCESS" {
		return nil, fmt.Errorf("DNS record does not exist for subdomain: %s", subdomain)
	}

	return &response, nil
}
