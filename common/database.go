package common

import (
	"github.com/jinzhu/gorm"
	"oceanlearn.teach/ginessential/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := "mysql"
	args := "root:000820@tcp(127.0.0.1:3306)/ginessential?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("fail to connect the database, err:" + err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
