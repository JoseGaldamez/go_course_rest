package database

import (
	"fmt"
	"os"

	"github.com/JoseGaldamez/go_course_rest/internal/users"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConfigDatabase() *gorm.DB {
	dsn := getDSN() // URL
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// db = db.Debug() // use in debug
	_ = db.AutoMigrate(&users.User{}) // create table with the struct information

	return db
}

func getDSN() string {
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"))
}
