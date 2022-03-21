package main

import (
	"flag"
	"fmt"
	"gin-fast/conf"
	"gin-fast/logger"
	"gin-fast/models"
	"gin-fast/routers"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"os"
)

var host string
var port string
var isDebugMode bool
var isErrMsg bool
var isOrmDebug bool


func init()  {
	flag.StringVar(&host, "h", "0.0.0.0", "主机")
	flag.StringVar(&port, "p", "", "监听端口")
	flag.BoolVar(&isDebugMode, "debug", true, "是否开启debug")
	flag.BoolVar(&isErrMsg, "err", true, "是否返回错误信息")
	flag.BoolVar(&isOrmDebug, "orm", true, "是否开启gorm的debug信息")
	flag.Parse()

	// 配置项初始化
	conf.SetUp()

	// 初始化日志logger
	if err := logger.InitLogger(conf.LoggerCfg); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
	}

	// mysql初始化连接
	//models.SetUp(isOrmDebug)

	// redis初始化连接
	//_ = models.InitRedisClient()
	//models.RedisExample()
}

// @title gin-fast API
// @version 1.0
// @description gin快速启动脚手架
// @termsOfService https://tianbaoo.github.io/
// @contact.name Tianbao
// @contact.url https://tianbaoo.github.io/
// @contact.email gtb365@163.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:7890
// @BasePath /api
func main() {

	// 从命令行获取端口参数，获取不到时，默认为空值, 然后从配置文件的log的Mode参数读取字段
	if isDebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(conf.LoggerCfg.Mode)  // debug release test
		//gin.SetMode(gin.ReleaseMode)
	}

	// 从命令行获取端口参数，获取不到时，从配置文件读取
	if len([]rune(port)) < 4 {
		port = conf.HttpServer.Port
	}

	// 将终端显示的数据写入日志文件
	var f *os.File
	if err := os.Mkdir("logs", os.ModePerm); err != nil {
		f, _ = os.OpenFile("logs/gin.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	} else {
		f, _ = os.Create("logs/gin.log")
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 程序退出时断开数据库连接
	defer models.DB.Close()
	defer f.Close()

	// 路由初始化
	router := routers.InitRouter(isErrMsg, isDebugMode)

	// 测试log输出
	zap.L().Error("this is hello func", zap.String("user", "111"), zap.Int("age", 123))

	// 启动监听路由
	router.Run(fmt.Sprintf("%s:%s", host, port))

}
