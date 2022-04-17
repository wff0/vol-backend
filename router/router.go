package router

import (
	"github.com/gin-gonic/gin"
	"vol-backend/api/v1/system"
)

func Router() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.POST("/api/createEvent", system.CreateEvent)
	}

	r.Run(":8080")
}
