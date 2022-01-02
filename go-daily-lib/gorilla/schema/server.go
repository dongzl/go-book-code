package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"log"
	"net/http"
	"reflect"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Login</title>
</head>
<body>
<form action="/login" method="post">
  <label>Username:</label>
  <input name="username"><br>
  <label>Password:</label>
  <input name="password" type="password"><br>
  <button type="submit">登录</button>
</form>
</body>
</html>`)
}

type User struct {
	Username string `schema:"username"`
	Password string `schema:"password"`
}

var (
	decoder = schema.NewDecoder()
)

func dologin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	u := User{}
	decoder.Decode(&u, r.PostForm)
	if u.Username == "dj" && u.Password == "handsome" {
		http.Redirect(w, r, "/", 301)
		return
	}

	http.Redirect(w, r, "/login", 301)
}

type Person struct {
	Name    string   `schema:"name"`
	Age     int      `schema:"age"`
	Hobbies []string `schema:"hobbies"`
}

func init() {
	decoder.RegisterConverter([]string{}, func(s string) reflect.Value {
		return reflect.ValueOf(strings.Split(s, ","))
	})
}

func doinfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	p := Person{}
	decoder.Decode(&p, r.PostForm)

	fmt.Println(p)
	fmt.Fprintf(w, "Name:%s Age:%d Hobbies:%v", p.Name, p.Age, p.Hobbies)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.Handle("/login", handlers.MethodHandler{
		"GET":  http.HandlerFunc(login),
		"POST": http.HandlerFunc(dologin),
	})
	log.Fatal(http.ListenAndServe(":8080", r))
}
