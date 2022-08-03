package main

import (
	"github.com/a-h/rest-api/api/handlers/greet"
	"github.com/akrylysov/algnhsa"
)

func greeting(name string) string {
	return name
}

func main() {
	h := greet.NewHomeHandler(greeting)
	algnhsa.ListenAndServe(h, nil)
}
