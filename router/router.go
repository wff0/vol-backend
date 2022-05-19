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

	api.GET("/getCaptcha", system.GenerateCaptcha)
	api.POST("/verifyCaptcha", system.CaptchaVerify)
	api.POST("/register", system.UserRegister)

	{
		api.POST("/admin/login", system.AdminLogin)
		api.POST("/user/login", system.UserLogin)

		adminGroup := api.Group("admin", middleware.Auth())
		{
			adminGroup.POST("/event/create", system.CreateEvent)
			adminGroup.POST("/event/update", system.EditEvent)
			adminGroup.POST("/event/delete", system.DeleteEvent)
			adminGroup.POST("/event/getList", system.GetEventList)
			adminGroup.POST("/event/finish", system.FinishEvent)

			adminGroup.POST("/vol/create", system.CreateUser)
			adminGroup.POST("/vol/update", system.AdminEditUser)
			adminGroup.POST("/vol/delete", system.DeleteUser)
			adminGroup.POST("/vol/getList", system.GetUserInfoList)

			adminGroup.POST("/news/create", system.CreateNews)
			adminGroup.POST("/news/update", system.EditNews)
			adminGroup.POST("/news/delete", system.DeleteNews)
			adminGroup.POST("/news/getList", system.GetNewsList)

			adminGroup.POST("/eventApply/getList", system.GetAllEventApplyListByEventID)
			adminGroup.POST("/eventApply/verify", system.AdminVerifyUser)
		}

		userGroup := api.Group("user", middleware.Auth())
		{
			userGroup.GET("/news/get", system.GetNewsWithID)
			userGroup.POST("/news/getAllList", system.GetNewsListToUser)

			userGroup.GET("/activity/get", system.GetActivityWithID)
			userGroup.POST("/activity/getAllList", system.GetActivityListToUser)
			userGroup.POST("/activity/apply", system.CreateEventApply)

			userGroup.GET("/eventApply/getList", system.GetAllEventApplyListByUserID)
			userGroup.POST("/eventApply/update", system.UpdateRemark)
			userGroup.POST("/eventApply/delete", system.DeleteEventApplyRecord)

			userGroup.GET("/myspace/get", system.GetUserInfoByID)
			userGroup.POST("/myspace/update", system.UserEditUserInfo)
		}
	}

	r.Run(":8080")
}
