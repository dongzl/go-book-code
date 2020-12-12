package main

import (
	"database/sql"
	"fmt"
)

func foo1() error {
	return sql.ErrNoRows
}

func bar1() error {
	return foo1()
}

func main() {
	err := bar1()
	if err == sql.ErrNoRows {
		fmt.Printf("data not found, %+v\n", err)
		return
	}
	if err != nil {
		// Unknown error
	}
}
