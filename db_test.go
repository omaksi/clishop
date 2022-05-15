package db

import "testing"

func TestConnect(t *testing.T) {
	db := connect()
	if db == nil {
		t.Log("could not connect to db")
		t.Fail()
	}
}
