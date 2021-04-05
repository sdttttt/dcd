package huck

import (
	"net/http"
	"sync/atomic"

	"github.com/labstack/echo/v4"
)

type Counter struct {
	path  string
	count uint64
	store *CounterStorage
}

func NewCounterHandler(path string) *Counter {
	count := uint64(0)

	if counterStore.counterMap[path] != 0 {
		count = counterStore.counterMap[path]
	}

	return &Counter{
		path:  path,
		count: count,
		store: counterStore,
	}
}

func (counter *Counter) Handler(ctx echo.Context) error {
	atomic.AddUint64(&counter.count, 1)
	counter.store.Save(counter.path, counter.count)

	resp := &HuckResponse{
		Code: 200,
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (counter *Counter) Path() string {
	return counter.path
}

func (counter *Counter) GetCount() uint64 {
	return counter.count
}
