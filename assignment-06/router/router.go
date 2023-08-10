package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ngoduongkha/go-2023/assignment-06/controller/middleware"
	v1 "github.com/ngoduongkha/go-2023/assignment-06/controller/v1"
	"github.com/ngoduongkha/go-2023/assignment-06/utils"
)

func InitializeRouter() (router *gin.Engine) {
	router = gin.Default()
	v1route := router.Group("/api/v1")
	v1route.Use(
		middleware.CORSMiddleware,
		middleware.AuthMiddleware,
	)
	{
		auth := v1route.Group("/auth")
		{
			auth.POST("/login", v1.POSTLogin)
			auth.POST("/signup", v1.POSTRegister)
		}
		user := v1route.Group("/user")
		{
			user.GET("/:username", utils.AuthOnly, v1.GETUser)
			user.PUT("", utils.AuthOnly, v1.PUTUser)
		}
	}
	return
}
