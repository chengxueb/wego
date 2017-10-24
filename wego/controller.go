package wego

import (
	"fmt"
	"net/http"
)

func NewControllerRegister() *ControllerRegister {
	return &ControllerRegister{}
}

type ControllerRegister struct {
}

func (this *ControllerRegister) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request")
}
