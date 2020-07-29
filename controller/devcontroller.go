package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(ctx *gin.Context) {
	user := ctx.PostForm("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": user})
}
