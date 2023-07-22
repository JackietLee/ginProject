package routers

import (
	"ginProject/controllers/admin"
	"ginProject/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin", middleware.InitMiddleware)
	{
		adminRouters.GET("/", func(c *gin.Context) {
			c.String(200, "后台首页")
		})
		adminRouters.GET("/userIndex", admin.UserController{}.UserIndex)
		adminRouters.GET("/userAdd", admin.UserController{}.UserAdd)
		adminRouters.GET("/userDel", admin.UserController{}.UserDel)
		adminRouters.GET("/userUpdate", admin.UserController{}.UserUpdate)
	}

}
