package platform

//noinspection ALL
import (
	redisg "cloud.google.com/go/redis/apiv1"
	"context"
	"fmt"
	"github.com/go-redis/redis/v7"
	redispb "google.golang.org/genproto/googleapis/cloud/redis/v1"
)

// CacheConfig Config is the required properties to use the cache.
type CacheConfig struct {
	Name string
	Host string
	Port int32
}

// OpenCache knows how to open Redis connection based on the instance name.
func OpenCache(ctx context.Context, cfg CacheConfig) (*redis.Client, error) {
	var host string
	var port int32
	if cfg.Name != "" {
		c, err := redisg.NewCloudRedisClient(ctx)
		if err != nil {
			return nil, err
		}

		req := redispb.GetInstanceRequest{
			Name: cfg.Name,
		}

		ri, err := c.GetInstance(ctx, &req)
		if err != nil {
			return nil, err
		}

		host = ri.GetHost()
		port = ri.GetPort()
	} else {
		host = cfg.Host
		port = cfg.Port
	}

	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", host, port),
		DB:   0, // use default DB
	})

	return client, nil
}

// CacheStatusCheck returns nil if it can successfully talk to the cache. It
// returns a non-nil error otherwise.
func CacheStatusCheck(ctx context.Context, cc *redis.Client) error {
	if cc == nil {
		return fmt.Errorf("cache is nil")
	}
	_, err := cc.WithContext(ctx).Ping().Result()
	return err
}
