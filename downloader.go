package main

import (
	"fmt"

	resty "gopkg.in/resty.v0"
)

type downloader interface {
	Get(url string) ([]byte, error)
}

type restyDownloader struct {
	client *resty.Client
}

func newRestyDownloader(unsplashToken string) restyDownloader {
	client := resty.New()
	client.SetHeader("Authorization", "Client-ID "+unsplashToken)
	client.SetHeader("Accept-Version", "v1")

	return restyDownloader{
		client: client,
	}

}

func (r restyDownloader) Get(url string) ([]byte, error) {
	fmt.Println(r.client.Header)
	resp, err := r.client.R().Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}
