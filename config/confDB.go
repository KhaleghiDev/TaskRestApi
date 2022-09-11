package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfDb struct {
	Hose     string
	Db       string
	Username string
	Password string
	Port     string
}

func ConnMysql(data ConfDb) (db *gorm.DB) {
	dsn := data.Username + ":" + data.Password + "@tcp(" + data.Hose + ":" + data.Port + ")/" + data.Db + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Print(db)
	if err != nil {
		fmt.Print("Error 1 : config start db ", err)
	}
	return
}
