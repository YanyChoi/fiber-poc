package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Router struct {
	app *fiber.App
}

func NewRouter() *Router {
	app := fiber.New()
	app.Use(logger.New())
	addOAuthRouters(app)
	return &Router{app: app}
}

func (r *Router) Run() {
	r.app.Listen(":8080")
}