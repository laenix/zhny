package routers

import (
	"net/http"
	"zhny/controller"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.GET("/ip", func(c *gin.Context) {
		c.String(http.StatusOK, c.ClientIP())
	})
	r.POST("/test", controller.Test)
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	return r
}
