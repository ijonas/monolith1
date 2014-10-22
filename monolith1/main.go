package main

import (
	"github.com/GeertJohan/go.rice"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(rice.MustFindBox("assets").HTTPBox()))
	http.ListenAndServe(":8080", nil)
}
