package system

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"log"
	"strconv"
	"vol-backend/entity"
	"vol-backend/global"
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
	conn := global.VB_REDIS_POOL.Get()
	_, err = conn.Do("DEL", RedisActivityListKey)
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

func GetActivityWithID(c *gin.Context) {
	ids := c.Query("activityID")
	id, _ := strconv.Atoi(ids)
	activity, err := entity.SelectEventById(uint(id))
	if err != nil {
		log.Printf("GetActivityWithID err:%s", err)
		response.FailWithMessage("内部发生错误", c)
		return
	}
	response.OkWithData(map[string]interface{}{"activity": activity}, c)
}

const RedisActivityListKey = "redis_activity_list"
const RedisActivityListTTL = 2 * 60

func GetActivityListToUser(c *gin.Context) {
	conn := global.VB_REDIS_POOL.Get()
	defer conn.Close()
	redisList, redisError := redis.String(conn.Do("GET", RedisActivityListKey))
	if errors.Is(redisError, redis.ErrNil) {
		list, err := entity.GetAllActivityList()
		if err != nil {
			log.Printf("GetActivityListToUser err:%s", err)
			response.FailWithMessage("内部发生错误", c)
			return
		}

		jsonList, err := json.Marshal(list)
		if err != nil {
			log.Printf("GetActivityListToUser err:%s", err)
			response.FailWithMessage("内部发生错误", c)
			return
		}

		_, err = conn.Do("SETEX", RedisActivityListKey, RedisActivityListTTL, jsonList)
		if err != nil {
			log.Printf("GetActivityListToUser err:%s", err)
			response.FailWithMessage("内部发生错误", c)
			return
		}
		response.OkWithData(map[string]interface{}{"list": list}, c)
	} else {
		response.OkWithData(map[string]interface{}{"list": json.RawMessage(redisList)}, c)
	}
}

func FinishEvent(c *gin.Context) {
	ids := c.PostForm("activityID")
	id, _ := strconv.Atoi(ids)
	err := entity.UpdateEventStatus(uint(id), entity.EventEnd)
	if err != nil {
		log.Printf("FinishEvent err:%s", err)
		response.FailWithMessage("内部发生错误", c)
		return
	}
	response.OkWithMessage("操作成功", c)
}
