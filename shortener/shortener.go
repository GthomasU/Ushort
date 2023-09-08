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

func GetOriginalUrl(urlId string) string {
	return redis.GetOriginalUrl(urlId)
}
