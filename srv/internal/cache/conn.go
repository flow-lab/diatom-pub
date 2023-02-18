package cache

import (
	"fmt"
	"github.com/flow-lab/diatom-pub/internal/helper"
	"github.com/go-redis/redis/v7"
)

// NewClient returns a new redis client.
func NewClient() (*redis.Client, error) {
	var (
		redisHost = helper.MustGetEnv("REDIS_HOST")
		redisPort = helper.MustGetEnv("REDIS_PORT")
	)

	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
	})

	return client, nil
}
