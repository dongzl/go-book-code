package main

import (
	"github.com/unrolled/render"
	"net/http"
)

func createMatchHandler(formatter *render.Render) http.HandlerFunc  {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			Test string
		}{ "This is a test"})
	}
}
