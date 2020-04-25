package main

import (
	"github.com/choyri/kns/controller"
	"net/http"
)

func InitRoute() {
	http.HandleFunc("/", controller.Index)
}
