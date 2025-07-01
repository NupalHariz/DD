package category

import (
	"context"
	"errors"
	"fmt"

	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/log"
	"github.com/reyhanmichiels/go-pkg/v2/parser"
	"github.com/reyhanmichiels/go-pkg/v2/redis"
	"github.com/reyhanmichiels/go-pkg/v2/sql"
)

type Interface interface {
	Create(ctx context.Context, param entity.CategoryInputParam) (entity.Category, error)
	GetAll(ctx context.Context, param entity.CategoryParam) ([]entity.Category, error)
}

type category struct {
	db    sql.Interface
	log   log.Interface
	json  parser.JSONInterface
	redis redis.Interface
}

type InitParam struct {
	Db    sql.Interface
	Log   log.Interface
	Json  parser.JSONInterface
	Redis redis.Interface
}

func Init(param InitParam) Interface {
	return &category{
		db:    param.Db,
		log:   param.Log,
		json:  param.Json,
		redis: param.Redis,
	}
}

func (c *category) Create(ctx context.Context, param entity.CategoryInputParam) (entity.Category, error) {
	category, err := c.createSQL(ctx, param)
	if err != nil {
		return category, err
	}

	err = c.deleteCache(ctx, deleteCategoryPattern)
	if err != nil {
		c.log.Error(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
	}

	return category, nil
}

func (c *category) GetAll(ctx context.Context, param entity.CategoryParam) ([]entity.Category, error) {
	var categories []entity.Category

	marshalledParam, err := c.json.Marshal(param)
	if err != nil {
		return categories, err
	}

	if !param.BypassCache {
		categories, err = c.getCacheList(ctx, fmt.Sprintf(getAllCategoryByKey, marshalledParam))
		switch {
		case errors.Is(err, redis.Nil):
			c.log.Warn(ctx, fmt.Sprintf(entity.ErrorRedisNil, err.Error()))
		case err != nil:
			c.log.Warn(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
		default:
			return categories, nil
		}
	}

	categories, err = c.getAllSQL(ctx, param)
	if err != nil {
		return categories, err
	}

	err = c.upsertCacheList(ctx, fmt.Sprintf(getAllCategoryByKey, marshalledParam), categories, c.redis.GetDefaultTTL(ctx))
	if err != nil {
		c.log.Warn(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
	}

	return categories, nil
}
