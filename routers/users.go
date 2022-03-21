package routers

import (
	"gin-fast/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterUsersRouter(group *gin.RouterGroup)  {
	group.GET("/hello", handlers.HelloHandler)  // 测试接口
	group.POST("/login", handlers.UsersLoginHandler)  // 登录获取tokent 24小时过去
	group.POST("/register", handlers.UsersRegisterHandler)  // 注册用户
}

func RegisterUsersRouterWithAuth(group *gin.RouterGroup) {
	group.PUT("/userinfo", handlers.UsersSetInfoHandler)  // 更新用户信息
	group.PUT("/pwd", handlers.UsersSetPwdHandler)
	group.GET("/users", handlers.UsersListHandler) // 获取所有用户信息
}
