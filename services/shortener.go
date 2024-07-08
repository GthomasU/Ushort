package services

import (
	"Ushort/storage"
	"fmt"
)

type ServiceShortener struct {
	storage storage.StorageI
}

func NewServiceShortener() ServiceShortener {
	redisClient := storage.NewRedisClient()
	return ServiceShortener{
		storage: redisClient,
	}
}

func (ss ServiceShortener) CreateShortUrl(originalUrl string) (*string, error) {
	urlId := createRandomString(10)
	shortedUrl := fmt.Sprintf("https://localhost:3000/r/%s", urlId)
	if ss.storage.SaveNewUrl(urlId, originalUrl) {
		return &shortedUrl, nil
	}
	return nil, fmt.Errorf("cannot created shorted url")
}

func (ss ServiceShortener) GetOriginalUrl(urlId string) (string, error) {
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

func (ss ServiceShortener) RemoveOriginalUrl(urlId string) bool {
	return ss.storage.RemoveOriginalUrl(urlId)
}

func (ss ServiceShortener) UpdateOriginalUrl(urlId, originalUrl string) bool {
	return ss.storage.UpdateUrl(urlId, originalUrl)
}
