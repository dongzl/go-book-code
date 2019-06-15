package main

import (
	"flag"
	"fmt"
)

func main()  {
	//var name string
	//flag.StringVar(&name, "name", "everyone", "The greeting opject.")
	//var name = *flag.String("name", "everyone", "The greeting opject.")
	name := *flag.String("name", "everyone", "The greeting opject.")
	flag.Parse()
	fmt.Printf("Hello, %v!\n", name)
}
