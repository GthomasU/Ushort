package services

import (
	"Ushort/storage"
	"fmt"
)

type ServiceShortener struct {
	storage      storage.StorageI
	urlGenerator IUrlGenerator
	domain       string
	sslActive    bool
}

func NewServiceShortener(urlGenerator IUrlGenerator, sslActive bool, storage storage.StorageI) ServiceShortener {
	return ServiceShortener{
		storage:      storage,
		urlGenerator: urlGenerator,
		domain:       "localhost",
		sslActive:    false,
	}
}

func (ss *ServiceShortener) SetSslActive(active bool) {
	ss.sslActive = active
}
func (ss *ServiceShortener) CreateShortUrl(originalUrl string) (*string, error) {
	urlId, error := ss.urlGenerator.CreateRandomString(10)
	if error != nil {
		return nil, error
	}
	protocol := "http"
	if ss.sslActive {
		protocol = fmt.Sprintf("%ss", protocol)
	}
	shortedUrl := fmt.Sprintf("%s://%s:3000/r/%s", protocol, ss.domain, *urlId)
	if ss.storage.SaveNewUrl(*urlId, originalUrl) {
		return &shortedUrl, nil
	}
	return nil, fmt.Errorf("cannot created shorted url")
}

func (ss *ServiceShortener) GetOriginalUrl(urlId string) (string, error) {
	if len(urlId) == 0 {
		return "", InvalidUrlId{}
	}
	result, err := ss.storage.GetOriginalUrl(urlId)
	if err != nil {
		switch err.(type) {
		case storage.RecordNotFound:
			return "", UrlNotFound{}
		default:
			return "", err
		}
	}
	return result, nil
}

func (ss *ServiceShortener) RemoveOriginalUrl(urlId string) (bool, error) {
	if len(urlId) == 0 {
		return false, InvalidUrlId{}
	}
	return ss.storage.RemoveOriginalUrl(urlId), nil
}

func (ss *ServiceShortener) UpdateOriginalUrl(urlId, originalUrl string) (bool, error) {
	if len(urlId) == 0 {
		return false, InvalidUrlId{}
	}
	result := ss.storage.UpdateUrl(urlId, originalUrl)
	return result, nil

}
