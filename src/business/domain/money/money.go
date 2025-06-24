package money

import (
	"context"

	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/log"
	"github.com/reyhanmichiels/go-pkg/v2/sql"
)

type Interface interface {
	Create(ctx context.Context, param entity.MoneyInputParam) error
	Get(ctx context.Context, param entity.MoneyParam) (entity.Money, error)
	Update(ctx context.Context, updateParam entity.MoneyUpdateParam, moneyParam entity.MoneyParam) error
}

type money struct {
	db  sql.Interface
	log log.Interface
}

type InitParam struct {
	Db  sql.Interface
	Log log.Interface
}

func Init(param InitParam) Interface {
	return &money{
		db:  param.Db,
		log: param.Log,
	}
}

func (m *money) Create(ctx context.Context, param entity.MoneyInputParam) error {
	err := m.createSql(ctx, param)
	if err != nil {
		return err
	}

	return nil
}

func (m *money) Get(ctx context.Context, param entity.MoneyParam) (entity.Money, error) {
	money, err := m.getSQL(ctx, param)
	if err != nil {
		return entity.Money{}, err
	}

	return money, nil
}

func (m *money) Update(ctx context.Context, updateParam entity.MoneyUpdateParam, moneyParam entity.MoneyParam) error {
	err := m.updateSQL(ctx, updateParam, moneyParam)
	if err != nil {
		return err
	}

	return nil
}
