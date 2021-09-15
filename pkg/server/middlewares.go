package server

import "github.com/gofiber/fiber/v2/middleware/logger"

func (app *FiberApp) middlewares() *FiberApp {
	app.Use(logger.New())
	return app
}
