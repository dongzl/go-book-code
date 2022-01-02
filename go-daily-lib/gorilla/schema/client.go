package main

import (
	"fmt"
	"github.com/gorilla/schema"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	encoder = schema.NewEncoder()
)

func main() {
	client := &http.Client{}
	form := url.Values{}

	u := &User{
		Username: "dj",
		Password: "handsome",
	}
	encoder.Encode(u, form)

	res, _ := client.PostForm("http://localhost:8080/login", form)
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(data))
	res.Body.Close()
}
