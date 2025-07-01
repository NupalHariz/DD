package assignment

import (
	"context"
	"fmt"
	"time"

	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/codes"
	"github.com/reyhanmichiels/go-pkg/v2/errors"
)

const (
	getAllAssignmentByKey      = "DD:assignment:gets:%s"
	deleteAssignmentKeyPattern = "DD:assignment:*"
)

func (a *assignment) upsertCacheList(ctx context.Context, key string, assignments []entity.Assignment, ttl time.Duration) error {
	a.log.Debug(ctx, fmt.Sprintf("create cache list with key: %s", key))

	marshalledAssignment, err := a.json.Marshal(assignments)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheMarshal, err.Error())
	}

	a.log.Debug(ctx, fmt.Sprintf("SETTING REDIS WITH KEY %s, BODY %v, and TT %s", key, assignments, ttl))
	err = a.redis.SetEX(ctx, key, string(marshalledAssignment), ttl)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheSetSimpleKey, err.Error())
	}

	a.log.Debug(ctx, fmt.Sprintf("success to set cache list with key %s and body %v", key, assignments))

	return nil
}

func (a *assignment) getCacheList(ctx context.Context, key string) ([]entity.Assignment, error) {
	var assignments []entity.Assignment

	a.log.Debug(ctx, fmt.Sprintf("trying to get assignment list from cache with key: %s", key))

	marshalAssignments, err := a.redis.Get(ctx, key)
	if err != nil {
		return assignments, errors.NewWithCode(codes.CodeCacheGetSimpleKey, err.Error())
	}

	err = a.json.Unmarshal([]byte(marshalAssignments), &assignments)
	if err != nil {
		return assignments, errors.NewWithCode(codes.CodeCacheUnmarshal, err.Error())
	}

	a.log.Debug(ctx, fmt.Sprintf("success to get assignment list from cache with key: %s", key))

	return assignments, nil
}

func (a *assignment) deleteCache(ctx context.Context, key string) error {
	a.log.Debug(ctx, "remove assignment from cache")

	err := a.redis.Del(ctx, key)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheDeleteSimpleKey, err.Error())
	}

	a.log.Debug(ctx, "success to remove assignment from cache")
	return nil
}
