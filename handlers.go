package huck

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Counter struct {
	path  string
	count uint64
}

func NewCounterHandler(path string) *Counter {
	return &Counter{
		path:  path,
		count: 0,
	}
}

func (counter *Counter) Handler(ctx echo.Context) error {
	counter.count++

	resp := &HuckResponse{
		Code: 200,
	}

	return ctx.JSON(http.StatusOK, resp)
}
