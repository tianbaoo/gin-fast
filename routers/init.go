package routers

import (
	_ "gin-fast/docs"
	"gin-fast/logger"
	"gin-fast/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	// "go.uber.org/zap"
)

type UrlGroup func(group *gin.RouterGroup)

func InitRouter(isErrMsg, isDebugMode bool) *gin.Engine {
	router := gin.New()
	// 测试日志功能
	// var (
	// 	name = "tianbao"
	// 	age  = 26
	// )
	// 记录日志并使用zap.Xxx(key, val)记录相关字段
	// zap.L().Debug("this is hello func", zap.String("user", name), zap.Int("age", age))
	// zap.L().Info("hello world", zap.String("user", name), zap.Int("age", age))
	// zap.L().Error("this is hello func", zap.String("user", name), zap.Int("age", age))
	router.Use(logger.GinLogger(), logger.GinRecovery(true)) // 重写gin原来的继承的2个中间件
	//router.Use(gin.Logger())
	//router.Use(gin.Recovery())
	router.Use(middlewares.ErrorHandleMiddleware(isErrMsg))

	//// 设置静态资源路由路径
	//if isDebugMode {
	//	var temp map[string]string
	//	if err := json.Unmarshal([]byte(conf.ProjectCfg.StaticUrlMapPath), &temp); err == nil {
	//		for url, path := range temp {
	//			router.Static(url, path)
	//		}
	//	}
	//}
	//
	//// 加载HTML模版
	//router.LoadHTMLGlob(conf.ProjectCfg.TemplateGlob)

	//
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 无需token验证路由组
	routerGroupWithNoAuth := router.Group("api")
	RegisterUsersRouter(routerGroupWithNoAuth)  // 用户注册相关

	// 需要token验证，使用middlewares.AuthJwtTokenMiddleware中间件验证
	routerGroupWithAuth := router.Group("api/v1")
	routerGroupWithAuth.Use(middlewares.AuthJwtTokenMiddleware())  // 进行鉴权
	RegisterUsersRouterWithAuth(routerGroupWithAuth)
	RegisterCmsRouter(routerGroupWithAuth)

	return router
}
