package dcd

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type DataStatistical interface {
	Handler(echo.Context) error
}

type DCDResponse struct {
	Code uint8 `json: "code"`
}

func Run(conf *string) {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":10086"))
}

// Handler
func hello(ctx echo.Context) error {

	resp := &DCDResponse{
		Code: 200,
	}

	return ctx.JSON(http.StatusOK, resp)
}
