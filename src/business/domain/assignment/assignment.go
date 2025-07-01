package assignment

import (
	"context"
	"fmt"

	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/errors"
	"github.com/reyhanmichiels/go-pkg/v2/log"
	"github.com/reyhanmichiels/go-pkg/v2/parser"
	"github.com/reyhanmichiels/go-pkg/v2/redis"
	"github.com/reyhanmichiels/go-pkg/v2/sql"
)

type Interface interface {
	Create(ctx context.Context, param entity.AssignmentInputParam) error
	Update(ctx context.Context, updateParam entity.AssignmentUpdateParam, assignmentParam entity.AssignmentParam) error
	GetAll(ctx context.Context, param entity.AssignmentParam) ([]entity.Assignment, error)
}

type assignment struct {
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
	return &assignment{
		db:    param.Db,
		log:   param.Log,
		json:  param.Json,
		redis: param.Redis,
	}
}

func (a *assignment) Create(ctx context.Context, param entity.AssignmentInputParam) error {
	err := a.createSQL(ctx, param)
	if err != nil {
		return err
	}

	err = a.deleteCache(ctx, deleteAssignmentKeyPattern)
	if err != nil {
		a.log.Warn(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
	}

	return nil
}

func (a *assignment) Update(ctx context.Context, updateParam entity.AssignmentUpdateParam, assignmentParam entity.AssignmentParam) error {
	err := a.updateSQL(ctx, updateParam, assignmentParam)
	if err != nil {
		return err
	}

	err = a.deleteCache(ctx, deleteAssignmentKeyPattern)
	if err != nil {
		a.log.Warn(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
	}

	return nil
}

func (a *assignment) GetAll(ctx context.Context, param entity.AssignmentParam) ([]entity.Assignment, error) {
	var assignments []entity.Assignment

	marshalledParam, err := a.json.Marshal(param)
	if err != nil {
		return assignments, err
	}

	if !param.BypassCache {
		assignments, err = a.getCacheList(ctx, fmt.Sprintf(getAllAssignmentByKey, string(marshalledParam)))
		switch {
		case errors.Is(err, redis.Nil):
			a.log.Warn(ctx, fmt.Sprintf(entity.ErrorRedisNil, err.Error()))
		case err != nil:
			a.log.Warn(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
		default:
			return assignments, nil
		}
	}

	assignments, err = a.getAllSQL(ctx, param)
	if err != nil {
		return assignments, err
	}

	err = a.upsertCacheList(ctx, fmt.Sprintf(getAllAssignmentByKey, string(marshalledParam)), assignments, a.redis.GetDefaultTTL(ctx))
	if err != nil {
		a.log.Error(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
	}

	return assignments, nil
}
