package redis

import (
	"os"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
)

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
			c, err := redis.DialURL(uri)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do(COMMANDS.PING)
			return err
		},
	}
}

func TestConnection(pool *redis.Pool) {
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do(COMMANDS.PING)
	if err != nil {
		panic(err)
	}
}

func Disconnect(pool *redis.Pool) {
	pool.Close()
}
