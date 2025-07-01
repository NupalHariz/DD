package dailyassignment

import (
	"context"
	"fmt"
	"time"

	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/codes"
	"github.com/reyhanmichiels/go-pkg/v2/errors"
)

const (
	getAllDailyAssignmentByKey   = "DD:daily-assignment:getList:%s"
	deleteDailyAssignmentPattern = "DD:daily-assignment:*"
)

func (d *dailyAssignment) upsertCacheList(ctx context.Context, key string, dailyAssignment []entity.DailyAssignment, ttl time.Duration) error {
	d.log.Debug(ctx, fmt.Sprintf("create list daily assignment cache with key %s and body %v", key, dailyAssignment))

	marshalledDailyAssignment, err := d.json.Marshal(dailyAssignment)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheMarshal, err.Error())
	}

	err = d.redis.SetEX(ctx, key, string(marshalledDailyAssignment), ttl)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheSetSimpleKey, err.Error())
	}

	d.log.Debug(ctx, fmt.Sprintf("success to create list daily assignment cache with key %s and body %v", key, dailyAssignment))

	return nil
}

func (d *dailyAssignment) getCacheList(ctx context.Context, key string) ([]entity.DailyAssignment, error) {
	var dailyAssignment []entity.DailyAssignment

	d.log.Debug(ctx, fmt.Sprintf("trying to get list daily assignment cache with key: %s", key))

	marshalledDailyAssignment, err := d.redis.Get(ctx, key)
	if err != nil {
		return dailyAssignment, errors.NewWithCode(codes.CodeCacheGetSimpleKey, err.Error())
	}

	err = d.json.Unmarshal([]byte(marshalledDailyAssignment), &dailyAssignment)
	if err != nil {
		return dailyAssignment, errors.NewWithCode(codes.CodeCacheUnmarshal, err.Error())
	}

	d.log.Debug(ctx, fmt.Sprintf("success to get list daily assignment cache with key: %s", key))

	return dailyAssignment, nil
}

func (d *dailyAssignment) deleteCache(ctx context.Context, key string) error {
	d.log.Debug(ctx, "removing daily assignment cache")

	err := d.redis.Del(ctx, key)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheDeleteSimpleKey, err.Error())
	}

	return nil
}

