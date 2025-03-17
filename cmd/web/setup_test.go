package main

import (
	"go-design-pattern-go-find-pet-app/models"
	"os"
	"testing"
)

var testApp application

func TestMain(m *testing.M) {
	testApp = application{
		Models: *models.New(nil),
	}
	os.Exit(m.Run())
}
