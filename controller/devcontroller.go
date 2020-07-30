package controller

import (
	"net/http"
	"zhny/database"
	"zhny/model"

	"github.com/gin-gonic/gin"
)

func Test(ctx *gin.Context) {
	user := ctx.PostForm("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": user})
}

func Binddev(ctx *gin.Context) {
	DB := database.GetDB()
	devid := ctx.PostForm("devid")
	devpass := ctx.PostForm("devpass")
	name, _ := ctx.Get("name")
	var dev model.Devs
	DB.Table("devs").Where("devid = ?", devid).First(&dev)
	//判断设备是否存在
	if dev.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "设备不存在"})
		return
	}
	// 判断设备密码是否正确
	if devpass != dev.Devpass {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "设备密码错误"})
		return
	}
	DB.Table("devs").Model(&dev).Update("belong", name)
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": devid, "msg": "设备绑定成功"})
	DB.Exec("CREATE TABLE ? AS SELECT * FROM devdata where 1=2", devid)
}

func Ctrldev(ctx *gin.Context) {
	DB := database.GetDB()
	devid := ctx.PostForm("devid")
	cmd := ctx.PostForm("cmd")
	name, _ := ctx.Get("name")
	var dev model.Devs
	DB.Table("devs").Where("devid = ?", devid).First(&dev)
	//判断设备是否存在
	if dev.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "设备不存在"})
		return
	}
	// 判断设备是否属于该用户
	if dev.Belong != name {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "此设备不属于你"})
		return
	}
	DB.Table("devs").Model(&dev).Update("cmd", cmd)
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": cmd})
}

func Devreport(ctx *gin.Context) {
	DB := database.GetDB()
	devid := ctx.PostForm("devid")
	devpass := ctx.PostForm("devpass")
	data := ctx.PostForm("data")
	var dev model.Devs
	DB.Table("devs").Where("devid = ?", devid).First(&dev)
	//判断设备是否存在
	if dev.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "设备不存在"})
		return
	}
	// 判断设备密码是否正确
	if devpass != dev.Devpass {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "设备密码存在"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": data})
}

func Devactive(ctx *gin.Context) {
	DB := database.GetDB()
	devid := ctx.PostForm("devid")
	cmd := ctx.PostForm("cmd")
	name, _ := ctx.Get("name")
	var dev model.Devs
	DB.Table("devs").Where("devid = ?", devid).First(&dev)
	//判断设备是否存在
	if dev.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "设备不存在"})
		return
	}
	// 判断设备是否属于该用户
	if dev.Belong != name {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "此设备不属于你"})
		return
	}
	DB.Table("devs").Model(&dev).Update("cmd", cmd)
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": cmd})
}

func Devadd(ctx *gin.Context) {
	DB := database.GetDB()
	devid := ctx.PostForm("devid")
	devpass := ctx.PostForm("devpass")
	newdev := model.Devs{
		Devid:   devid,
		Devpass: devpass,
	}
	DB.Create(&newdev)
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": newdev})
}
