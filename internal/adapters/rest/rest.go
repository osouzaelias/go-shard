package rest

import (
	"github.com/gin-gonic/gin"
)

func (a Adapter) Get(c *gin.Context) {
	shard, err := a.api.GetShard(c.Request.Context(), c.Query("tenant"), c.Query("customerId"))
	if err != nil {
		c.JSON(422, gin.H{"Error": err})
	}

	c.JSON(200, gin.H{"ShardId": shard.ID})
}

func (a Adapter) HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"Status": "UP"})
}
