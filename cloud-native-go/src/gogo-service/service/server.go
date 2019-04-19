package main

import (
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"net/http"
)

func NewServer() *negroni.Negroni  {
	formatter := render.New(render.Options{IndentJSON: true })
	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render)  {
	mx.HandleFunc("/test", testHandler(formatter)).Methods("GET")
}

func testHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			Test string
		}{"This is a test"})
	}
}