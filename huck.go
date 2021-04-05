package huck

import (
	"github.com/labstack/echo/v4"
)

var kernel *HuckKernel = NewHuckKernel()

type HuckResponse struct {
	Code uint8 `json: "code"`
}

type HuckKernel struct {
	stat map[string]DataStatistical
}

type DataStatistical interface {
	Handler(ctx echo.Context) error
	Path() string
}

func NewHuckKernel() *HuckKernel {
	return &HuckKernel{
		stat: make(map[string]DataStatistical),
	}
}
