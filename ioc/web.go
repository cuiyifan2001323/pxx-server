// @Author 齐静春
// @Date 2023-11-14 21:48:00

package ioc

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"pxx-server/web"
	"strings"
)

func InitWebServer(mlds []gin.HandlerFunc, userHandles *web.HandleUser) *gin.Engine {
	g := gin.Default()
	g.Use(mlds...)
	userHandles.RegisterRoutes(g)
	return g
}

func InitGinMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		cors.New(cors.Config{
			// 业务请求中可以带上的头
			AllowHeaders: []string{"Content-Type", "Authorization"},
			// 允许前端访问的响应头字段
			ExposeHeaders: []string{"Authorization"},
			// 是否允许带上用户认证 信息（比如 cookie）。
			AllowCredentials: true,
			// 哪些来源是允许的
			AllowOriginFunc: func(origin string) bool {
				return strings.HasPrefix(origin, "http://localhost")
			},
		}),
		//(&middleware.LoginMiddlewareBuilder{}).CheckLogin(),
	}
}
