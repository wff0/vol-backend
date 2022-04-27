package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"strconv"
	"vol-backend/entity"
	"vol-backend/model/request"
	"vol-backend/model/response"
)

func Login(c *gin.Context) {
	user := entity.UserInfo{}
	username := c.PostForm("username")
	password := c.PostForm("password")
	user.Username = username
	user.Password = password
	newUser, err := entity.FindUserByUsernameAndPassword(user)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMessage("用户不存在", c)
		log.Printf("FindUserByUsernameAndPassword err:%s", err)
		return
	}
	if err != nil {
		response.FailWithMessage("内部发生错误", c)
		log.Printf("FindUserByUsernameAndPassword err:%s", err)
		return
	}
	if newUser.Role == 1 {
		//c.SetCookie("userID",
		//	strconv.Itoa(int(newUser.ID)),
		//	3600,
		//	"/",
		//	"localhost",
		//	false,
		//	true)
		response.OkWithData(newUser.ID, c)
	} else if newUser.Role == 2 {
		response.FailWithMessage("当前用户不是管理员", c)
	} else {
		log.Printf("FindUserByUsernameAndPassword err:%s", err)
		response.FailWithMessage("内部发生错误", c)
	}
}

func Cookie(c *gin.Context) {
	c.SetCookie("userID",
		"1",
		3600,
		"/",
		"localhost",
		false,
		true)
	response.OkWithMessage("登录成功", c)
}

func CreateUser(c *gin.Context) {
	var user entity.UserInfo
	err := c.Bind(&user)
	if err != nil {
		response.FailWithMessage("内部发生错误", c)
		log.Printf("CreateUser err:%s", err)
		return
	}
	err = entity.CreateUserInfo(user)
	if err != nil {
		response.FailWithMessage("内部发生错误", c)
		log.Printf("CreateUser err:%s", err)
		return
	}
	response.OkWithMessage("添加成功", c)
}

func EditUser(c *gin.Context) {
	var user entity.UserInfo
	err := c.Bind(&user)
	if err != nil {
		response.FailWithMessage("内部发生错误", c)
		log.Printf("CreateUser err:%s", err)
		return
	}
	err = entity.UpdateUserInfoByID(user, user.ID)
	if err != nil {
		response.FailWithMessage("内部发生错误", c)
		log.Printf("CreateUser err:%s", err)
		return
	}
	response.OkWithMessage("编辑成功", c)
}

func DeleteUser(c *gin.Context) {
	userID := c.PostForm("user_id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		log.Printf("DeleteUser err:%s", err)
		response.FailWithMessage("内部发生错误", c)
		return
	}
	err = entity.DeleteUserByID(uint(id))
	if err != nil {
		log.Printf("DeleteUser err:%s", err)
		response.FailWithMessage("内部发生错误", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func GetUserInfoList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.BindJSON(&pageInfo)
	if err != nil {
		log.Printf("GetUserInfoList err:%s", err)
		response.FailWithMessage("内部发生错误", c)
		return
	}
	list, err := entity.GetUserInfoList(pageInfo.Page, pageInfo.PageSize)
	if err != nil {
		log.Printf("GetUserInfoList err:%s", err)
		response.FailWithMessage("内部发生错误", c)
		return
	}
	count, err := entity.GetUserCount()
	if err != nil {
		log.Printf("GetUserInfoList err:%s", err)
		response.FailWithMessage("内部发生错误", c)
		return
	}
	response.OkWithData(map[string]interface{}{"list": list, "total": count}, c)
}
