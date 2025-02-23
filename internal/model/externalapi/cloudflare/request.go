package cloudflare

type UpsertDNSRecordRequest struct {
	Type    string  `json:"type"`
	Name    string  `json:"name"`
	Content string  `json:"content"`
	TTL     int     `json:"ttl"`
	Proxied bool    `json:"proxied"`
	Comment *string `json:"comment"`
}
