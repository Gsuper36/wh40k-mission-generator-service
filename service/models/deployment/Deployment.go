package deployment

import (
	"errors"
	"net/url"
)

type Deployment struct {
	imageUrl string
}

func (d Deployment) ImageUrl() string {
	return d.imageUrl
}

func (d *Deployment) SetImageUrl(img string) error {
	if !d.isValidUrl(img) {
		return errors.New("invalid url")
	}

	d.imageUrl = img

	return nil
}

func (d Deployment) isValidUrl(path string) bool {
	_, err := url.ParseRequestURI(path)

	if err != nil {
		return false
	}

	u, err := url.Parse(path)

	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func NewDeployment(img string) (*Deployment, error) {
	d := &Deployment{}

	err := d.SetImageUrl(img)

	if err != nil {
		return &Deployment{}, err
	}

	return d, nil
}