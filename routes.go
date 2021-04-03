package huck

import (
	"log"

	"github.com/labstack/echo/v4"
)

func (hk *HuckKernel) ParseConfigHandler(conf *Configuration) {
	for _, counter := range conf.Counter {
		hk.stat[counter] = NewCounterHandler(counter)
	}
}

func (hk *HuckKernel) RegisterHandlerToEcho(e *echo.Echo) {
	for path, stat := range hk.stat {
		log.Println("Register Handler to " + path)
		e.GET(path, stat.Handler)
	}
}
