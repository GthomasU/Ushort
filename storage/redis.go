package storage

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client        *redis.Client
	redis_context context.Context
}

const HASH_URL_ORIGINAL_TO_SHORTED string = "original_url_to_shorted_url"

func NewRedisClient() *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &RedisClient{
		client:        client,
		redis_context: context.Background(),
	}

}
func (rc RedisClient) UpdateUrl(urlId, originalUrl string) bool {
	_ = rc.client.HSet(rc.redis_context, HASH_URL_ORIGINAL_TO_SHORTED, urlId, originalUrl)
	result := rc.client.HGet(rc.redis_context, HASH_URL_ORIGINAL_TO_SHORTED, urlId)
	urlUpdated := result.Val()
	return urlUpdated == originalUrl
}

func (rc RedisClient) SaveNewUrl(urlId, originalUrl string) bool {
	fieldsAdded := rc.client.HSet(rc.redis_context, HASH_URL_ORIGINAL_TO_SHORTED, urlId, originalUrl)
	return fieldsAdded.Val() == 1
}

func (rc RedisClient) GetOriginalUrl(urlId string) (string, error) {
	result := rc.client.HGet(rc.redis_context, HASH_URL_ORIGINAL_TO_SHORTED, urlId)
	originalUrl := result.Val()
	if len(originalUrl) > 0 {
		return originalUrl, nil
	} else {
		return "", RecordNotFound{}
	}

}

func (rc RedisClient) RemoveOriginalUrl(urlId string) bool {
	result := rc.client.HDel(rc.redis_context, HASH_URL_ORIGINAL_TO_SHORTED, urlId)
	return result.Val() == 1
}
