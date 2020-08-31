package repo

import "github.com/garyburd/redigo/redis"

type cache struct {
	pool *redis.Pool
}

func InitCache(pool *redis.Pool) UserRepository {
	return &cache{
		pool,
	}
}

func (c *cache) findUserByID(id string) {
	conn := c.pool.Get()
	conn.Do("GET", id)
}
