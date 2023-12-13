package db

import (
	"time"

	"github.com/go-redis/redis"
)

const (
	NO_TIMEOUT = 0
)

type RedisClient interface {
	Ping() *redis.StatusCmd
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}

type connection struct {
	cli RedisClient
}

func New(cli RedisClient) (*connection, error) {
	return &connection{cli}, nil
}

func (c *connection) Ping() error {
	return c.cli.Ping().Err()
}

// Insert implements router.DBConn.
func (c *connection) Insert(key string, value []byte) error {
	return c.cli.Set(key, value, NO_TIMEOUT).Err()
}
