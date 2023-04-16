package config

import (
	"fmt"
	"log"
	aData "todolist-api/features/activity/data"
	tData "todolist-api/features/todo/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(dc DBConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dc.Mysql_User, dc.Mysql_Password, dc.Mysql_Host, dc.Mysql_Port, dc.Mysql_DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("database connection error : ", err.Error())
		return nil
	}
	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(aData.Activity{})
	db.AutoMigrate(tData.Todo{})
}
