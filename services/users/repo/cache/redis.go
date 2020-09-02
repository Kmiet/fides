package cache

import (
	"github.com/Kmiet/fides/internal/storage"
	"github.com/Kmiet/fides/internal/storage/redis"
	redigo "github.com/garyburd/redigo/redis"
)

func composeKey(parts ...interface{}) string {
	var key string = storage.DATA_TYPES.USERS
	for _, part := range parts {
		key += storage.CACHE_KEY_SEPARATOR + part.(string)
	}
	return key
}

// Command Wrappers
func (c *cache) get(id string) ([]byte, error) {
	return redigo.Bytes(
		c.pool.Get().Do(
			redis.COMMANDS.GET,
			composeKey(id),
		),
	)
}

func (c *cache) set(id string, value string) {
	c.pool.Get().Do(redis.COMMANDS.SET, composeKey(id))
}
