package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"zhny/config"
	"zhny/database"
	"zhny/routers"
)

func main() {
	config.InitConfig()
	db := database.InitDB()
	defer db.Close()
	r := gin.Default()
	r = routers.CollectRouter(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}