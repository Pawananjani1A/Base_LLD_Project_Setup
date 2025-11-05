package redis

import (
	"base_app/config"
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// Redis wraps a connected Redis client.
type Redis struct {
	Client *redis.Client
}

var redisInstance *Redis

// InitRedis initializes a singleton Redis client using provided config.
// Safe to call multiple times; the first successful call wins.
func InitRedis() (*Redis, error) {

	if redisInstance != nil {
		return redisInstance, nil
	}

	redisConf := config.GetRedisConf()

	dialTimeout := redisConf.DialTimeout
	readTimeout := redisConf.ReadTimeout
	writeTimeout := redisConf.WriteTimeout

	if dialTimeout <= 0 {
		dialTimeout = 5
	}
	if readTimeout <= 0 {
		readTimeout = 3
	}
	if writeTimeout <= 0 {
		writeTimeout = 3
	}

	client := redis.NewClient(&redis.Options{
		Addr:         redisConf.Addr,
		Password:     redisConf.Password,
		DB:           redisConf.Db,
		DialTimeout:  time.Duration(dialTimeout) * time.Second,
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
	})

	// Ping to verify connectivity
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	redisInstance = &Redis{Client: client}
	return redisInstance, nil
}

// GetRedis returns the initialized Redis singleton, or nil if not initialized.
func GetRedis() *Redis {
	return redisInstance
}

// Close terminates the underlying Redis client.
func (r *Redis) Close() error {
	if r == nil || r.Client == nil {
		return nil
	}
	return r.Client.Close()
}
