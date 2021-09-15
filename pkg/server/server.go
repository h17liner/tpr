package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type FiberApp struct {
	*fiber.App
}

func setup() *FiberApp {

	app := &FiberApp{App: fiber.New()}

	return app.middlewares().routes()
}

func Run() {
	app := setup()
	panic(app.Listen(viper.GetString("host") + ":" + viper.GetString("port")))
}
