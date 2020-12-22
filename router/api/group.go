package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leetpy/kafka-dashboard/lib/kafka"
)

// ListGroup 返回所有的 group
func ListGroup(c *gin.Context) {
	groups := kafka.Groups()
	c.JSON(http.StatusOK, gin.H{
		"error_code":    0,
		"error_message": "",
		"data":          groups,
	})
}

// ListGroupOffset group offset
func ListGroupOffset(c *gin.Context) {
	group := c.Query("group")
	offsets := kafka.GroupOffsets(group)
	c.JSON(http.StatusOK, gin.H{
		"error_code":    0,
		"error_message": "",
		"data":          offsets,
	})
}
