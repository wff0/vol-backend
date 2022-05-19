package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"strconv"
	"vol-backend/entity"
	"vol-backend/model/response"
)

func CreateEventApply(c *gin.Context) {
	query := c.PostForm("activityID")
	remark := c.PostForm("remark")
	cookie, _ := c.Request.Cookie("userID")
	userID, _ := strconv.Atoi(cookie.Value)
	activityID, _ := strconv.Atoi(query)

	event, err := entity.SelectEventById(uint(activityID))
	if err != nil {
		log.Printf("CreateEventApply err:%s", err)
		response.FailWithMessage("内部发生错误", c)
		return
	}
	if event.Status != entity.EventSignUp {
		response.FailWithMessage("活动已停止报名", c)
		return
	}

	// 检查是否重复报名
	record, err := entity.SearchUserIfApplyEvent(uint(activityID), uint(userID))
	// 未报名过相同活动
	if errors.Is(gorm.ErrRecordNotFound, err) {
		var eventApply entity.EventApply
		eventApply.UserID = userID
		eventApply.EventID = activityID
		eventApply.Remark = remark
		eventApply.ApplyStatus = entity.EventApplyProcessing
		err = entity.CreateEventApply(eventApply)
		if err != nil {
			log.Printf("CreateEventApply err:%s", err)
			response.FailWithMessage("内部发生错误", c)
			return
		}
		response.OkWithMessage("报名成功", c)
		return
	} else if record.ID != 0 {
		// 报名过相同活动
		response.OkWithMessage("不可重复报名", c)
		return
	}
	log.Printf("CreateEventApply err:%s", err)
	response.OkWithMessage("发生未知错误", c)
}

func GetAllEventApplyListByEventID(c *gin.Context) {
	formID := c.PostForm("eventID")
	eventID, _ := strconv.Atoi(formID)
	list, err := entity.GetAllEventApplyListByEventID(uint(eventID))
	if err != nil {
		log.Printf("GetAllEventApplyListByEventID err:%s", err)
		response.OkWithMessage("内部发生错误", c)
		return
	}
	response.OkWithData(map[string]interface{}{"list": list}, c)
}

func GetAllEventApplyListByUserID(c *gin.Context) {
	cookie, _ := c.Request.Cookie("userID")
	userID, _ := strconv.Atoi(cookie.Value)
	list, err := entity.GetAllEventApplyListByUserID(uint(userID))
	if err != nil {
		log.Printf("GetAllEventApplyListByUserID err:%s", err)
		response.FailWithMessage("内部发生错误", c)
		return
	}
	response.OkWithData(map[string]interface{}{"list": list}, c)
}

func UpdateRemark(c *gin.Context) {
	query := c.PostForm("event_apply_id")
	remark := c.PostForm("remark")
	id, _ := strconv.Atoi(query)

	record, err := entity.SelectEventApplyRecordByID(uint(id))
	if record.ApplyStatus != entity.EventApplyProcessing {
		response.FailWithMessage("管理员已经审评不可再操作", c)
		return
	}

	err = entity.UpdateEventApplyRemark(remark, uint(id))
	if err != nil {
		log.Printf("GetAllEventApplyListByUserID err:%s", err)
		response.OkWithMessage("内部发生错误", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

func DeleteEventApplyRecord(c *gin.Context) {
	query := c.PostForm("event_apply_id")
	id, _ := strconv.Atoi(query)

	record, err := entity.SelectEventApplyRecordByID(uint(id))
	if record.ApplyStatus != entity.EventApplyProcessing {
		response.FailWithMessage("管理员已经审评不可再操作", c)
		return
	}

	err = entity.DeleteEventApplyByID(uint(id))
	if err != nil {
		log.Printf("DeleteEventApplyRecord err:%s", err)
		response.OkWithMessage("内部发生错误", c)
		return
	}
	response.OkWithMessage("取消成功", c)
}

func AdminVerifyUser(c *gin.Context) {
	status := c.PostForm("verify_status")
	tmpID := c.PostForm("event_apply_id")
	id, _ := strconv.Atoi(tmpID)

	record, err := entity.SelectEventApplyRecordByID(uint(id))
	if err != nil {
		log.Printf("AdminVerifyUser err:%s", err)
		response.FailWithMessage("内部发生错误", c)
		return
	}
	if record.ApplyStatus != entity.EventApplyProcessing {
		response.FailWithMessage("已对该用户操作过", c)
		return
	}

	err = entity.UpdateEventApplyStatus(status, uint(id))
	if err != nil {
		log.Printf("AdminVerifyUser err:%s", err)
		response.OkWithMessage("内部发生错误", c)
		return
	}
	response.OkWithMessage("操作成功", c)
}
