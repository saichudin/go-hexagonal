package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = ""
const DB_NAME = "majoo-emenu-19jun23"
const DB_HOST = "127.0.0.1"
const DB_PORT = "8888"

// const DB_USERNAME = "dev_zeppelin"
// const DB_PASSWORD = "h%r#<:z4r@j9~53A"
// const DB_NAME = "emenu"
// const DB_HOST = "10.99.65.226"
// const DB_PORT = "3306"

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectMysql()
	return Db
}

func connectMysql() *gorm.DB {
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(err)
	if err != nil {
		panic(err)
	}

	return db
}
