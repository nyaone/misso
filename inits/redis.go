package inits

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"misso/config"
	"misso/global"
	"time"
)

func Redis() error {
	// Parse connect string
	redisConfig, err := redis.ParseURL(config.Config.System.Redis)
	if err != nil {
		return fmt.Errorf("failed to parse redis connection string: %v", err)
	}

	// Connect to server
	global.Redis = redis.NewClient(redisConfig)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Try connection
	err = global.Redis.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %v", err)
	}

	return nil
}
