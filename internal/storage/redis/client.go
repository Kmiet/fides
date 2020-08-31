package redis

import (
	"os"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
)

const _PING = "PING"

func InitClient(uri string) *redis.Pool {
	ps, ok := os.LookupEnv("REDIS_POOL_SIZE")
	if !ok {
		ps = "2"
	}
	poolSize, _ := strconv.Atoi(ps)

	return &redis.Pool{
		MaxIdle:     poolSize,
		IdleTimeout: 300 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", uri)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do(_PING)
			return err
		},
	}
}

func TestConnection(pool *redis.Pool) {
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do(_PING)
	if err != nil {
		panic(err)
	}
}

func Disconnect(pool *redis.Pool) {
	pool.Close()
}
