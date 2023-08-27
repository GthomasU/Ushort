package shortener

import (
	"fmt"
)

func CreateShortUrl(originalUrl string) string {
	randomString := createRandomString(10)
	shortedUrl := fmt.Sprintf("https://localhost:3000/%s", randomString)
	return shortedUrl
}
