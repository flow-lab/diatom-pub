package platform

import (
	"context"
	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestOpenCache(t *testing.T) {
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

	t.Run("should connect", func(t *testing.T) {
		result, err := rc.Ping().Result()
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "PONG", result)
	})
}
