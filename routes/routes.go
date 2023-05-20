package routes

import (
	"github.com/RianIhsan/ex-go-crud-icc/controllers"
	"github.com/gofiber/fiber/v2"
)

func RunRoute(app *fiber.App) {
	app.Get("/api/reads", controllers.Reads)
	app.Get("/api/read/:id", controllers.Read)
	app.Post("/api/create", controllers.Create)
	app.Put("/api/update/:id", controllers.Update)
	app.Delete("/api/delete/:id", controllers.Delete)
}
