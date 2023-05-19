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
	if err := godotenv.Load(); err != nil {
		log.Fatal("Tidak bisa memuat file .env")
	}

	database.ConnectDB()

	port := os.Getenv("PORT")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://fscrud.netlify.app",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	routes.RunRoute(app)

	app.Listen(":" + port)

}
