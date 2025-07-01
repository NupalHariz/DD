package category

import (
	"context"
	"fmt"
	"time"

	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/codes"
	"github.com/reyhanmichiels/go-pkg/v2/errors"
)

const (
	getCategoryByKey     = "DD:category:get:%s"
	getAllCategoryByKey  = "DD:category:getList:%s"
	deleteCategoryPattern = "DD:category:*"
)

func (c *category) upsertCache(ctx context.Context, key string, category entity.Category, ttl time.Duration) error {
	c.log.Debug(ctx, fmt.Sprintf("create category cache with key %s and body %v", key, category))

	marshalledCategory, err := c.json.Marshal(&category)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheMarshal, err.Error())
	}

	err = c.redis.SetEX(ctx, key, string(marshalledCategory), ttl)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheSetSimpleKey, err.Error())
	}

	c.log.Debug(ctx, fmt.Sprintf("success to create category cachse with key %s and body %v", key, category))

	return nil
}

func (c *category) upsertCacheList(ctx context.Context, key string, categories []entity.Category, ttl time.Duration) error {
	c.log.Debug(ctx, fmt.Sprintf("create list category cache with key %s and body %v", key, categories))

	marshalledCategories, err := c.json.Marshal(&categories)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheMarshal, err.Error())
	}

	err = c.redis.SetEX(ctx, key, string(marshalledCategories), ttl)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheSetSimpleKey, err.Error())
	}

	c.log.Debug(ctx, fmt.Sprintf("success to create list category cachse with key %s and body %v", key, categories))

	return nil
}

func (c *category) getCache(ctx context.Context, key string) (entity.Category, error) {
	var category entity.Category

	c.log.Debug(ctx, fmt.Sprintf("trying to get category cache with key: %s", key))

	marshalledCategory, err := c.redis.Get(ctx, key)
	if err != nil {
		return category, errors.NewWithCode(codes.CodeCacheGetSimpleKey, err.Error())
	}

	err = c.json.Unmarshal([]byte(marshalledCategory), &category)
	if err != nil {
		return category, errors.NewWithCode(codes.CodeCacheUnmarshal, err.Error())
	}

	c.log.Debug(ctx, fmt.Sprintf("success to get category cache with key: %s", key))

	return category, nil
}

func (c *category) getCacheList(ctx context.Context, key string) ([]entity.Category, error) {
	var categories []entity.Category

	c.log.Debug(ctx, fmt.Sprintf("trying to get list category cache with key: %s", key))

	marshalledCategories, err := c.redis.Get(ctx, key)
	if err != nil {
		return categories, errors.NewWithCode(codes.CodeCacheGetSimpleKey, err.Error())
	}

	err = c.json.Unmarshal([]byte(marshalledCategories), &categories)
	if err != nil {
		return categories, errors.NewWithCode(codes.CodeCacheUnmarshal, err.Error())
	}

	c.log.Debug(ctx, fmt.Sprintf("success to get list category cache with key: %s", key))

	return categories, nil
}

func (c *category) deleteCache(ctx context.Context, key string) error {
	c.log.Debug(ctx, "removing category cache")

	err := c.redis.Del(ctx, key)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheDeleteSimpleKey, err.Error())
	}

	return nil
}
