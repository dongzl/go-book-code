package main

import (
	"fmt"
	"github.com/gorilla/schema"
	"io/ioutil"
	"net/http"
	"net/url"
)

type PersonClient struct {
	Name    string `schema:"name"`
	Age     int    `schema:"age"`
	Hobbies string `schema:"hobbies"`
}

var (
	encoderClient = schema.NewEncoder()
)

func main() {
	client := &http.Client{}
	form := url.Values{}

	p := &PersonClient{
		Name:    "dj",
		Age:     18,
		Hobbies: "Game,Programming",
	}
	encoderClient.Encode(p, form)

	res, _ := client.PostForm("http://localhost:8080/info", form)
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(data))
	res.Body.Close()
}
