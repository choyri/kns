package main

import (
	"github.com/choyri/kns/controller"
	"net/http"
)

func InitRoute() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/import", controller.Import)
	http.HandleFunc("/search", controller.Search)
	http.HandleFunc("/export", controller.Export)
}
