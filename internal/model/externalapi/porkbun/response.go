package porkbunmodel

type CreateDNSResponse struct {
	Status string `json:"status"`
	Id     string `json:"id"`
}

type UpdateDNSResponse struct {
	Status string `json:"status"`
}

type RetrieveDNSByTypeSubdomain struct {
	Status  string   `json:"status"`
	Records []Record `json:"records"`
}

type Record struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Content string `json:"content"`
	TTL     string `json:"ttl"`
	Prio    string `json:"prio"`
	Notes   string `json:"notes"`
}
