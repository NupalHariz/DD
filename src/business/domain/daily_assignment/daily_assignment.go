package dailyassignment

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
	Create(ctx context.Context, param entity.DailyAssignmentInputParam) error
	Update(ctx context.Context, updateParam entity.DailyAssignmentUpdateParam, dailyAssignmentParam entity.DailyAssignmentParam) error
	GetAll(ctx context.Context, param entity.DailyAssignmentParam) ([]entity.DailyAssignment, error)
	UpdateDailyAssignmentToFalse(ctx context.Context) error
}

type dailyAssignment struct {
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
	return &dailyAssignment{
		db:    param.Db,
		log:   param.Log,
		json:  param.Json,
		redis: param.Redis,
	}
}

func (d *dailyAssignment) Create(ctx context.Context, param entity.DailyAssignmentInputParam) error {
	err := d.createSQL(ctx, param)
	if err != nil {
		return err
	}

	err = d.deleteCache(ctx, deleteDailyAssignmentPattern)
	if err != nil {
		d.log.Error(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
	}

	return nil
}

func (d *dailyAssignment) Update(
	ctx context.Context,
	updateParam entity.DailyAssignmentUpdateParam,
	dailyAssignmentParam entity.DailyAssignmentParam,
) error {
	err := d.updateSQL(ctx, updateParam, dailyAssignmentParam)
	if err != nil {
		return err
	}

	err = d.deleteCache(ctx, deleteDailyAssignmentPattern)
	if err != nil {
		d.log.Error(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
	}

	return nil
}

func (d *dailyAssignment) GetAll(ctx context.Context, param entity.DailyAssignmentParam) ([]entity.DailyAssignment, error) {
	var dailyAssignments []entity.DailyAssignment

	marshalledParam, err := d.json.Marshal(param)
	if err != nil {
		return dailyAssignments, err
	}

	if !param.BypassCache {
		dailyAssignments, err = d.getCacheList(ctx, fmt.Sprintf(getAllDailyAssignmentByKey, string(marshalledParam)))
		switch {
		case errors.Is(err, redis.Nil):
			d.log.Warn(ctx, fmt.Sprintf(entity.ErrorRedisNil, err.Error()))
		case err != nil:
			d.log.Warn(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
		default:
			return dailyAssignments, nil
		}
	}

	dailyAssignments, err = d.getAllSQL(ctx, param)
	if err != nil {
		return []entity.DailyAssignment{}, err
	}

	err = d.upsertCacheList(ctx, fmt.Sprintf(getAllDailyAssignmentByKey, string(marshalledParam)), dailyAssignments, d.redis.GetDefaultTTL(ctx))
	if err != nil {
		d.log.Error(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
	}

	return dailyAssignments, nil
}

func (d *dailyAssignment) UpdateDailyAssignmentToFalse(ctx context.Context) error {
	err := d.updateDailyAssignmentToFalse(ctx)
	if err != nil {
		return err
	}

	err = d.deleteCache(ctx, deleteDailyAssignmentPattern)
	if err != nil {
		d.log.Error(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
	}

	return nil
}
