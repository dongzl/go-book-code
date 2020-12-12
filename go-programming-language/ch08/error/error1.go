package main

import (
	"database/sql"
	"fmt"
)

func foo() error {
	return sql.ErrNoRows
}

func bar() error {
	return foo()
}

func main() {
	err := bar()
	if err != nil {
		fmt.Printf("got err, %+v\n", err)
	}
}
