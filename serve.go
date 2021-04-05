package huck

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run(conf_filename string) {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	conf := FromConfigFile(conf_filename)

	kernel.ParseConfigHandler(conf)
	kernel.RegisterHandlerToEcho(e)

	log.Println(conf.Counter)

	// Start server
	e.Logger.Fatal(e.Start(":10086"))
}
