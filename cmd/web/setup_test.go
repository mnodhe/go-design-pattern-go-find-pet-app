package main

import (
	"go-design-pattern-go-find-pet-app/configuration"
	"os"
	"testing"
)

var testApp application

func TestMain(m *testing.M) {
	testApp = application{
		App: configuration.New(nil),
	}
	os.Exit(m.Run())
}
