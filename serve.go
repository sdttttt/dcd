package huck

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type DataStatistical interface {
	Handler(ctx echo.Context) error

	Path() string
}

type HuckResponse struct {
	Code uint8 `json: "code"`
}

func Run(conf_filename string) {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	// Routes
	e.GET("/", hello)
	conf := FromConfigFile(conf_filename)

	kernel.ParseConfigHandler(conf)
	kernel.RegisterHandlerToEcho(e)

	log.Println(conf.Counter)

	// Start server
	e.Logger.Fatal(e.Start(":10086"))
}

// Handler
func hello(ctx echo.Context) error {

	resp := &HuckResponse{
		Code: 200,
	}

	return ctx.JSON(http.StatusOK, resp)
}
