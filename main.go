package main

import (
	"log"
	"os"

	"github.com/RianIhsan/ex-go-crud-icc/database"
	"github.com/RianIhsan/ex-go-crud-icc/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	database.ConnectDB()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Tidak bisa memuat file .env")
	}

	port := os.Getenv("PORT")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, X-Requested-With, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: false,
	}))

	routes.RunRoute(app)

	app.Listen(":" + port)

}
