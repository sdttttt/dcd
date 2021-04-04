package huck

import (
	"net/http"
	"sync/atomic"

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
	atomic.AddUint64(&counter.count, 1)

	resp := &HuckResponse{
		Code: 200,
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (counter *Counter) Path() string {
	return counter.path
}

func (counter *Counter) Persistence() {

}

func (counter *Counter) GetCount() uint64 {
	return counter.count
}
