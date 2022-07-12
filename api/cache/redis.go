package cache

import (
	"encoding/json"
	"time"
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/VladRomanciuc/Go-classes/api/models"
)

type RedisCache struct {
	host    string
	db      int
	pswd    string
	expires time.Duration
}

func NewRedisCache(host string, pswd string, db int, exp time.Duration) models.PostCache {
	return &RedisCache{
		host:    host,
		db:      db,
		pswd:	pswd,
		expires: exp,
	}
}

func (cache *RedisCache) getClient() *redis.Client {
	
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: cache.pswd,
		DB:       cache.db,
	})
}

func (cache *RedisCache) Set(key string, post *models.Post) {
	client := cache.getClient()

	// serialize Post object to JSON
	json, err := json.Marshal(post)
	if err != nil {
		panic(err)
	}
	c := context.Background()
	client.Set(c, key, json, cache.expires*time.Second)
}

func (cache *RedisCache) Get(key string) *models.Post {
	client := cache.getClient()
	
	c := context.Background()
	val, err := client.Get(c, key).Result()
	if err != nil {
		return nil
	}

	post := models.Post{}
	err = json.Unmarshal([]byte(val), &post)
	if err != nil {
		panic(err)
	}

	return &post
}

func (cache *RedisCache) Del(key string) *models.Post {
	client := cache.getClient()
	
	c := context.Background()
	del, err := client.Del(c, key).Result()
	if err != nil {
		return nil
	}

	post := models.Post{}
	err = json.Unmarshal([]byte(string(del)), &post)
	if err != nil {
		panic(err)
	}

	return &post
}
