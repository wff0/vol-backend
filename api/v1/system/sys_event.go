package system

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"vol-backend/entity"
	"vol-backend/model/request"
	"vol-backend/model/response"
)

func CreateEvent(c *gin.Context) {
	var event entity.Event
	err := c.ShouldBind(&event)
	if err != nil {
		log.Printf("CreateEvent err:%s", err)
		response.FailWithMessage("内部发生错误", c)
		return
	}
	cookie, _ := c.Request.Cookie("userID")
	userID, _ := strconv.Atoi(cookie.Value)
	event.UserID = userID
	err = entity.CreateEvent(event)
	if err != nil {
		log.Printf("CreateEvent err:%s", err)
		response.FailWithMessage("内部发生错误", c)
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

func DeleteEvent(c *gin.Context) {
	eventID := c.PostForm("event_id")
	id, err := strconv.Atoi(eventID)
	if err != nil {
		log.Printf("DeleteEvent err:%s", err)
		response.FailWithMessage("内部发生错误", c)
		return
	}
	err = entity.DeleteEventByID(uint(id))
	if err != nil {
		log.Printf("DeleteEvent err:%s", err)
		response.FailWithMessage("内部发生错误", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func GetEventList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.BindJSON(&pageInfo)
	if err != nil {
		log.Printf("GetEventList err:%s", err)
		response.FailWithMessage("内部发生错误", c)
		return
	}
	list, err := entity.GetEventList(pageInfo.Page, pageInfo.PageSize)
	if err != nil {
		log.Printf("GetEventList err:%s", err)
		response.FailWithMessage("内部发生错误", c)
		return
	}
	count, err := entity.GetEventCount()
	if err != nil {
		log.Printf("GetEventList err:%s", err)
		response.FailWithMessage("内部发生错误", c)
		return
	}
	response.OkWithData(map[string]interface{}{"list": list, "total": count}, c)
}
