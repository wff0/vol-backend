package system

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"vol-backend/entity"
	"vol-backend/model/request"
	"vol-backend/model/response"
)

func CreateNews(c *gin.Context) {
	var news entity.News
	err := c.ShouldBind(&news)
	if err != nil {
		response.FailWithMessage("内部发生错误", c)
		log.Printf("CreateNews err:%s", err)
		return
	}
	cookie, _ := c.Request.Cookie("userID")
	userID, _ := strconv.Atoi(cookie.Value)
	news.UserID = userID
	err = entity.CreateNews(news)
	if err != nil {
		response.FailWithMessage("内部发生错误", c)
		log.Printf("CreateNews err:%s", err)
		return
	}
	response.OkWithMessage("添加成功", c)
}

func EditNews(c *gin.Context) {
	var news entity.News
	err := c.Bind(&news)
	if err != nil {
		response.FailWithMessage("内部发生错误", c)
		log.Printf("EditNews err:%s", err)
		return
	}
	err = entity.UpdateNewsByID(news, news.ID)
	if err != nil {
		response.FailWithMessage("内部发生错误", c)
		log.Printf("EditNews err:%s", err)
		return
	}
	response.OkWithMessage("编辑成功", c)
}

func DeleteNews(c *gin.Context) {
	newsID := c.PostForm("news_id")
	id, err := strconv.Atoi(newsID)
	if err != nil {
		log.Printf("DeleteNews err:%s", err)
		response.FailWithMessage("内部发生错误", c)
		return
	}
	err = entity.DeleteNewsByID(uint(id))
	if err != nil {
		log.Printf("DeleteNews err:%s", err)
		response.FailWithMessage("内部发生错误", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func GetNewsList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.BindJSON(&pageInfo)
	if err != nil {
		log.Printf("GetNewsList err:%s", err)
		response.FailWithMessage("内部发生错误", c)
		return
	}
	list, err := entity.GetNewsList(pageInfo.Page, pageInfo.PageSize)
	if err != nil {
		log.Printf("GetNewsList err:%s", err)
		response.FailWithMessage("内部发生错误", c)
		return
	}
	count, err := entity.GetNewsCount()
	if err != nil {
		log.Printf("GetNewsList err:%s", err)
		response.FailWithMessage("内部发生错误", c)
		return
	}
	response.OkWithData(map[string]interface{}{"list": list, "total": count}, c)
}
