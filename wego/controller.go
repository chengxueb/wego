package wego

import (
	"errors"
	"net/http"
	"reflect"
	"strings"
)

func NewControllerRegister() *ControllerRegister {
	return &ControllerRegister{}
}

type ControllerRegister struct {
	RequestPath   string
	RequestMethod string
}

func (this *ControllerRegister) InvokeMethod() error {
	r, ok := UrlRouter[this.RequestPath]

	if ok && strings.ToUpper(r.RequestMethod) == strings.ToUpper(this.RequestMethod) {
		v := reflect.ValueOf(r.Controller)
		v.MethodByName(r.Method).Call(nil)
		return nil
	}

	return errors.New("Error: This is controller or method not found!")
}

func (this *ControllerRegister) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	this.RequestPath = r.URL.Path
	this.RequestMethod = r.Method

	this.InvokeMethod()
}
