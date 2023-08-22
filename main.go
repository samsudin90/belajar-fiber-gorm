package main

import (
	"belajar-fiber-gorm/config"
	"belajar-fiber-gorm/controllers/auth"
	"belajar-fiber-gorm/database"
	"belajar-fiber-gorm/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DatabaseInit()

	app := fiber.New()

	app.Post("/login", auth.Login)
	app.Post("/upload", auth.Upload)

	api := app.Group("/api")
	api.Static("/public", config.ProjectRootPathh+"/public/assets")
	user := api.Group("/user")

	routes.RouteInit(user)

	app.Listen(":8000")
}
