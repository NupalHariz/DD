package budget

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
	Create(ctx context.Context, param entity.BudgetInputParam) error
	UpdateExpense(ctx context.Context, updateParam entity.BudgetUpdateParam) error
	Update(ctx context.Context, updateParam entity.BudgetUpdateParam, budgetParam entity.BudgetParam) error
	GetAll(ctx context.Context, budgetParam entity.BudgetParam) ([]entity.Budget, error)
}

type budget struct {
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
	return &budget{
		db:    param.Db,
		log:   param.Log,
		json:  param.Json,
		redis: param.Redis,
	}
}

func (b *budget) Create(ctx context.Context, param entity.BudgetInputParam) error {
	err := b.createSQL(ctx, param)
	if err != nil {
		return err
	}

	err = b.deleteCache(ctx, deleteBudgetPattern)
	if err != nil {
		b.log.Warn(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
	}

	return nil
}

func (b *budget) UpdateExpense(ctx context.Context, updateParam entity.BudgetUpdateParam) error {
	err := b.updateExpenseSQL(ctx, updateParam)
	if err != nil {
		return err
	}

	err = b.deleteCache(ctx, deleteBudgetPattern)
	if err != nil {
		b.log.Warn(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
	}

	return nil
}

func (b *budget) Update(ctx context.Context, updateParam entity.BudgetUpdateParam, budgetParam entity.BudgetParam) error {
	err := b.updateSQL(ctx, updateParam, budgetParam)
	if err != nil {
		return err
	}

	err = b.deleteCache(ctx, deleteBudgetPattern)
	if err != nil {
		b.log.Warn(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
	}

	return nil
}

func (b *budget) GetAll(ctx context.Context, param entity.BudgetParam) ([]entity.Budget, error) {
	var budgets []entity.Budget
	marshalledBudgets, err := b.json.Marshal(param)
	if err != nil {
		return budgets, err
	}

	if !param.BypassCache {
		budgets, err = b.getCacheList(ctx, fmt.Sprintf(getAllBudgetByKey, string(marshalledBudgets)))
		switch {
		case errors.Is(err, redis.Nil):
			b.log.Warn(ctx, fmt.Sprintf(entity.ErrorRedisNil, err.Error()))
		case err != nil:
			b.log.Warn(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
		default:
			return budgets, nil
		}
	}

	budgets, err = b.getAllSQL(ctx, param)
	if err != nil {
		return budgets, err
	}

	err = b.upsertCacheList(ctx, fmt.Sprintf(getAllBudgetByKey, string(marshalledBudgets)), budgets, b.redis.GetDefaultTTL(ctx))
	if err != nil {
		b.log.Error(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
	}

	return budgets, nil
}
