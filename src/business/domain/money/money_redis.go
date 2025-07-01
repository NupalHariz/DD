package money

import (
	"context"
	"fmt"
	"time"

	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/codes"
	"github.com/reyhanmichiels/go-pkg/v2/errors"
)

const (
	getMoneyByKey         = "DD:money:get:%s"
	getAllMoneyByKey      = "DD:money:gets:%s"
	deleteMoneyKeyPattern = "DD:money:*"
)

func (m *money) upsertCache(ctx context.Context, key string, money entity.Money, ttl time.Duration) error {
	m.log.Debug(ctx, fmt.Sprintf("create cache with key %s and money %v", key, money))

	marshalledMoney, err := m.json.Marshal(money)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheMarshal, err.Error())
	}

	err = m.redis.SetEX(ctx, key, string(marshalledMoney), ttl)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheSetSimpleKey, err.Error())
	}

	m.log.Debug(ctx, fmt.Sprintf("success to set cache with key %s and body %v", key, money))
	return nil
}

func (m *money) upsertCacheList(ctx context.Context, key string, moneys []entity.Money, ttl time.Duration) error {
	m.log.Debug(ctx, fmt.Sprintf("create cache list with key: %s", key))

	marshalledMoney, err := m.json.Marshal(moneys)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheMarshal, err.Error())
	}

	m.log.Debug(ctx, fmt.Sprintf("SETTING REDIS WITH KEY %s, BODY %v, and TT %s", key, moneys, ttl))
	err = m.redis.SetEX(ctx, key, string(marshalledMoney), ttl)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheSetSimpleKey, err.Error())
	}

	m.log.Debug(ctx, fmt.Sprintf("success to set cache list with key %s and body %v", key, moneys))

	return nil
}

func (m *money) getCache(ctx context.Context, key string) (entity.Money, error) {
	var money entity.Money

	m.log.Debug(ctx, fmt.Sprintf("trying to get money from cache with key: %s", key))

	marshalMoney, err := m.redis.Get(ctx, key)
	if err != nil {
		return money, errors.NewWithCode(codes.CodeCacheGetSimpleKey, err.Error())
	}

	err = m.json.Unmarshal([]byte(marshalMoney), &money)
	if err != nil {
		return money, errors.NewWithCode(codes.CodeCacheUnmarshal, err.Error())
	}

	m.log.Debug(ctx, fmt.Sprintf("success to get money from cache with key: %s", key))
	
	return money, nil
}

func (m *money) getCacheList(ctx context.Context, key string) ([]entity.Money, error) {
	var moneys []entity.Money

	m.log.Debug(ctx, fmt.Sprintf("trying to get money list from cache with key: %s", key))

	marshalMoney, err := m.redis.Get(ctx, key)
	if err != nil {
		return moneys, errors.NewWithCode(codes.CodeCacheGetSimpleKey, err.Error())
	}

	err = m.json.Unmarshal([]byte(marshalMoney), &moneys)
	if err != nil {
		return moneys, errors.NewWithCode(codes.CodeCacheUnmarshal, err.Error())
	}

	m.log.Debug(ctx, fmt.Sprintf("success to get money list from cache with key: %s", key))
	
	return moneys, nil
}

func (m *money) deleteCache(ctx context.Context, key string) (error) {
	m.log.Debug(ctx, "remove money from cache")

	err := m.redis.Del(ctx, key)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheDeleteSimpleKey, err.Error())
	}

	m.log.Debug(ctx, "success to remove money from cache")
	return nil
}
