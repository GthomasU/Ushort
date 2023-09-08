package redis

import (
	"context"

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

func GetOriginalUrl(urlId string) string {
	result := client.HGet(redis_context, HASH_URL_ORIGINAL_TO_SHORTED, urlId)
	originalUrl := result.Val()
	return originalUrl
}
