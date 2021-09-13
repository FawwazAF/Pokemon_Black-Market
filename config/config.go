package config

import (
	"project/pokemon/models"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var HTTP_PORT int

func InitDb() {

	//Set connection string here, use mysql username password and schema at your pc
	connectionString := "root:Minus12345@tcp(localhost:3306)/pokemon_schema?charset=utf8&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitPort() {
	var err error

	// HTTP_PORT = 8080 //Port Setting

	HTTP_PORT, err = strconv.Atoi("8080")
	if err != nil {
		panic(err)
	}
}

func InitMigrate() {
	DB.AutoMigrate(&models.Pokemon{})
	DB.AutoMigrate(&models.Seller{})
	DB.AutoMigrate(&models.Transaction{})
}
