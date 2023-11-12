package domain

type Shard struct {
	ID     string `json:"id"`
	Tenant string `json:"tenant"`
	Total  uint8  `json:"total"`
}
