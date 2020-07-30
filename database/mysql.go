package database

import (
	"fmt"

	"zhny/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	//加载config中的参数
	driverName := viper.GetString("mysql.driverName")
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	datebase := viper.GetString("mysql.database")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	charset := viper.GetString("mysql.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		datebase,
		charset,
	)
	//连接数据库
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	//自动建立表
	db.SingularTable(true)
	db.AutoMigrate(&model.Users{})
	db.AutoMigrate(&model.Devs{})
	db.AutoMigrate(&model.Devdata{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
