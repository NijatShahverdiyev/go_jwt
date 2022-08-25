package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *gorm.DB

func init() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		dbHost, username, dbName, password)
	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		log.Fatalln("Wrong database url")
	}
	db = conn
	fmt.Println("connected to database")
	db.Debug().AutoMigrate(&User{})

}

func GetDB() *gorm.DB {
	return db
}
