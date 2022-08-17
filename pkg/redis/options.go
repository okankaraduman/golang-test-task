package redis

import "time"

// Option -.
type Option func(*Redis)

// MaxPoolSize -.
func MaxPoolSize(size int) Option {
	return func(c *Redis) {
		c.poolSize = size
	}
}

// ConnAttempts -.
func ConnAttempts(attempts int) Option {
	return func(c *Redis) {
		c.maxRetries = attempts
	}
}

// PoolTimeout -.
func PoolTimeout(timeout time.Duration) Option {
	return func(c *Redis) {
		c.poolTimeout= timeout
	}
}
