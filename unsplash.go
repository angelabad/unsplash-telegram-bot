package main

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"
	"github.com/tidwall/gjson"

	resty "gopkg.in/resty.v0"
)

func getRandomImage(logger *logrus.Logger) (Photo, error) {

	var photo Photo

	resty.SetHeader("Authorization", "Client-ID "+unsplashID)
	resty.SetHeader("Accept-Version", "v1")
	resp, err := resty.R().Get("https://api.unsplash.com/photos/random")
	if err != nil {
		return photo, err
	}

	if err = json.Unmarshal(resp.Body(), &photo); err != nil {
		return photo, err
	}

	return photo, nil
}

func searchImages(query string, logger *logrus.Logger) ([]Photo, error) {
	var photos []Photo

	resty.SetHeader("Authorization", "Client-ID "+unsplashID)
	resty.SetHeader("Accept-Version", "v1")
	resp, err := resty.R().
		SetQueryParam("query", query).
		Get("https://api.unsplash.com/search/photos")
	if err != nil {
		return photos, err
	}

	results := gjson.Get(string(resp.Body()), "results")
	if err = json.Unmarshal([]byte(results.String()), &photos); err != nil {
		return photos, err
	}

	return photos, nil
}

func getFeaturedImages(logger *logrus.Logger) ([]Photo, error) {
	var photos []Photo

	resty.SetHeader("Authorization", "Client-ID "+unsplashID)
	resty.SetHeader("Accept-Version", "v1")
	resp, err := resty.R().
		Get("https://api.unsplash.com/photos/curated")
	if err != nil {
		return photos, err
	}

	if err = json.Unmarshal(resp.Body(), &photos); err != nil {
		return photos, err
	}

	return photos, nil
}
