package domain

type Shard struct {
	ShardID  string `json:"shard_id"`
	TenantID string `json:"tenant_id"`
	Address  string `json:"address"`
}
