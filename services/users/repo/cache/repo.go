package cache

import (
	redigo "github.com/garyburd/redigo/redis"
)

type Repository interface {
	get(id string) ([]byte, error)
	set(id string, value string)

	GetUserByID(id string) ([]byte, error)
	SetUserByID(id string, body string)
}

type cache struct {
	pool *redigo.Pool
}

func InitRepository(pool *redigo.Pool) Repository {
	return &cache{
		pool,
	}
}

func (c *cache) GetUserByID(id string) ([]byte, error) {
	return c.get(id)
}

func (c *cache) SetUserByID(id string, body string) {
	c.set(id, body)
	return
}
