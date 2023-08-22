package routes

import (
	usercontroller "belajar-fiber-gorm/controllers/userController"
	"belajar-fiber-gorm/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r fiber.Router) {
	r.Get("/", middleware.Auth, usercontroller.Index)
	r.Get("/:id", usercontroller.Show)
	r.Post("/", usercontroller.Create)
	r.Put("/:id", usercontroller.Update)
	r.Delete("/:id", usercontroller.Delete)
}
