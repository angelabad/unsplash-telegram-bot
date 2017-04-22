package main

import resty "gopkg.in/resty.v0"

func (p *Photo) getSmallImage() ([]byte, error) {
	resty.SetHeader("Authorization", "Client-ID "+unsplashID)
	resty.SetHeader("Accept-Version", "v1")

	resp, err := resty.R().Get(p.Urls.Small)
	if err != nil {
		return nil, err
	}

	return resp.Body(), nil
}

// Photo struct
type Photo struct {
	ID          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Color       string `json:"color"`
	Downloads   int    `json:"downloads"`
	Likes       int    `json:"likes"`
	LikedByUser bool   `json:"liked_by_user"`
	Exif        struct {
		Make         string `json:"make"`
		Model        string `json:"model"`
		ExposureTime string `json:"exposure_time"`
		Aperture     string `json:"aperture"`
		FocalLength  string `json:"focal_length"`
		Iso          int    `json:"iso"`
	} `json:"exif"`
	CurrentUserCollections []interface{} `json:"current_user_collections"`
	Urls                   struct {
		Raw     string `json:"raw"`
		Full    string `json:"full"`
		Regular string `json:"regular"`
		Small   string `json:"small"`
		Thumb   string `json:"thumb"`
	} `json:"urls"`
	Categories []interface{} `json:"categories"`
	Links      struct {
		Self             string `json:"self"`
		HTML             string `json:"html"`
		Download         string `json:"download"`
		DownloadLocation string `json:"download_location"`
	} `json:"links"`
	User struct {
		ID               string `json:"id"`
		UpdatedAt        string `json:"updated_at"`
		Username         string `json:"username"`
		Name             string `json:"name"`
		FirstName        string `json:"first_name"`
		LastName         string `json:"last_name"`
		PortfolioURL     string `json:"portfolio_url"`
		Bio              string `json:"bio"`
		Location         string `json:"location"`
		TotalLikes       int    `json:"total_likes"`
		TotalPhotos      int    `json:"total_photos"`
		TotalCollections int    `json:"total_collections"`
		ProfileImage     struct {
			Small  string `json:"small"`
			Medium string `json:"medium"`
			Large  string `json:"large"`
		} `json:"profile_image"`
		Links struct {
			Self      string `json:"self"`
			HTML      string `json:"html"`
			Photos    string `json:"photos"`
			Likes     string `json:"likes"`
			Portfolio string `json:"portfolio"`
			Following string `json:"following"`
			Followers string `json:"followers"`
		} `json:"links"`
	} `json:"user"`
}
