package cache

import (
	"fmt"
	utils "github.com/flow-lab/utils"
	"github.com/go-redis/redis/v7"
)

// NewClient returns a new redis client.
func NewClient() (*redis.Client, error) {
	var (
		redisHost = utils.MustGetEnv("REDIS_HOST")
		redisPort = utils.MustGetEnv("REDIS_PORT")
	)

	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
	})

	return client, nil
}
