package repository

import (
	"context"
	"fmt"

	"github.com/okankaraduman/golang-test-task/internal/entity"
	"github.com/okankaraduman/golang-test-task/pkg/redis"
)

const _defaultEntityCap = 64

// KeyValueRepo -.
type KeyValueRepo struct {
	*redis.Redis
}

func (r *KeyValueRepo) GetMessages(ctx context.Context) ([]entity.Pair, error) {
	pairs := make([]entity.Pair, 0, _defaultEntityCap)
	stringpairs, err := r.Redis.Rdcli.HGetAll(ctx, "keys").Result()

	if err != nil {
		return nil, fmt.Errorf(" KeyValueRepo - GetMessages - r.Rdcli.HGetAll: %w", err)
	}
	for key, value := range stringpairs {
		pair := entity.Pair{
			Sender:  key,
			Message: value,
		}
		pairs = append(pairs, pair)
	}

	return pairs, nil
}
