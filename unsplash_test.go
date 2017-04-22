package main

import (
	"reflect"
	"testing"

	"github.com/Sirupsen/logrus"
)

var logger = logrus.New()

func TestGetRandomImage(t *testing.T) {
	photo, err := getRandomImage(logger)
	if err != nil {
		t.Error("Connection error: ", err.Error())
	}
	if reflect.TypeOf(photo).Name() != "Photo" {
		t.Error("Return is not a Photo struct")
	}
}

func TestSearchImages(t *testing.T) {
	photos, err := searchImages("photo", logger)
	if err != nil {
		t.Error("Connection error: ", err.Error())
	}
	if reflect.TypeOf(photos).String() != "[]main.Photo" {
		t.Error("Result is not a Photo objects slice")
	}
}
