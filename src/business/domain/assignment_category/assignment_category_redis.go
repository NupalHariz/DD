package assignmentcategory

import (
	"context"
	"fmt"
	"time"

	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/codes"
	"github.com/reyhanmichiels/go-pkg/v2/errors"
)

const (
	getAllAssignmentCategoryByKey   = "DD:assignment-category:getList:%s"
	deleteAssignmentCategoryPattern = "DD:assignment-category:*"
)

func (a *assignmentCategory) upsertCacheList(ctx context.Context, key string, assignmentCategory []entity.AssignmentCategory, ttl time.Duration) error {
	a.log.Debug(ctx, fmt.Sprintf("create list assignment category cache with key %s and body %v", key, assignmentCategory))

	marshalledAssignmentCategory, err := a.json.Marshal(assignmentCategory)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheMarshal, err.Error())
	}

	err = a.redis.SetEX(ctx, key, string(marshalledAssignmentCategory), ttl)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheSetSimpleKey, err.Error())
	}

	a.log.Debug(ctx, fmt.Sprintf("success to create list assignment category cache with key %s and body %v", key, assignmentCategory))

	return nil
}

func (a *assignmentCategory) getCacheList(ctx context.Context, key string) ([]entity.AssignmentCategory, error) {
	var assignmentCategory []entity.AssignmentCategory

	a.log.Debug(ctx, fmt.Sprintf("trying to get list assignment category cache with key: %s", key))

	marshalledAssignmentCategory, err := a.redis.Get(ctx, key)
	if err != nil {
		return assignmentCategory, errors.NewWithCode(codes.CodeCacheGetSimpleKey, err.Error())
	}

	err = a.json.Unmarshal([]byte(marshalledAssignmentCategory), &assignmentCategory)
	if err != nil {
		return assignmentCategory, errors.NewWithCode(codes.CodeCacheUnmarshal, err.Error())
	}

	a.log.Debug(ctx, fmt.Sprintf("success to get list assignment category cache with key: %s", key))

	return assignmentCategory, nil
}

func (a *assignmentCategory) deleteCache(ctx context.Context, key string) error {
	a.log.Debug(ctx, "removing assignment category cache")

	err := a.redis.Del(ctx, key)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheDeleteSimpleKey, err.Error())
	}

	return nil
}
