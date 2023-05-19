package config

import (
	"fmt"
	"log"
	"os"

	"github.com/RianIhsan/ex-go-crud-icc/database"
	"github.com/RianIhsan/ex-go-crud-icc/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db_name = ""
	db_user = ""
	db_pass = ""
	db_host = ""
	db_port = ""
)

func BootDatabase() {
	//Check Default database NAME
	if dbNameEnv := os.Getenv("DB_NAME"); dbNameEnv != "" {
		db_name = dbNameEnv
	}
	//Check default Database USER
	if dbUserEnv := os.Getenv("DB_USER"); dbUserEnv != "" {
		db_user = dbUserEnv
	}
	//Check default Database PASSWORD
	if dbPassEnv := os.Getenv("DB_PASS"); dbPassEnv != "" {
		db_pass = dbPassEnv
	}
	//Check default Database HOST
	if dbHostEnv := os.Getenv("DB_HOST"); dbHostEnv != "" {
		db_host = dbHostEnv
	}
	//Check Default Database PORT
	if dbPortEnv := os.Getenv("DB_PORT"); dbPortEnv != "" {
		db_port = dbPortEnv
	}
}

func ConnectDB() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_pass, db_host, db_port, db_name)
	database.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Tidak bisa terkoneksi dengan database")
	} else {
		fmt.Println("Berhasil terhubung ke database")
	}
}

func RunMigration() {
	if err := database.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Gagal migrasi ke database")
	} else {
		fmt.Println("Data berhasil dimigrasikan ke database")
	}
}
