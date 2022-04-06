package system

import (
	"github.com/gin-gonic/gin"
	"log"
	"vol-backend/entity"
	"vol-backend/model/response"
)

func CreateEvent(c *gin.Context) {
	var event entity.Event
	err := c.Bind(&event)
	if err != nil {
		response.FailWithMessage("内部发生错误", c)
		return
	}
	err = entity.CreateEvent(event)
	if err != nil {
		response.FailWithMessage("内部发生错误", c)
		log.Printf("CreateEvent err:%s", err)
		return
	}
	response.OkWithMessage("添加成功", c)
}

func EditEvent(c *gin.Context) {
	var event entity.Event
	err := c.Bind(&event)
	if err != nil {
		response.FailWithMessage("内部发生错误", c)
		return
	}
	err = entity.UpdateEventAll(event, event.ID)
	if err != nil {
		response.FailWithMessage("内部发生错误", c)
		log.Printf("CreateEvent err:%s", err)
		return
	}
	response.OkWithMessage("编辑成功", c)
}
