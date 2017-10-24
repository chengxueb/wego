package wego

import (
	"fmt"
	"net/http"
)

var (
	WegoApp *Application
	Config  *WegoConfig
)

func init() {
	Config = NewConfig()
	WegoApp = NewApplication()
}

func NewServer() *http.Server {
	s := &http.Server{}
	s.Addr = Config.WebConfig.Addr
	s.ReadTimeout = Config.WebConfig.ReadTimeout
	s.WriteTimeout = Config.WebConfig.WriteTimeout

	return s
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

	go func() {
		if err := WegoApp.Server.ListenAndServe(); err != nil {
			fmt.Println(err)
			appRunning <- false
		}
	}()

	<-appRunning
}
