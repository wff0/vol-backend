package router

import (
	"github.com/gin-gonic/gin"
	"vol-backend/api/v1/system"
	"vol-backend/middleware"
)

func Router() {
	r := gin.Default()
	r.Use(middleware.Cors())
	//r.GET("/api/cookie", system.Cookie)
	api := r.Group("/api")
	{
		api.POST("/admin/login", system.Login)

		adminGroup := api.Group("admin", middleware.Auth())
		{
			adminGroup.POST("/event/create", system.CreateEvent)
			adminGroup.POST("/event/update", system.EditEvent)
			adminGroup.POST("/event/delete", system.DeleteEvent)
			adminGroup.POST("/event/getList", system.GetEventList)

			adminGroup.POST("/vol/create", system.CreateUser)
			adminGroup.POST("/vol/update", system.EditUser)
			adminGroup.POST("/vol/delete", system.DeleteUser)
			adminGroup.POST("/vol/getList", system.GetUserInfoList)

			adminGroup.POST("/news/create", system.CreateNews)
			adminGroup.POST("/news/update", system.EditNews)
			adminGroup.POST("/news/delete", system.DeleteNews)
			adminGroup.POST("/news/getList", system.GetNewsList)
		}

	}

	r.Run(":8080")
}
