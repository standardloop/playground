package redis

import (
	"api-std/config"
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

func RedisPoolInit() {

	redisDBNum, _ := strconv.Atoi(config.Env.RedisDBNum)
	redisPool := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", config.Env.RedisHost, config.Env.RedisPort),
		Password: config.Env.RedisPass,
		DB:       redisDBNum,

		PoolSize:     10,               // Max number of connections
		MinIdleConns: 5,                // Maintain at least 5 idle connections
		PoolTimeout:  30 * time.Second, // Max wait time for a connection from the pool
	})

	ctx := context.Background()
	_, err := redisPool.Ping(ctx).Result()
	if err != nil {
		slog.Error("Could not connect to Redis: %v", err)
		return
	}
	slog.Info("RedisPool has been succesfully initalized")
	RedisPool = redisPool
}

func RedisPoolDestroy() {
	RedisPool.Close()
}

func RedisPoolPing() error {
	ctx := context.Background()
	_, err := RedisPool.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

var RedisPool *redis.Client = nil
