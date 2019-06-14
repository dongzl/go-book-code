package main

import (
	"flag"
	"./lib"
)

var name string

func init()  {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main()  {
	flag.Parse()
	lib.Hello(name)
}
