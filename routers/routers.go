package routers

import (
	"net/http"
	"zhny/controller"
	"zhny/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	//测试的
	r.GET("/ip", func(c *gin.Context) {
		c.String(http.StatusOK, c.ClientIP())
	})
	r.POST("/test", controller.Test)
	//注册登录
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	//用户信息
	r.POST("/console", middleware.AuthMiddleware(), controller.Console)
	r.POST("/userinfo", middleware.AuthMiddleware(), controller.Userinfo)
	//用户绑定控制设备
	r.POST("/binddev", middleware.AuthMiddleware(), controller.Binddev)
	r.POST("/ctrldev", middleware.AuthMiddleware(), controller.Ctrldev)
	//读取设备信息
	r.POST("/readall", middleware.AuthMiddleware(), controller.Readall)
	r.POST("/readdev", middleware.AuthMiddleware(), controller.Readdev)
	//设备报告执行命令
	r.POST("/devreport", controller.Devreport)
	r.POST("/devactive", controller.Devactive)
	//添加设备
	r.POST("/devadd", controller.Devadd)
	return r
}
