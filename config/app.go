package config

import (
	"fmt"

	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"github.com/joho/godotenv"
	"os"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func Connect(){
	// d,err := gorm.Open("mysql","DSN=qk90y9a8xzdyw06vgmu8:pscale_pw_sRfwX7WZky4HI1d4ntgFoFZe5zfSSV6XSMLcJdUKWPu@tcp(aws.connect.psdb.cloud)/budget-app?tls=true")

	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}
	
	d, err := gorm.Open(mysql.Open(os.Getenv("DSN")))

	if err != nil {
		panic(err)
		
	}

	db = d
	fmt.Println("Connected to DB")
}

func GetDb() *gorm.DB {
	return db
}