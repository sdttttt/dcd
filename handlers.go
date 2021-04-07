package huck

import (
	"net/http"
	"sync/atomic"

	"github.com/labstack/echo/v4"
)

// Counter statisticians.
type Counter struct {
	name  string
	path  string
	count uint64
	store *CounterStorage
}

// NewCounterHandler to initializer a Counter.
func NewCounterHandler(name string, path string) *Counter {
	count := uint64(0)

	if counterStore.counterMap[name] != 0 {
		count = counterStore.counterMap[name]
	}

	return &Counter{
		name:  name,
		path:  path,
		count: count,
		store: counterStore,
	}
}

// Handler is Counter logic.
func (counter *Counter) Handler(ctx echo.Context) error {
	atomic.AddUint64(&counter.count, 1)
	counter.store.Save(counter.name, counter.count)

	resp := &Response{
		Code: 200,
	}

	return ctx.JSON(http.StatusOK, resp)
}

// Path is route url.
func (counter *Counter) Path() string {
	return counter.path
}

// GetCount is get count of counter.
func (counter *Counter) GetCount() uint64 {
	return counter.count
}
