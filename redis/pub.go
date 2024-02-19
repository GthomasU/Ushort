package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client
var redis_context context.Context

const (
	HASH_URL_ORIGINAL_TO_SHORTED = "original_url_to_shorted_url"
)

func init() {
	redis_context = context.Background()
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func SaveNewUrl(urlId, originalUrl string) bool {
	fieldsAdded := client.HSet(redis_context, HASH_URL_ORIGINAL_TO_SHORTED, urlId, originalUrl)
	return fieldsAdded.Val() == 1
}

func GetOriginalUrl(urlId string) (string, error) {
	result := client.HGet(redis_context, HASH_URL_ORIGINAL_TO_SHORTED, urlId)
	originalUrl := result.Val()
	fmt.Print(originalUrl)
	if len(originalUrl) > 0 {
		return originalUrl, nil
	} else {
		return "", RecordNotFound{}
	}

}

func RemoveOriginalUrl(urlId string) bool {
	result := client.HDel(redis_context, HASH_URL_ORIGINAL_TO_SHORTED, urlId)
	return result.Val() == 1
}
