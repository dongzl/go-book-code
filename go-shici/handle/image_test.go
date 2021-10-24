package handle

import (
	"go-shici/db"
	"testing"
)

func Test_CreateImg(t *testing.T) {
	poems, err := db.QueryPoemsByAuthor("李白")
	if err != nil {
		return
	}
	for index:= range poems{
		CreateShiImage(poems[index])
	}
}
