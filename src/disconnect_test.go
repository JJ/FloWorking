package main

import (
	"testing"
)

func TestDeleteAct(t *testing.T) {
	cl, _, _ := connectToMongo()
	err2 := disconnectFromMongo(cl)
	if err2 != nil {
		t.Fail()
	}
}