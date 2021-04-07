package huck

import (
	"github.com/labstack/echo/v4"
)

var kernel = NewHuckKernel()

// Response is Default huck http response.
type Response struct {
	Code uint8 `json: "code"`
}

// Kernel have all huck state.
type Kernel struct {
	stat map[string]DataStatistical
}

// DataStatistical is a basic statistician common implementation interface.
type DataStatistical interface {
	Handler(ctx echo.Context) error
	Path() string
}

// NewHuckKernel to initializer a HuckKernel.
func NewHuckKernel() *Kernel {
	return &Kernel{
		stat: make(map[string]DataStatistical),
	}
}
