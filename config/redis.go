package config

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func ConnectRedis() {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "redis:6379" // fallback mặc định nếu không có biến môi trường
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // thêm nếu Redis có password
		DB:       0,
	})

	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ Connected to Redis")
}
