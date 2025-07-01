package budget

import (
	"context"
	"fmt"
	"time"

	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/codes"
	"github.com/reyhanmichiels/go-pkg/v2/errors"
)

const (
	getAllBudgetByKey   = "DD:budget:getList:%s"
	deleteBudgetPattern = "DD:budget:*"
)

func (b *budget) upsertCacheList(ctx context.Context, key string, budgets []entity.Budget, ttl time.Duration) error {
	b.log.Debug(ctx, fmt.Sprintf("create list budget cache with key %s and body %v", key, budgets))

	marshalledBudgets, err := b.json.Marshal(budgets)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheMarshal, err.Error())
	}

	err = b.redis.SetEX(ctx, key, string(marshalledBudgets), ttl)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheSetSimpleKey, err.Error())
	}

	b.log.Debug(ctx, fmt.Sprintf("success to create list budget cache with key %s and body %v", key, budgets))

	return nil
}

func (b *budget) getCacheList(ctx context.Context, key string) ([]entity.Budget, error) {
	var budgets []entity.Budget

	b.log.Debug(ctx, fmt.Sprintf("trying to get list budget cache with key: %s", key))

	marshalledBudgets, err := b.redis.Get(ctx, key)
	if err != nil {
		return budgets, errors.NewWithCode(codes.CodeCacheGetSimpleKey, err.Error())
	}

	err = b.json.Unmarshal([]byte(marshalledBudgets), &budgets)
	if err != nil {
		return budgets, errors.NewWithCode(codes.CodeCacheUnmarshal, err.Error())
	}

	b.log.Debug(ctx, fmt.Sprintf("success to get list budget cache with key: %s", key))

	return budgets, nil
}

func (b *budget) deleteCache(ctx context.Context, key string) error {
	b.log.Debug(ctx, "removing budget cache")

	err := b.redis.Del(ctx, key)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheDeleteSimpleKey, err.Error())
	}

	return nil
}

