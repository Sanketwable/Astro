package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

var (
	PORT                = 0
	DBURL               = ""
	DBDRIVER            = ""
	DBPORT              = ""
	HOST                = ""
)

// Load is a func
func Load() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error is : ", err)
	}
	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		PORT = 8080
	}

	DBDRIVER = os.Getenv("DB_DRIVER")
	DBPORT = os.Getenv("DB_PORT")
	DBURL = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),os.Getenv("DB_NAME"))

}
