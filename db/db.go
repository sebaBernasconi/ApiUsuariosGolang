package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConectarDb() {

	dsn := "root:1234@tcp(127.0.0.1:3306)/api_usuarios_golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error al conectarse con la BD")
	}

	fmt.Print("Conexion exitosa")
	DB = db
}
