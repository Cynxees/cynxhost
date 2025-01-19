package porkbunmodel

type BasePorkbunRequest struct {
	SecretApiKey string `json:"secretapikey"`
	ApiKey       string `json:"apikey"`
}

type CreateDNSRequest struct {
	BasePorkbunRequest
	Name    string `json:"name"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

type UpdateDnsRequest struct {
	BasePorkbunRequest
	Content string `json:"content"`
}
