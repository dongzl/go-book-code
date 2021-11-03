package v1

import "testing"

func TestOpen(t *testing.T) {
	db, err := Open("/tmp/mindb")
	if err != nil {
		t.Error(err)
	}
	t.Log(db)
}
