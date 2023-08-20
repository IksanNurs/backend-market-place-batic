package database

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	db, err = gorm.Open(""+os.Getenv("DBMS"), ""+os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+")/"+os.Getenv("DB_NAME")+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	// err = db.AutoMigrate(&models.User{}).Error
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func GetDB() *gorm.DB {
	return db
}
