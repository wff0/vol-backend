package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"gorm.io/gorm"
	"log"
	"strconv"
	"vol-backend/entity"
	"vol-backend/model/request"
	"vol-backend/model/response"
)

func AdminLogin(c *gin.Context) {
	captchaID := c.PostForm("captchaID")
	verifyCode := c.PostForm("verifyCode")

	if !store.Verify(captchaID, verifyCode, true) {
		response.FailWithMessage("验证码错误", c)
		return
	}

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
	if newUser.Password != user.Password {
		response.FailWithMessage("密码错误！", c)
		return
	}
	if err != nil {
		response.FailWithMessage("内部发生错误", c)
		log.Printf("FindUserByUsernameAndPassword err:%s", err)
		return
	}
	if newUser.Role == 1 {
		response.OkWithData(newUser.ID, c)
	} else if newUser.Role == 2 {
		response.FailWithMessage("当前用户不是管理员", c)
	} else {
		log.Printf("FindUserByUsernameAndPassword err:%s", err)
		response.FailWithMessage("内部发生错误", c)
	}
}

func UserLogin(c *gin.Context) {
	captchaID := c.PostForm("captchaID")
	verifyCode := c.PostForm("verifyCode")

	if !store.Verify(captchaID, verifyCode, true) {
		response.FailWithMessage("验证码错误", c)
		return
	}

	user := entity.UserInfo{}
	username := c.PostForm("username")
	password := c.PostForm("password")
	user.Username = username
	user.Password = password
	newUser, err := entity.FindUserByUsernameAndPassword(user)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMessage("用户名不存在，请先注册！", c)
		log.Printf("FindUserByUsernameAndPassword err:%s", err)
		return
	}
	if newUser.Password != user.Password {
		response.FailWithMessage("密码错误！", c)
		return
	}
	if err != nil {
		response.FailWithMessage("内部发生错误", c)
		log.Printf("FindUserByUsernameAndPassword err:%s", err)
		return
	}
	if newUser.Role != 0 {
		response.OkWithData(newUser.ID, c)
	} else {
		log.Printf("FindUserByUsernameAndPassword err:%s", err)
		response.FailWithMessage("内部发生错误", c)
	}
}

func UserRegister(c *gin.Context) {
	captchaID := c.PostForm("captchaID")
	verifyCode := c.PostForm("verifyCode")

	if !store.Verify(captchaID, verifyCode, true) {
		response.FailWithMessage("验证码错误", c)
		return
	}

	username := c.PostForm("username")
	password := c.PostForm("password")
	gender := c.PostForm("gender")
	school := c.PostForm("school")
	classroom := c.PostForm("classroom")

	var userInfo entity.UserInfo
	userInfo.Username = username
	userInfo.Password = password
	userInfo.Gender = gender
	userInfo.School = school
	userInfo.Classroom = classroom

	err := entity.CreateUserInfo(userInfo)
	if err != nil {
		log.Printf("UserRegister err:%s", err)
		response.FailWithMessage("内部发生错误", c)
		return
	}
	response.OkWithMessage("注册成功", c)
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

func AdminEditUser(c *gin.Context) {
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

func GetUserInfoByID(c *gin.Context) {
	cookie, _ := c.Request.Cookie("userID")
	userID, _ := strconv.Atoi(cookie.Value)
	record, err := entity.GetUserInfoByID(uint(userID))
	if err != nil {
		log.Printf("GetUserInfoByID err:%s", err)
		response.FailWithMessage("内部发生错误", c)
		return
	}
	response.OkWithData(map[string]interface{}{"userInfo": record}, c)
}

func UserEditUserInfo(c *gin.Context) {
	var user entity.UserInfo
	err := c.Bind(&user)
	if err != nil {
		response.FailWithMessage("内部发生错误", c)
		log.Printf("CreateUser err:%s", err)
		return
	}
	err = entity.UserUpdateUserInfoByID(user, user.ID)
	if err != nil {
		response.FailWithMessage("内部发生错误", c)
		log.Printf("CreateUser err:%s", err)
		return
	}
	response.OkWithMessage("编辑成功", c)
}

var (
	store  = base64Captcha.DefaultMemStore
	driver = base64Captcha.DefaultDriverDigit
)

func GenerateCaptcha(c *gin.Context) {
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := captcha.Generate()
	if err != nil {
		response.FailWithMessage("内部发生错误", c)
		log.Printf("GenerateCaptcha err:%s", err)
		return
	}
	response.OkWithData(map[string]interface{}{"id": id, "img": b64s}, c)
}

func CaptchaVerify(c *gin.Context) {
	id := c.PostForm("id")
	captcha := c.PostForm("captcha")

	if store.Verify(id, captcha, true) {
		response.OkWithMessage("验证码成功", c)
	} else {
		response.FailWithMessage("验证码失败", c)
	}
}
