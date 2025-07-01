package money

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
	Create(ctx context.Context, param entity.MoneyInputParam) error
	Get(ctx context.Context, param entity.MoneyParam) (entity.Money, error)
	Update(ctx context.Context, updateParam entity.MoneyUpdateParam, moneyParam entity.MoneyParam) error
	GetAll(ctx context.Context, param entity.MoneyParam) ([]entity.Money, error)
}

type money struct {
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
	return &money{
		db:    param.Db,
		log:   param.Log,
		json:  param.Json,
		redis: param.Redis,
	}
}

func (m *money) Create(ctx context.Context, param entity.MoneyInputParam) error {
	err := m.createSql(ctx, param)
	if err != nil {
		return err
	}

	err = m.deleteCache(ctx, deleteMoneyKeyPattern)
	if err != nil {
		m.log.Error(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
	}

	return nil
}

func (m *money) Get(ctx context.Context, param entity.MoneyParam) (entity.Money, error) {
	var money entity.Money

	marshalledParam, err := m.json.Marshal(param)
	if err != nil {
		return money, err
	}

	if !param.BypassCache {
		money, err = m.getCache(ctx, fmt.Sprintf(getMoneyByKey, marshalledParam))
		switch {
		case errors.Is(err, redis.Nil):
			m.log.Warn(ctx, fmt.Sprintf(entity.ErrorRedisNil, err.Error()))
		case err != nil:
			m.log.Warn(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
		default:
			return money, nil
		}
	}

	money, err = m.getSQL(ctx, param)
	if err != nil {
		return entity.Money{}, err
	}

	err = m.upsertCache(ctx, fmt.Sprintf(getMoneyByKey, string(marshalledParam)), money, m.redis.GetDefaultTTL(ctx))
	if err != nil {
		m.log.Error(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
	}

	return money, nil
}

func (m *money) Update(ctx context.Context, updateParam entity.MoneyUpdateParam, moneyParam entity.MoneyParam) error {
	err := m.updateSQL(ctx, updateParam, moneyParam)
	if err != nil {
		return err
	}

	err = m.deleteCache(ctx, deleteMoneyKeyPattern)
	if err != nil {
		m.log.Error(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
	}

	return nil
}

func (m *money) GetAll(ctx context.Context, param entity.MoneyParam) ([]entity.Money, error) {
	var moneys []entity.Money

	marshalledParam, err := m.json.Marshal(param)
	if err != nil {
		return moneys, err
	}

	if !param.BypassCache {
		moneys, err = m.getCacheList(ctx, fmt.Sprintf(getMoneyByKey, marshalledParam))
		switch {
		case errors.Is(err, redis.Nil):
			m.log.Warn(ctx, fmt.Sprintf(entity.ErrorRedisNil, err.Error()))
		case err != nil:
			m.log.Warn(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
		default:
			return moneys, nil
		}
	}

	moneys, err = m.getAllSQL(ctx, param)
	if err != nil {
		return moneys, err
	}

	err = m.upsertCacheList(ctx, fmt.Sprintf(getMoneyByKey, string(marshalledParam)), moneys, m.redis.GetDefaultTTL(ctx))
	if err != nil {
		m.log.Error(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
	}

	return moneys, err
}
