package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetSmallImage(t *testing.T) {
	// TODO: This is duplicate code
	// TODO: logger is from unsplash_test
	photo, err := getRandomImage(logger)
	if err != nil {
		t.Error("Connection error: ", err.Error())
	}
	if reflect.TypeOf(photo).Name() != "Photo" {
		t.Error("Return is not a Photo struct")
	}

	_, err = url.ParseRequestURI(photo.Urls.Small)
	if err != nil {
		t.Error("Photo url is not valid url")
	}

	smallphoto, err := photo.getSmallImage()
	if err != nil {
		t.Error("Error getting image: ", err.Error())
	}

	//Check if result is []byte type
	var typeOfBytes = reflect.TypeOf([]byte(nil))
	if reflect.TypeOf(smallphoto) != typeOfBytes {
		t.Error("Resulting image is not []byte type")
	}
}
