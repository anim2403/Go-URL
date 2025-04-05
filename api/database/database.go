package database

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)
// Ctx is the context for Redis operations
var Ctx = context.Background()
// CreateClient creates a new Redis client
func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		DB:       dbNo,
	})
	return rdb
}