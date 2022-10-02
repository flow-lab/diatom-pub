package platform

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v7"
	"time"
)

// GetThrottleExpiration gives when throttle duration will expire.
func GetThrottleExpiration(ctx context.Context, rc *redis.Client, key string) (time.Duration, error) {
	key = fmt.Sprintf("%s_t", key)
	return rc.WithContext(ctx).TTL(key).Result()
}

// Throttle increments the requests count for a specific key and set expiration if it is a new period.
func Throttle(ctx context.Context, rc *redis.Client, key string, expire time.Duration) (int64, error) {
	key = fmt.Sprintf("%s_t", key)

	i, err := rc.WithContext(ctx).Incr(key).Result()
	if err != nil {
		return 0, err
	}

	if i == 1 {
		// the key created, set expire
		ok, err := rc.WithContext(ctx).Expire(key, expire).Result()
		if err != nil {
			// try to remove the key
			if _, e := rc.WithContext(ctx).Del(key).Result(); e != nil {
				return 0, fmt.Errorf("unable to remove key %s, %s and expire failed: %s", key, e.Error(), err.Error())
			}
		} else if !ok {
			return 0, fmt.Errorf("unable to set expiration on key %s", key)
		}
	}

	return i, nil
}
