package boot

import (
	"log"
	"os"

	"github.com/RianIhsan/ex-go-crud-icc/config"
	"github.com/RianIhsan/ex-go-crud-icc/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func BootApp() {
	//LOAD ENV FILE
	if err := godotenv.Load(); err != nil {
		log.Fatal("Gagal memuat file ENV")
	}

	//Check Default PORT SERVER
	if portEnv := os.Getenv("PORT"); portEnv != "" {
		config.PORT = portEnv
	}

	config.BootDatabase()
	config.ConnectDB()
	config.RunMigration()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     config.AllowOrigins,
		AllowHeaders:     config.AllowHeaders,
		AllowCredentials: config.AllowCredentials,
	}))

	//Initialize Route
	routes.SetupRoute(app)

	//Initialize Server
	app.Listen(config.PORT)

}
