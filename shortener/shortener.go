package shortener

import (
	"Ushort/redis"
	"fmt"
)

func CreateShortUrl(originalUrl string) (*string, error) {
	urlId := createRandomString(10)
	shortedUrl := fmt.Sprintf("https://localhost:3000/r/%s", urlId)
	if redis.SaveNewUrl(urlId, originalUrl) {
		return &shortedUrl, nil
	}
	return nil, fmt.Errorf("cannot created shorted url")
}

func GetOriginalUrl(urlId string) (string, error) {
	result, err := redis.GetOriginalUrl(urlId)
	if err != nil {
		if _, ok := err.(redis.RecordNotFound); ok {
			return "", UrlNotFound{}
		}
	}
	return result, nil
}

func RemoveOriginalUrl(urlId string) bool {
	return redis.RemoveOriginalUrl(urlId)
}

func UpdateOriginalUrl(urlId, originalUrl string) bool {
	return redis.UpdateUrl(urlId, originalUrl)

}
