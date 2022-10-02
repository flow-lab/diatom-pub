package platform

import (
	"context"
	"fmt"
	"github.com/alicebob/miniredis/v2"
	"strconv"
	"testing"
	"time"
)

func TestThrottle(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()

	port, err := strconv.Atoi(s.Port())
	if err != nil {
		panic(err)
	}

	rc, err := OpenCache(context.Background(), CacheConfig{
		Host: s.Host(),
		Port: int32(port),
	})
	if err != nil {
		panic(err)
	}

	t.Run("should create new key if not exist", func(t *testing.T) {
		key := fmt.Sprintf("%v", time.Now().Unix())
		count, err := Throttle(context.Background(), rc, key, 1*time.Second)
		if err != nil {
			t.Error(err)
		} else if count != 1 {
			t.Error("expected count to be 1 was", count)
		}
	})

	t.Run("should increase 3 times", func(t *testing.T) {
		key := "unittest"
		var count int64
		for i := 0; i < 3; i++ {
			count, err = Throttle(context.Background(), rc, key, 3*time.Second)
		}
		if err != nil {
			t.Error(err)
		} else if count != 3 {
			t.Error("expected count to be 3 was", count)
		}
	})

	t.Run("another unittest", func(t *testing.T) {
		key := "unittest_reset"
		count, err := Throttle(context.Background(), rc, key, 1*time.Second)
		if err != nil {
			t.Fatal(err)
		}

		s.FastForward(1100 * time.Millisecond)

		count, err = Throttle(context.Background(), rc, key, 1*time.Second)
		if err != nil {
			t.Error(err)
		} else if count != 1 {
			t.Error("expected count to be 1 was", count)
		}
	})

	t.Run("should give throttle expiration", func(t *testing.T) {
		key := "unittest_exp"
		_, _ = Throttle(context.Background(), rc, key, 1*time.Second)
		expiration, err := GetThrottleExpiration(context.Background(), rc, key)
		if err != nil {
			t.Fatal(err)
		}

		if expiration != 1*time.Second {
			t.Error("should be 1 sec")
		}
	})
}
