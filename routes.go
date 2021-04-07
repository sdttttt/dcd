package huck

import (
	"log"

	"github.com/labstack/echo/v4"
)

// ParseConfigHandler is All processors in the parse configuration struct are added to HuckKernel.
func (hk *Kernel) ParseConfigHandler(conf *Configuration) {
	for _, counter := range conf.Counter {
		hk.stat[counter.Path] = NewCounterHandler(counter.Name, counter.Path)
	}
}

// RegisterHandlerToEcho is Add all the processors in the HuckKernel to Echo.
func (hk *Kernel) RegisterHandlerToEcho(e *echo.Echo) {
	for path, stat := range hk.stat {
		log.Println("Register Handler to " + path)
		e.GET(path, stat.Handler)
	}
}
