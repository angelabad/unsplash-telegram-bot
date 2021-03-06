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
	resty.SetQueryParams(map[string]string{
		"query":    query,
		"page":     "1",
		"per_page": "5",
	})
	resp, err := resty.R().
		Get("https://api.unsplash.com/search/photos")
	if err != nil {
		return photos, err
	}

	results := gjson.Get(string(resp.Body()), "results")
	err = json.Unmarshal([]byte(results.String()), &photos)

	return photos, err
}

func getFeaturedImages(logger *logrus.Logger) ([]Photo, error) {
	var photos []Photo

	resty.SetHeader("Authorization", "Client-ID "+unsplashID)
	resty.SetHeader("Accept-Version", "v1")
	resty.SetQueryParams(map[string]string{
		"page":     "1",
		"per_page": "5",
		"order_by": "popular",
	})
	resp, err := resty.R().
		Get("https://api.unsplash.com/photos")
	if err != nil {
		return photos, err
	}

	if err = json.Unmarshal(resp.Body(), &photos); err != nil {
		return photos, err
	}

	return photos, nil
}

func getLatestImages(logger *logrus.Logger) ([]Photo, error) {
	var photos []Photo

	resty.SetHeader("Authorization", "Client-ID "+unsplashID)
	resty.SetHeader("Accept-Version", "v1")
	resty.SetQueryParams(map[string]string{
		"page":     "1",
		"per_page": "5",
		"order_by": "latest",
	})
	resp, err := resty.R().
		Get("https://api.unsplash.com/photos")
	if err != nil {
		return photos, err
	}

	if err = json.Unmarshal(resp.Body(), &photos); err != nil {
		return photos, err
	}

	return photos, nil
}
