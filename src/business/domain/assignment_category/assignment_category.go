package assignmentcategory

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
	Create(ctx context.Context, param entity.AssignmentCategoryInputParam) error
	GetAll(ctx context.Context, param entity.AssignmentCategoryParam) ([]entity.AssignmentCategory, error)
}

type assignmentCategory struct {
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
	return &assignmentCategory{
		db:    param.Db,
		log:   param.Log,
		json:  param.Json,
		redis: param.Redis,
	}
}

func (a *assignmentCategory) Create(ctx context.Context, param entity.AssignmentCategoryInputParam) error {
	err := a.createSQL(ctx, param)
	if err != nil {
		return err
	}

	err = a.deleteCache(ctx, deleteAssignmentCategoryPattern)
	if err != nil {
		a.log.Error(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
	}

	return nil
}

func (a *assignmentCategory) GetAll(ctx context.Context, param entity.AssignmentCategoryParam) ([]entity.AssignmentCategory, error) {
	var categories []entity.AssignmentCategory

	marshalledParam, err := a.json.Marshal(param)
	if err != nil {
		return categories, err
	}

	if !param.BypassCache {
		categories, err = a.getCacheList(ctx, fmt.Sprintf(getAllAssignmentCategoryByKey, string(marshalledParam)))
		switch {
		case errors.Is(err, redis.Nil):
			a.log.Warn(ctx, fmt.Sprintf(entity.ErrorRedisNil, err.Error()))
		case err != nil:
			a.log.Warn(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
		default:
			return categories, nil
		}
	}

	categories, err = a.getAllSQL(ctx, param)
	if err != nil {
		return []entity.AssignmentCategory{}, err
	}

	err = a.upsertCacheList(ctx, fmt.Sprintf(getAllAssignmentCategoryByKey, string(marshalledParam)), categories, a.redis.GetDefaultTTL(ctx))
	if err != nil {
		a.log.Error(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
	}

	return categories, nil
}
