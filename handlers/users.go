package handlers

import (
	"gin-fast/models"
	"gin-fast/pkg/jwt"
	"gin-fast/pkg/util"
	"gin-fast/serializers"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UsersLoginHandler 登录接口
// @Summary 用户登录接口
// @Description 用户登录接口
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param data body serializers.Login true "请求body中需携带的参数"
// @Success 200 {string} string "{"code":200,"msg":"返回成功","data":{"email":"gtb365@163.com","id":93997900482051, "name":"","phone":"15926187015","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDAwNjczMzksImlkIjo5Mzk5NzkwMDQ4MjA1MSwidXNlcm5hbWUiOiJ0aWFuYmFvIn0.luFD6rMu3eYrnvWaIu8uBUiKszFGnkFrf8b9-0djnPc","username":"tianbao"}}"
// @Failure 400 {string} string "{"code": 400, "msg": "返回错误"}"
// @Router /api/login [post]

func HelloHandler(ctx *gin.Context) {
	response := Response{Ctx: ctx}
	response.Response(nil, nil)
	return
}

func UsersLoginHandler(ctx *gin.Context) {
	response := Response{Ctx: ctx}
	var loginUser serializers.Login
	if err := ctx.ShouldBindJSON(&loginUser); err != nil {
		zap.L().Error("UsersLoginHandler func", zap.String("user", "admin"), zap.Int("msg", 0))
		panic(err)
	}
	user := loginUser.GetUser()
	isLoginUser := user.CheckPassword()
	if !isLoginUser {
		response.BadRequest("密码错误")
		return
	}
	token, err := jwt.GenToken(user.ID, user.Username)
	if err != nil {
		panic(err)
	}
	data, _ := util.PrecisionLost(user)
	data["token"] = token
	response.Response(data, nil)
	return
}

// UsersRegisterHandler 注册接口
// @Summary 用户注册接口
// @Description 用户注册接口
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param data body serializers.Register true "请求body中需携带的参数"
// @Success 200 {string} json "{"code": 200, "msg": "返回成功"}"
// @Failure 400 {string} string "{"code": 400, "msg": "返回错误"}"
// @Router /api/register [post]
func UsersRegisterHandler(ctx *gin.Context) {

	response := Response{Ctx: ctx}
	var registerUser serializers.Register
	if err := ctx.ShouldBind(&registerUser); err != nil {
		zap.L().Error("注册用户时，前端传参数有问题", zap.Error(err))
		panic(err)
	}
	zap.L().Error("参数", zap.String("username",registerUser.Username),
		zap.String("password",registerUser.Password),
		zap.String("phone",registerUser.Phone),
		zap.String("email",registerUser.Email))
	// 创建user对象，并将前端传来的参数序列化到user对象结构体中
	user := registerUser.GetUser()
	// 校验用户名是否已存在
	status := user.CheckDuplicateUsername()
	if status == false {
		response.BadRequest("用户名已存在")
		return
	}
	// 将密码加密进行存储
	if err := user.SetPassword(user.Password); err != nil {
		panic(err)
	}
	// 用户激活状态改为true
	user.IsActive = true
	// 数据库创建用户
	models.DB.Create(&user)
	zap.L().Info("创建用户: ", zap.String("user ", user.Username))
	response.Response(nil, nil)
}

// UsersSetInfoHandler 修改用户信息
func UsersSetInfoHandler(ctx *gin.Context) {
	response := Response{Ctx: ctx}
	// 封装GetBodyData，从request的body中获取参数
	jsonData, err := util.GetBodyData(ctx)
	if err != nil {
		response.BadRequest("参数解析失败")
		return
	}
	_, ok := jsonData["password"]
	if ok {
		response.BadRequest("不允许通过此接口修改密码")
		return
	}
	if jsonData == nil {
		response.BadRequest("参数为空")
		return
	}
	currentUser := jwt.AssertUser(ctx)
	if currentUser != nil {
		models.DB.Model(&currentUser).Updates(jsonData)
		response.Response(currentUser, nil)
		return
	}
}

// UsersSetPwdHandler 修改密码
func UsersSetPwdHandler(ctx *gin.Context) {
	response := Response{Ctx: ctx}
	currentUser := jwt.AssertUser(ctx)
	if currentUser == nil {
		response.Unauthenticated("未验证登录")
		return
	}
	var user serializers.Account
	if err := ctx.ShouldBindJSON(&user); err != nil {
		response.BadRequest(err.Error())
		return
	}
	if user.Username != currentUser.Username {
		response.BadRequest("当前登录用户用户名与输入用户名不符，用户名不可修改")
		return
	}
	if user.OldPwd == user.NewPwd {
		response.BadRequest("本次密码与上一次密码相同")
		return
	}
	if isPwd := currentUser.IsPasswordEqual(user.OldPwd); !isPwd {
		response.BadRequest("原密码错误，请重新确认")
		return
	}
	if err := currentUser.SetPassword(user.NewPwd); err != nil {
		response.BadRequest(err.Error())
		return
	}
	models.DB.Save(&currentUser)
	response.Response(nil, nil)
}

// UsersListHandler 获取用户列表
func UsersListHandler(ctx *gin.Context) {
	// http://127.0.0.1:7890/api/v1/users?page=3 通过page参数获取分页
	response := Response{Ctx: ctx}
	var pager serializers.Pager
	pager.InitPager(ctx)
	var users []models.Account
	db := models.DB.Model(&users)
	db.Count(&pager.Total)
	db.Offset(pager.OffSet).Limit(pager.PageSize).Find(&users)
	pager.GetPager()
	response.Response(users, pager)
}
