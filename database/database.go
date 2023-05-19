package database

import (
	"fmt"
	"os"

	"github.com/RianIhsan/ex-go-crud-icc/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Tidak bisa konek ke database")
	} else {
		fmt.Println("Telah terkoneksi ke database")
	}

	database.AutoMigrate(&models.User{})

	DB = database

}
