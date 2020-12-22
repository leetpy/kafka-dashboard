package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leetpy/kafka-dashboard/router/api"
)

func Load(r *gin.Engine) *gin.Engine {
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/groups", api.ListGroup)
		apiv1.GET("/group/offset", api.ListGroupOffset)
	}
	return r
}
