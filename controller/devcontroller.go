package controller

import (
	"net/http"
	"time"
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
	if dev.Belong != "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "设备已被绑定"})
		return
	}
	// 判断设备密码是否正确
	if devpass != dev.Devpass {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "设备密码错误"})
		return
	}
	DB.Table("devs").Model(&dev).Update("belong", name)
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": devid, "msg": "设备绑定成功"})
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
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": cmd, "msg": "命令发布成功"})
}

func Devreport(ctx *gin.Context) {
	DB := database.GetDB()
	devid := ctx.PostForm("devid")
	devpass := ctx.PostForm("devpass")
	temperature := ctx.PostForm("temperature")
	humidity := ctx.PostForm("humidity")
	co2 := ctx.PostForm("co2")

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
	devdata := model.Devdata{
		Devid:          devid,
		Devtemperature: temperature,
		Devhumidity:    humidity,
		Devco2:         co2,
		Time:           time.Now(),
	}
	DB.Table("devdata").Create(&devdata)
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": devdata})
}

func Devactive(ctx *gin.Context) {
	DB := database.GetDB()
	devid := ctx.PostForm("devid")
	devpass := ctx.PostForm("devpass")
	var dev model.Devs
	DB.Table("devs").Where("devid = ?", devid).First(&dev)
	//判断设备是否存在
	if dev.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "设备不存在"})
		return
	}
	if dev.Devpass != devpass {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "设备密码错误"})
		return
	}
	if dev.Cmd == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "无命令"})
		return
	}
	DB.Table("devs").Model(&dev).Update("cmd", "")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "done"})
}

func Readall(ctx *gin.Context) {
	name, _ := ctx.Get("name")
	DB := database.GetDB()
	var devs []model.Devs
	var devlist []string
	var temperaturelist []string
	var humiditylist []string
	var co2list []string
	var devpd []model.Dev
	DB.Table("devs").Where("belong = ?", name).Scan(&devs)
	for _, dev := range devs {
		var devdata model.Devdata
		DB.Table("devdata").Where("devid = ?", dev.Devid).Find(&devdata)
		devp := model.Dev{
			Devid:          dev.Devid,
			Belong:         dev.Belong,
			Cmd:            dev.Cmd,
			Devtemperature: devdata.Devtemperature,
			Devhumidity:    devdata.Devhumidity,
			Devco2:         devdata.Devco2,
			Time:           devdata.Time,
		}
		devpd = append(devpd, devp)
		devlist = append(devlist, dev.Devid)
		temperaturelist = append(temperaturelist, devdata.Devtemperature)
		humiditylist = append(humiditylist, devdata.Devhumidity)
		co2list = append(co2list, devdata.Devco2)
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": devpd, "devs": devlist, "temperatures": temperaturelist, "humiditys": humiditylist, "co2s": co2list})
}

func Readdev(ctx *gin.Context) {
	name, _ := ctx.Get("name")
	devid := ctx.PostForm("devid")
	DB := database.GetDB()
	var dev model.Devs
	DB.Table("devs").Where("devid = ?", devid).First(&dev)
	//判断设备是否存在
	if dev.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "设备不存在"})
		return
	}
	if dev.Belong != name {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "该设备不属于你"})
		return
	}
	var datas []model.Devdata
	var Time []time.Time
	var Devtemperature []string
	var Devhumidity []string
	var Devco2 []string
	DB.Table("devdata").Where("devid = ?", devid).Scan(&datas)
	for _, devdata := range datas {
		Time = append(Time, devdata.Time)
		Devtemperature = append(Devtemperature, devdata.Devtemperature)
		Devhumidity = append(Devhumidity, devdata.Devhumidity)
		Devco2 = append(Devco2, devdata.Devco2)
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "devdatas": gin.H{"Time": Time, "Devtemperature": Devtemperature, "Devhumidity": Devhumidity, "Devco2": Devco2}})
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
