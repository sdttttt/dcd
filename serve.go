package huck

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type DataStatistical interface {
	Handler(echo.Context) error
}

type HuckResponse struct {
	Code uint8 `json: "code"`
}

func Run(conf *string) {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	// Routes
	e.GET("/", hello)

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
