package router

import (
	"jinrongdati/app/controller"
	"jinrongdati/app/middleware"

	"github.com/gin-gonic/gin"
)

// APIRouter ...
func APIRouter(r *gin.Engine) *gin.Engine {

	// Index
	r.GET("/", controller.Index{}.Index)

	// 定义 api 路由
	api := r.Group("/api")
	{

		// api 下且无授权
		// api.GET("/", controller.User{}.Index)

		// api 下有授权
		api.Use(middleware.AuthenticationRequired())
		{

			// user
			// api.GET("/user", controller.User{}.Index)

		}
	}

	return r

}
