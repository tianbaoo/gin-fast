package routers

import (
	"gin-fast/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterCmsRouter(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/file", handlers.UploadFileHandler)
	routerGroup.GET("/file/:id", handlers.DownloadFileHandler)
}
