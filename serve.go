package huck

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Run huck application.
func Run(conFilename string) {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	conf = FromConfigFile(conFilename)

	kernel.ParseConfigHandler(conf)
	kernel.RegisterHandlerToEcho(e)

	log.Println(conf.Counter)

	// Start server
	e.Logger.Fatal(e.Start(":10086"))
}
