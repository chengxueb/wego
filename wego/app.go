package wego

import (
	"fmt"
	"net/http"
)

var (
	WegoApp *Application
)

func init() {
	WegoApp = NewApplication()
}

func NewServer() *http.Server {
	return &http.Server{}
}

func NewApplication() *Application {
	cr := NewControllerRegister()
	s := NewServer()
	s.Handler = cr
	app := &Application{Handler: cr, Server: s}
	return app
}

type Application struct {
	Handler *ControllerRegister
	Server  *http.Server
}

func (this *Application) Run() {
	var appRunning chan bool = make(chan bool)

	WegoApp.Server.Addr = WConfig.WebConfig.Addr
	WegoApp.Server.ReadTimeout = WConfig.WebConfig.ReadTimeout
	WegoApp.Server.WriteTimeout = WConfig.WebConfig.WriteTimeout

	go func() {
		if err := WegoApp.Server.ListenAndServe(); err != nil {
			fmt.Println(err)
			appRunning <- false
		}
	}()

	<-appRunning
}
