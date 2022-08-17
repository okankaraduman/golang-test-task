// Package postgres implements postgres connection.
package redis

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	//I should definitely check these out and replace with my own-logic
)

const (
	_defaultMaxPoolSize  = 1
	_defaultConnAttempts = 10
	_defaultConnTimeout  = time.Second
)

// Redis -.
type Redis struct {
	maxRetries int

	poolSize    int
	poolTimeout time.Duration
	ıdleTimeout time.Duration

	Rdcli *redis.Client
}

// New -.
func New(url string, opts ...Option) (*Redis, error) {
	re := &Redis{
		poolSize:    _defaultMaxPoolSize,
		maxRetries:  _defaultConnAttempts,
		poolTimeout: _defaultConnTimeout,
	}

	// Custom options
	for _, opt := range opts {
		opt(re)
	}

	redisConfig := redis.Options{
		Addr:     url,
		Password: os.Getenv("REDIS_PASSWORD"),
	}

	for re.maxRetries > 0 {

		re.Rdcli = redis.NewClient(&redisConfig)

		if err := re.Rdcli.Ping(context.Background()).Err(); err == nil {
			break
		}
		log.Printf("Redis is trying to connect, attempts left: %d", re.maxRetries)
		time.Sleep(re.poolTimeout)

		re.maxRetries--
	}
	return re, nil
}

func (re *Redis) Close() {
	if re.Rdcli.PoolStats().TotalConns != 0 {
		re.Rdcli.Close()
	}

}
