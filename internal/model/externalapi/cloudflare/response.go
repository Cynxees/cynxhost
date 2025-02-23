package cloudflare

type UpsertDNSRecordResponse struct {
	Result   DNSRecordResult `json:"result"`
	Success  bool            `json:"success"`
	Errors   []interface{}   `json:"errors"`
	Messages []interface{}   `json:"messages"`
}

type DNSRecordResult struct {
	ID                string                 `json:"id"`
	ZoneID            string                 `json:"zone_id"`
	ZoneName          string                 `json:"zone_name"`
	Name              string                 `json:"name"`
	Type              string                 `json:"type"`
	Content           string                 `json:"content"`
	Proxiable         bool                   `json:"proxiable"`
	Proxied           bool                   `json:"proxied"`
	TTL               int                    `json:"ttl"`
	Settings          map[string]interface{} `json:"settings"`
	Meta              map[string]interface{} `json:"meta"`
	Comment           string                 `json:"comment"`
	Tags              []string               `json:"tags"`
	CreatedOn         string                 `json:"created_on"`
	ModifiedOn        string                 `json:"modified_on"`
	CommentModifiedOn string                 `json:"comment_modified_on"`
}
