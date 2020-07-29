package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(ctx *gin.Context) {
	user := ctx.PostForm("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": user})
}
func Binddev(ctx *gin.Context) {
	devid := ctx.PostForm("devid")
	devpass := ctx.PostForm("devpass")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": user})
}
func Ctrldev(ctx *gin.Context) {
	user := ctx.PostForm("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": user})
}
