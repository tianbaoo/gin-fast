package middlewares

import (
	"gin-fast/handlers"
	"gin-fast/pkg/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

// AuthJwtTokenMiddleware jwt token验证有效性，并赋值ctx当前会话用户
func AuthJwtTokenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// method := ctx.Request.Method
		resp := handlers.Response{Ctx: ctx}
		// 从Request Header中获取token参数
		tokenStr := ctx.Request.Header.Get("Authentication")
		if len([]rune(tokenStr)) == 0 {
			resp.Unauthenticated("token不存在，请先登录")
			return
		}
		strs := strings.Split(tokenStr, " ")
		if strs[0] != "gin-fast" {
			resp.BadRequest("token格式不正确，${gin-fast token}")
		}
		// 将获取到的具体token参数进行校验
		claims, err := jwt.ValidateJwtToken(strs[1])
		if err != nil {
			resp.Unauthenticated("token校验失败，" + err.Error())
			return
		} else {
			// 解析获取用户ID
			CurrentUser := claims.GetUserByID()
			if CurrentUser != nil {
				// 返回值中把当前用户名返回
				ctx.Set("CurrentUser", CurrentUser)
			} else {
				resp.Unauthenticated("查询不到token对应用户")
				return
			}
		}
		// 让中间件执行下一步
		ctx.Next()
	}
}
