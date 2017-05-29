package main

import (
	"log"
	"reflect"
	"testing"

	"github.com/Sirupsen/logrus"
)

var logger = logrus.New()

func init() {
	err := parseConfig()
	if err != nil {
		log.Panic("Error parsing configuration: ", err.Error())
	}
}

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

func TestGetFeaturedImages(t *testing.T) {
	photos, err := getFeaturedImages(logger)
	if err != nil {
		t.Error("Connection error: ", err.Error())
	}
	if reflect.TypeOf(photos).String() != "[]main.Photo" {
		t.Error("Result is not a Photo objects slice")
	}
}

func TestGetLatestImages(t *testing.T) {
	photos, err := getLatestImages(logger)
	if err != nil {
		t.Error("Connection error: ", err.Error())
	}
	if reflect.TypeOf(photos).String() != "[]main.Photo" {
		t.Error("Result is not a Photo objects slice")
	}
}

func TestFeaturedImages(t *testing.T) {
	photos, err := getFeaturedImages(logger)
	if err != nil {
		t.Error("Connection error: ", err.Error())
	}
	if reflect.TypeOf(photos).String() != "[]main.Photo" {
		t.Error("Result is not a Photo slice")
	}
}
