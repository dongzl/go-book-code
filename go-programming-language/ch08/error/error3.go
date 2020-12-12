package main

import (
	"database/sql"
	"fmt"
)

func foo3() error {
	return fmt.Errorf("foo err, %v", sql.ErrNoRows)
}

func bar3() error {
	return foo3()
}

func main() {
	err := bar3()
	if err == sql.ErrNoRows {
		fmt.Printf("data not found, %+v\n", err)
		return
	}
	if err != nil {
		fmt.Printf("unknown")
		return
	}
}
